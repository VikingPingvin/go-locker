package agent

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
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
	Payload   string
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

// Parse raw CLI input parameters to internal data structures
func parseInput() *InputData {
	if len(FileInput) == 0 {
		log.Info().Msg("No input file was given.")
	}
	data := &InputData{
		FileInput: FileInput,
	}
	return data
}

// If --file cli input is not null, parse file
func parseFile(inputData *InputData) (fileInfo os.FileInfo, err error) {
	fileInfo, err = os.Stat(inputData.FileInput)
	if os.IsNotExist(err) {
		log.Error().Msgf("Parsing file Input error: %v", err)
		return fileInfo, err
	}

	log.Info().Msgf("File Name: %s", fileInfo.Name())
	log.Info().Msgf("File size: %d", fileInfo.Size())

	f, err := os.Open(inputData.FileInput)
	defer f.Close()
	if err != nil {
		log.Error().Msgf("Cannot open file %s", inputData.FileInput)
	}

	fw, _ := os.Create(inputData.FileInput + "bufwritten")
	defer fw.Close()
	writer := bufio.NewWriter(fw)
	reader := bufio.NewReader(f)

	buffer := make([]byte, 1024)
	for {
		n, ioErr := reader.Read(buffer)
		writer.Write(buffer[:n])
		//fmt.Println(n)
		if ioErr == io.EOF {
			break
		}
	}
	writer.Flush()
	return fileInfo, err
}

// ExecuteAgent : Entrypoint for Locker agent start
func ExecuteAgent() {
	// Handle input flags
	inputData := parseInput()

	//fileInfo, _ := parseFile(inputData)
	parseFile(inputData)

	// Start Agent
	agent := &ArtifactAgent{Port: "27001"}
	agent.Start()
}
