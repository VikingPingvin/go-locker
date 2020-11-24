package locker

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"
	"vikingPingvin/locker/locker/messaging/protobuf"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

type connectionType struct {
	// Send or Recieve
	intent int
}

// Struct containing file data from received MetaData message
type metaInfo struct {
	fileHash []byte
	fileName string
}

type Server interface {
	Start() bool
	Stop() bool
}

type ArtifactServer struct {
	Address string
	Port    string
}

func (s ArtifactServer) Start() bool {
	listenAddr := fmt.Sprintf("%s:%s", s.Address, s.Port)
	server, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Info().Msg("Error listening on port!")
		os.Exit(1)
	}

	defer server.Close()
	log.Info().Msg("Locker server started! Listening for connections...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		go handleConnection(connection)

	}
}

func handleConnection(connection net.Conn) {
	log.Info().Msgf("Locker client connected: %s", connection.RemoteAddr().String())
	timeStart := time.Now()
	defer func() {
		connection.Close()
		log.Info().Msg("Connection closing...")
	}()
	artifactReceived := false
	var artifactPath *os.File

	metaData := metaInfo{}

	//var hashInfoFromMeta []byte
	timeoutDuration := 5 * time.Second
	invalidCounterMax := 10
	invalidCounter := 0

	// writeFileBuffer defaults to 1 Mb, used for disk IO flushing
	writeFileBuffer := make([]byte, 0, 1000000)

	var writer *bufio.Writer

	// sizePrefix is 4 bytes protobug message size
	sizePrefix := make([]byte, 4)
	// nextPacket auto expanding slice, contains a protobuf packet
	nextPacket := make([]byte, 1)

	defer func() {
		artifactPath.Close()

		elapsedTime := time.Since(timeStart)
		log.Debug().Msgf("Client Connection took: %s", elapsedTime)
	}()

	for artifactReceived != true {
		connection.SetReadDeadline(time.Now().Add(timeoutDuration))

		_, _ = io.ReadFull(connection, sizePrefix)
		protoLength := int(binary.BigEndian.Uint32(sizePrefix))

		if protoLength > cap(nextPacket) {
			// Extend buffer size
			nextPacket = make([]byte, protoLength, 2*protoLength)
			nextPacket = nextPacket[:protoLength]
		} else if protoLength > len(nextPacket) {
			nextPacket = nextPacket[:protoLength]
		}

		_, _ = io.ReadFull(connection, nextPacket[:protoLength])

		decodedMessage := &protobuf.LockerMessage{}
		if err := proto.Unmarshal(nextPacket[:protoLength], decodedMessage); err != nil {
			log.Err(err).Msg("Error during unmarshalling")
		}

		// Determine ProtoBuf message using LockerMessage anyof structure
		if decodedMessage.GetMeta().ProtoReflect().IsValid() {
			artifactPath, metaData = handleProtoMeta(decodedMessage.GetMeta())
			writer = bufio.NewWriter(artifactPath)

		} else if decodedMessage.GetPackage().ProtoReflect().IsValid() {
			packageMessage := decodedMessage.GetPackage()

			// If received the last payload package, flush and close the tmp file
			if isPackageFinal := packageMessage.GetIsTerminated(); isPackageFinal == true {
				artifactReceived = true
				log.Info().Msg("Artifact payload received")

				writer.Write(writeFileBuffer)
				writer.Flush()
				artifactPath.Close()
				break
			}
			handleProtoPackage(packageMessage, writer, &writeFileBuffer)
		} else {
			if invalidCounter >= invalidCounterMax {
				log.Err(errors.New("Too many Invalid packets received")).Msg("Terminating connection")
				return
			}
			invalidCounter++
			fmt.Println("INVALID PROTOBUF MESSAGE")

		}
	}

	// Calculate Hash of tmp file
	if compareArtifactHash(metaData.fileHash, artifactPath) {
		//Rename file
		baseDir := filepath.Dir(artifactPath.Name())
		newPath := filepath.Join(baseDir, metaData.fileName)
		os.Rename(artifactPath.Name(), newPath)
		log.Info().Msgf("Artifact ready: %s", newPath)

	}
}

func handleProtoMeta(metaMessage *protobuf.FileMeta) (file *os.File, metaData metaInfo) {
	log.Info().
		Str("Artifact Name", metaMessage.GetFilename()).
		Str("NameSpace", metaMessage.GetNamespace()).
		Str("Project", metaMessage.GetProject()).
		Str("hash", fmt.Sprintf("%v", metaMessage.GetHash())).
		Msg("Artifact Meta info Recieved")

	metaData.fileHash = metaMessage.GetHash()
	metaData.fileName = metaMessage.GetFilename()

	// Create temp file where the payload will be appended
	return createTempFile(), metaData
}

func handleProtoPackage(packageMessage *protobuf.FilePackage, writer *bufio.Writer, writeFileBuffer *[]byte) {
	writePayloadToFile(writer, packageMessage, writeFileBuffer)
}

func createTempFile() *os.File {
	cwd, err := os.Getwd()
	if err != nil {
		log.Err(err).Str("tempfile", "getwd").Msg("Error getting working directory")
		os.Exit(1)
	}
	outputDir := filepath.Join(cwd, "out")
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, os.ModeDir)
	}

	tmpArtifact, err := ioutil.TempFile(outputDir, "~temp_locker_artifact_")
	if err != nil {
		log.Err(err).Str("tempfile", "ioutil").Msg("Error creating temp file on path")
	}
	log.Debug().Msgf("Using temp file: %s", tmpArtifact.Name())
	return tmpArtifact
}

func writePayloadToFile(writer *bufio.Writer, payload *protobuf.FilePackage, writeFileBuffer *[]byte) {

	// If payload > len + cap: flush io and reslice to size 0
	payloadBytes := payload.GetPayload()
	if len(payloadBytes) > len(*writeFileBuffer)+cap(*writeFileBuffer) {
		writer.Write(*writeFileBuffer)
		writer.Flush()

		*writeFileBuffer = (*writeFileBuffer)[:0]
	} else {
		*writeFileBuffer = append(*writeFileBuffer, payloadBytes...)
	}

}

// TODO: Duplicate hash command from agent.go:/hashFile
// Given a valid file path, returns a SHA256 hash
//func hashFile(path string) (hash []byte) {
//	f, err := os.Open(path)
//	defer f.Close()
//	if err != nil {
//		log.Err(err).Msgf("Cannot open file %s", path)
//	}
//
//	hasher := sha256.New()
//	if _, err := io.Copy(hasher, f); err != nil {
//		log.Err(err).Msg("Error calculating SHA256 Hash")
//	}
//	return hasher.Sum(nil)
//}

func compareArtifactHash(hashFromMeta []byte, tempPath *os.File) bool {
	calculatedHash := hashFile(tempPath.Name())

	if bytes.Compare(calculatedHash, hashFromMeta) == 0 {
		log.Info().Msg("Recieved Payload Hash is valid!")
		return true
	} else {
		log.Info().Msg("Recieved Payload Hash is Invalid!")
		return false
	}
}

// ExecuteServer : Entrypoint for Locker server start
func ExecuteServer() {
	server := &ArtifactServer{Address: "localhost", Port: "27001"}
	server.Start()

}
