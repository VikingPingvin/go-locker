package agent

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"vikingPingvin/locker/server"
	"vikingPingvin/locker/server/protobuf"

	"github.com/golang/protobuf/proto"
	"github.com/rs/zerolog/log"
)

// FileInput String Flag for Cobra CMD input
var InputArg string

// InputData Populated at the start of the program
type InputData struct {
	FileInput string
	FileName  string
	FileHash  []byte
}

type Agent interface {
	Start() bool
	Stop() bool
}

type ArtifactAgent struct {
	Port string
}

// Generic Interface for Protobuf messages
type protoBufMessage interface {
	ProtoMessage()
	Reset()
	String() string
}

func (a ArtifactAgent) Start(inputData *InputData) bool {
	connection, err := net.Dial("tcp", "localhost:27001")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	log.Info().Msg("Agent connected to Locker Server...")

	//message, err := server.CreateMessage_FileInfo(123, protobuf.MessageType_META, "Test string valami content field-ben...")
	message, err := server.CreateMessage_FileMeta(
		2234,
		protobuf.MessageType_META,
		"test_namespace",
		"test_project",
		inputData.FileName,
		inputData.FileHash)
	if err != nil {
		panic(err)
	}

	// Send Metadata message
	sendProtoBufMessage(connection, message)

	// Send Payload message(s)
	sendParsedPayloadBytes(connection, inputData)

	return true
}

// Parse raw CLI input parameters to internal data structures
func parseInputArguments() *InputData {
	var err error
	var inputPath string

	if len(InputArg) == 0 {
		err = errors.New("--file empty")
		log.Err(err).Str("agent", "parseInputArguments").Msgf("No input file was given.")
	}
	if err != nil {
		os.Exit(1)
	}

	inputPath = InputArg
	if !filepath.IsAbs(InputArg) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Err(err).Msg("Error during CWD PATH parsing")
			os.Exit(1)
		}
		log.Debug().Msgf("Relative path of input: %s", InputArg)
		inputPath = filepath.Join(cwd, InputArg)

	}
	data := &InputData{
		FileInput: inputPath,
	}
	return data
}

// If --file cli input is not null, parse file
func parseFileMetaData(inputData *InputData) (fileInfo os.FileInfo, err error) {
	fileInfo, err = os.Stat(inputData.FileInput)
	if os.IsNotExist(err) {
		log.Error().Msgf("Parsing file Input error: %v", err)
		return fileInfo, err
	}
	inputData.FileHash = hashFile(inputData.FileInput)
	inputData.FileName = fileInfo.Name()

	log.Info().
		Str("file name", fileInfo.Name()).
		Str("size", fmt.Sprintf("%d", fileInfo.Size())).
		Str("hash", fmt.Sprintf("%v", inputData.FileHash)).
		Msg("Artifact metadata parsing finished")

	return fileInfo, err
}

//func sendParsedPayloadBytes(bytes *[]byte, numBytes int) {
func sendParsedPayloadBytes(connection net.Conn, inputData *InputData) {
	//fw, _ := os.OpenFile("E:\\WorkSpace\\Go\\artifact-server\\readmecopy.channel", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//defer fw.Close()
	//writer := bufio.NewWriter(fw)
	//writer.Write((*bytes)[:numBytes])
	//writer.Flush()

	f, err := os.Open(inputData.FileInput)
	defer f.Close()
	if err != nil {
		log.Error().Msgf("Cannot open file %s", inputData.FileInput)
	}

	//fw, _ := os.Create(inputData.FileInput + "bufwritten")
	//defer fw.Close()
	//writer := bufio.NewWriter(fw)
	reader := bufio.NewReader(f)
	isPayloadFinal := false

	buffer := make([]byte, 1024)
	for {
		n, ioErr := reader.Read(buffer)
		if ioErr == io.EOF {
			isPayloadFinal = true
			// Send terminating payload protobuf message
			terminalMessage, err := server.CreateMessage_FilePackage(
				2234,
				protobuf.MessageType_PACKAGE,
				make([]byte, 1),
				isPayloadFinal,
			)
			if err != nil {
				log.Fatal().Msg("Fatal Error during payload protobuf assembly")
			}
			sendProtoBufMessage(connection, terminalMessage)
			break
		}
		//log.Debug().Str("payload", fmt.Sprintf("%v", (buffer)[:n])).Msgf("Protobuf Payload Type")
		message, err := server.CreateMessage_FilePackage(
			2234,
			protobuf.MessageType_PACKAGE,
			(buffer)[:n],
			isPayloadFinal)

		if err != nil {
			log.Fatal().Msg("Fatal Error during payload protobuf assembly")
		}

		sendProtoBufMessage(connection, message)
	}
}

// Generic protobuf sender that accepts an interface to the defined messages
func sendProtoBufMessage(connection net.Conn, message *protobuf.LockerMessage) {
	log.Debug().Str("Proto_message", fmt.Sprintf("%v", message)).Msg("sendProtoBufMessage")
	dataToSend, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, dataToSend)
	connection.Write(buffer.Bytes())
	log.Debug().Msgf("Protobuf Msg Size: %d", len(buffer.Bytes()))
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

// ExecuteAgent : Entrypoint for Locker agent start
func ExecuteAgent() {
	// Handle input flags
	inputData := parseInputArguments()

	parseFileMetaData(inputData)

	// Start Agent
	agent := &ArtifactAgent{Port: "27001"}
	agent.Start(inputData)
}
