package server

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"
	"vikingPingvin/locker/server/protobuf"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

type connectionType struct {
	// Send or Recieve
	intent int
}

type Server interface {
	Start() bool
	Stop() bool
}

type ArtifactServer struct {
	Address string
	Port    string
}

// TODO: Duplicate interface (agent.go)
// Generic Interface for Protobuf messages
type protoBufMessage interface {
	ProtoMessage()
	Reset()
	String() string
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
	defer func() {
		connection.Close()
		log.Info().Msg("Connection closing...")
	}()
	artifactReceived := false
	var artifactPath *os.File
	var hashInfoFromMeta []byte
	timeoutDuration := 5 * time.Second
	invalidCounterMax := 10
	invalidCounter := 0

	//payloadBuffer := make([]byte, 2048)
	sizePrefix := make([]byte, 4)
	nextPacket := make([]byte, 1)

	for artifactReceived != true {
		connection.SetReadDeadline(time.Now().Add(timeoutDuration))
		//n, _ := connection.Read(payloadBuffer)
		//bufReader := bufio.NewReader(connection).ReadBytes()

		_, _ = io.ReadFull(connection, sizePrefix)
		protoLength := int(binary.BigEndian.Uint32(sizePrefix))

		if protoLength > cap(nextPacket) {
			// Extend buffer size
			nextPacket = make([]byte, protoLength, 2*protoLength)
			nextPacket = nextPacket[:protoLength]
		} else if protoLength > len(nextPacket) {
			nextPacket = nextPacket[:protoLength]
		}

		//nextPacket := make([]byte, protoLength)
		_, _ = io.ReadFull(connection, nextPacket[:protoLength])

		decodedMessage := &protobuf.LockerMessage{}
		if err := proto.Unmarshal(nextPacket, decodedMessage); err != nil {
			log.Err(err).Msg("Error during unmarshalling")
		}

		// Determine ProtoBuf message using LockerMessage anyof structure
		if decodedMessage.GetMeta().ProtoReflect().IsValid() {
			artifactPath, hashInfoFromMeta = handleProtoMeta(decodedMessage.GetMeta())
		} else if decodedMessage.GetPackage().ProtoReflect().IsValid() {
			packageMessage := decodedMessage.GetPackage()
			if isPackageFinal := packageMessage.GetIsTerminated(); isPackageFinal == true {
				artifactReceived = true
				log.Info().Msg("Artifact fully received")
				compareArtifactHash(hashInfoFromMeta, artifactPath)
				break
			}
			handleProtoPackage(packageMessage, artifactPath)
		} else {
			if invalidCounter >= invalidCounterMax {
				log.Err(errors.New("Too many Invalid packets received")).Msg("Terminating connection")
				return
			}
			invalidCounter++
			fmt.Println("INVALID PROTOBUF MESSAGE")

		}
	}

}

func handleProtoMeta(metaMessage *protobuf.FileMeta) (file *os.File, hash []byte) {
	log.Info().
		Str("Artifact Name", metaMessage.GetFilename()).
		Str("NameSpace", metaMessage.GetNamespace()).
		Str("Project", metaMessage.GetProject()).
		Str("hash", fmt.Sprintf("%v", metaMessage.GetHash())).
		Msg("Artifact Meta info Recieved")

	// Create temp file where the payload will be appended
	return createTempFile(), metaMessage.GetHash()
}

func handleProtoPackage(packageMessage *protobuf.FilePackage, artifactPath *os.File) {
	//log.Debug().Str("payload", fmt.Sprintf("%s", string(packageMessage.GetPayload()))).Msg("Payload")
	writePayloadToFile(artifactPath, packageMessage.GetPayload())
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
	defer tmpArtifact.Close()
	log.Debug().Msgf("Using temp file: %s", tmpArtifact.Name())
	return tmpArtifact
}

func writePayloadToFile(filePath *os.File, payload []byte) {
	if filePath == nil {
		log.Err(errors.New("File pointer is empty. Can't write payload.")).Msg("Error writing payload. Maybe file handle is not created during META parsing.")
		os.Exit(111)
	}
	f, err := os.OpenFile(filePath.Name(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		log.Err(err).Msg("Error opening file for writing payload contents")
	}
	//fmt.Printf("%v", payload)
	writer := bufio.NewWriter(f)
	writer.Write(payload)
	writer.Flush()
}

// TODO: Duplicate hash command from agent.go:/hashFile
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

func compareArtifactHash(hashFromMeta []byte, tempPath *os.File) {
	calculatedHash := hashFile(tempPath.Name())

	if bytes.Compare(calculatedHash, hashFromMeta) == 0 {
		log.Info().Msg("Recieved Payload Hash is valid!")
	} else {
		log.Info().Msg("Recieved Payload Hash is Invalid!")
	}
}

// ExecuteServer : Entrypoint for Locker server start
func ExecuteServer() {
	server := &ArtifactServer{Address: "localhost", Port: "27001"}
	server.Start()

}
