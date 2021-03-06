package locker

import (
	"bufio"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"vikingPingvin/locker/locker/messaging"
	"vikingPingvin/locker/locker/messaging/protobuf"

	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

// If --file cli input is not null, parse file
func parseAndSendMetaData(connection net.Conn, artifactData *ArtifactData) (fileInfo os.FileInfo, err error) {
	fileInfo, err = os.Stat(artifactData.FilePath)
	if os.IsNotExist(err) {
		log.Error().Msgf("Parsing file Input error: %v", err)
		return fileInfo, err
	}
	artifactData.FileHash = hashFile(artifactData.FilePath)
	artifactData.FileName = fileInfo.Name()

	log.Info().
		Str("file name", fileInfo.Name()).
		Str("Namespace", fmt.Sprintf("%s/%s/%s", artifactData.NameSpace, artifactData.Project, artifactData.JobID)).
		Str("size", fmt.Sprintf("%d", fileInfo.Size())).
		Str("hash", fmt.Sprintf("%v", artifactData.FileHash)).
		Str("id", artifactData.ID.String()).
		Msg("Artifact metadata parsing finished")

	message, err := messaging.CreateMessage_FileMeta(
		artifactData.ID.Bytes(),
		protobuf.MessageType_META,
		artifactData.NameSpace,
		artifactData.Project,
		artifactData.JobID,
		artifactData.FileName,
		artifactData.FileHash)
	if err != nil {
		panic(err)
	}

	// Send Metadata message
	log.Info().Msg("Sending MetaData Packet")
	messaging.SendProtoBufMessage(connection, message)

	return fileInfo, err
}

//func parseAndSendPayload(bytes *[]byte, numBytes int) {
func parseAndSendPayload(connection net.Conn, artifactData *ArtifactData) {

	f, err := os.Open(artifactData.FilePath)
	defer f.Close()
	if err != nil {
		log.Error().Msgf("Cannot open file %s", artifactData.FilePath)
	}

	log.Info().Msg("Started sending Payload Packets...")
	reader := bufio.NewReader(f)
	isPayloadFinal := false

	buffer := make([]byte, 1024)
	for {
		n, ioErr := reader.Read(buffer)
		if ioErr == io.EOF {
			isPayloadFinal = true
			// Send terminating payload protobuf message
			terminalMessage, err := messaging.CreateMessage_FilePackage(
				artifactData.ID.Bytes(),
				protobuf.MessageType_PACKAGE,
				make([]byte, 1),
				isPayloadFinal,
			)
			if err != nil {
				log.Fatal().Msg("Fatal Error during payload protobuf assembly")
			}
			messaging.SendProtoBufMessage(connection, terminalMessage)
			break
		}

		message, err := messaging.CreateMessage_FilePackage(
			artifactData.ID.Bytes(),
			protobuf.MessageType_PACKAGE,
			(buffer)[:n],
			isPayloadFinal)

		if err != nil {
			log.Fatal().Msg("Fatal Error during payload protobuf assembly")
		}

		messaging.SendProtoBufMessage(connection, message)
	}
	log.Info().Msg("Finished sending Payload Packets...")
}

// Given a valid file path, returns a SHA256 hash
func hashFile(path string) (hash []byte) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Err(err).Msgf("Cannot open file %s", path)
	}

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		log.Err(err).Msg("Error calculating SHA256 Hash")
	}
	return hasher.Sum(nil)
}

func listenForACK(connection net.Conn, artifactData *ArtifactData) {

	// TODO: Make into const in message_handler (also server.go)
	// sizePrefix is 4 bytes protobug message size
	sizePrefix := make([]byte, 4)

	_, _ = io.ReadFull(connection, sizePrefix)
	protoLength := int(binary.BigEndian.Uint32(sizePrefix))

	ackPacketRaw := make([]byte, protoLength)
	_, _ = io.ReadFull(connection, ackPacketRaw)
	genericProto := &protobuf.LockerMessage{}
	if err := proto.Unmarshal(ackPacketRaw, genericProto); err != nil {
		log.Err(err).Msg("Error during unmarshalling")
	}

	if genericProto.GetAck().ProtoReflect().IsValid() {
		ackPacket := genericProto.GetAck()

		serverResult := ackPacket.GetServerSuccess()
		ackID, _ := xid.FromBytes(ackPacket.GetId())
		if ackID != artifactData.ID {
			log.Warn().
				Str("respone_id", ackID.String()).
				Str("original_id", artifactData.ID.String()).
				Msg("Response ID mismatch.")
		}

		log.Info().
			Str("id_back", ackID.String()).
			Msgf("ACK packet recieved from server with success flag: %v", serverResult)
	}
}

// Parse raw CLI input parameters to internal data structures
func parseInputArguments() []*ArtifactData {
	var err error
	var inputPath string
	const inputPathSeparator = ","

	// Initialize LockerConfig to avoid nil dereference errors
	//LockerConfig = &AgentConfig{}

	if len(LockerAgentConfig.Agent.ArgPath) == 0 {
		err = errors.New("--file empty")
		log.Err(err).Str("agent", "parseInputArguments").Msgf("No input file was given.")
	}
	if len(LockerAgentConfig.Agent.ArgNamespace) == 0 {
		err = errors.New("--namespace empty")
		log.Err(err).Str("agent", "parseInputArguments").Msgf("No input namespace was given.")
	}
	if err != nil {
		os.Exit(1)
	}

	//
	// Store information about Namespace, Project and Job-ID
	fullNameSpace := LockerAgentConfig.Agent.ArgNamespace
	namePaths := strings.Split(fullNameSpace, "/")
	if len(namePaths) != 3 {
		err = errors.New("Namespace must contain 3 values separated by '/'")
		log.Err(err).Msg("Namespace values not valid")
		os.Exit(1)
	}

	//
	// dataArray contains an *InputData, len(inputPathSlice) times
	inputPathSlice := strings.Split(LockerAgentConfig.Agent.ArgPath, inputPathSeparator)
	dataArray := make([]*ArtifactData, len(inputPathSlice))

	for i, path := range inputPathSlice {
		if !filepath.IsAbs(path) {
			cwd, err := os.Getwd()
			if err != nil {
				log.Err(err).Msg("Error during CWD PATH parsing")
				os.Exit(1)
			}
			log.Debug().Msgf("Relative path of input: %s", path)
			inputPath = filepath.Join(cwd, path)

		} else {
			inputPath = path
		}

		dataArray[i] = &ArtifactData{
			FilePath:  inputPath,
			NameSpace: namePaths[0],
			Project:   namePaths[1],
			JobID:     namePaths[2],
			ID:        xid.New(),
		}
	}

	return dataArray
}
