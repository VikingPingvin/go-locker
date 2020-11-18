package agent

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"vikingPingvin/locker/server"
	"vikingPingvin/locker/server/protobuf"

	"github.com/golang/protobuf/proto"
	"github.com/rs/zerolog/log"
)

// FileInput String Flag for Cobra CMD input
var FileInput string

// InputData Populated at the start of the program
type InputData struct {
	FileInput string
}

type Agent interface {
	Start() bool
	Stop() bool
}

type ArtifactAgent struct {
	Port string
}

func (a ArtifactAgent) Start() bool {
	connection, err := net.Dial("tcp", "localhost:27001")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	log.Info().Msg("Agent connected to Locker Server...")

	message, err := server.CreateMessage_FileInfo(123, protobuf.MessageType_META, "Test string valami content field-ben...")
	if err != nil {
		panic(err)
	}

	dataToSend, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}

	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, dataToSend)
	connection.Write(buffer.Bytes())

	fmt.Printf("Buffer Bytes: %d\n", buffer.Bytes())
	log.Info().Msg("Test string sent...")

	return true
}

func parseInput() *InputData {
	if len(FileInput) == 0 {
		log.Info().Msg("No input file was given.")
	}
	data := &InputData{
		FileInput: FileInput,
	}

	log.Info().Msgf("Input structure: %v", data)

	return data
}

func parseFile(filePath *string) (err error) {
	if _, err := os.Stat(*filePath); os.IsNotExist(err) {
		log.Error().Msgf("Parsing file Input error: %v", err)
	}
	return err
}

// ExecuteAgent : Entrypoint for Locker agent start
func ExecuteAgent() {
	// Handle input flags
	inputData := parseInput()

	parseFile(&inputData.FileInput)

	// Start Agent
	agent := &ArtifactAgent{Port: "27001"}
	agent.Start()
}
