package locker

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"

	"vikingPingvin/locker/locker/messaging"
	"vikingPingvin/locker/locker/messaging/protobuf"

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

func (a ArtifactAgent) Start(inputData *InputData) bool {
	connection, err := net.Dial("tcp", "localhost:27001")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	log.Info().Msg("Agent connected to Locker Server...")

	// Send Metadata message
	parseAndSendMetaData(connection, inputData)

	// Send Payload message(s)
	parseAndSendPayload(connection, inputData)

	// Listen for ACK from server
	listenForACK(connection)

	return true
}

// If --file cli input is not null, parse file
func parseAndSendMetaData(connection net.Conn, inputData *InputData) (fileInfo os.FileInfo, err error) {
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

	message, err := messaging.CreateMessage_FileMeta(
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
	log.Info().Msg("Sending MetaData Packet")
	messaging.SendProtoBufMessage(connection, message)

	return fileInfo, err
}

//func parseAndSendPayload(bytes *[]byte, numBytes int) {
func parseAndSendPayload(connection net.Conn, inputData *InputData) {

	f, err := os.Open(inputData.FileInput)
	defer f.Close()
	if err != nil {
		log.Error().Msgf("Cannot open file %s", inputData.FileInput)
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
				2234,
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
			2234,
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

func listenForACK(connection net.Conn) {

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

// ExecuteAgent : Entrypoint for Locker agent start
func ExecuteAgent() {
	// Handle input flags
	inputData := parseInputArguments()

	// Start Agent
	agent := &ArtifactAgent{Port: "27001"}
	agent.Start(inputData)
}