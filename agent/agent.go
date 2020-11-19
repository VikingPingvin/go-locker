package agent

import (
	"bufio"
	"bytes"
	"crypto/sha256"
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

	//message, err := server.CreateMessage_FileInfo(123, protobuf.MessageType_META, "Test string valami content field-ben...")
	message, err := server.CreateMessage_FileMeta(
		2234,
		protobuf.MessageType_META,
		inputData.FileName,
		inputData.FileHash)
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

	//fmt.Printf("Buffer Bytes: %d\n", buffer.Bytes())
	//log.Info().Msg("Test string sent...")

	return true
}

// Parse raw CLI input parameters to internal data structures
func parseInputArguments() *InputData {
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

	inputData.FileHash = hashFile(inputData.FileInput)
	inputData.FileName = fileInfo.Name()
	log.Info().Msgf("Calculated SHA256 Hash: %v", inputData.FileHash)

	f, err := os.Open(inputData.FileInput)
	defer f.Close()
	if err != nil {
		log.Error().Msgf("Cannot open file %s", inputData.FileInput)
	}

	fw, _ := os.Create(inputData.FileInput + "bufwritten")
	defer fw.Close()
	writer := bufio.NewWriter(fw)
	reader := bufio.NewReader(f)
	bufChannel := make(chan []byte)

	//go sendParsedPayload(&bufChannel)

	buffer := make([]byte, 1024)
	for {
		n, ioErr := reader.Read(buffer)
		if ioErr == io.EOF {
			close(bufChannel)
			break
		}
		writer.Write(buffer[:n])
		//func() { bufChannel <- buffer[:n] }()
		sendParsedPayloadBytes(&buffer, n)
	}

	writer.Flush()

	return fileInfo, err
}

// sendParsedPayload Recieves []byte from bufferChannel and operates on it.
//
// value <- *bufferChannel
func sendParsedPayload(bufferChannel *chan []byte) {
	fw, _ := os.OpenFile("E:\\WorkSpace\\Go\\artifact-server\\readmecopy.channel", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer fw.Close()
	writer := bufio.NewWriter(fw)
	for elem := range *bufferChannel {
		fmt.Println(len(elem))
		fmt.Println("GOROUTINE")
		writer.Write(elem)
	}
	writer.Flush()
}

func sendParsedPayloadBytes(bytes *[]byte, numBytes int) {
	fw, _ := os.OpenFile("E:\\WorkSpace\\Go\\artifact-server\\readmecopy.channel", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer fw.Close()
	writer := bufio.NewWriter(fw)
	writer.Write((*bytes)[:numBytes])
	writer.Flush()
}

func hashFile(path string) (hash []byte) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Error().Msgf("Cannot open file %s", path)
	}

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		log.Error().Msgf("Error calculating SHA256 Hash: %v", err)
	}
	return hasher.Sum(nil)
}

// ExecuteAgent : Entrypoint for Locker agent start
func ExecuteAgent() {
	// Handle input flags
	inputData := parseInputArguments()

	//fileInfo, _ := parseFile(inputData)
	parseFile(inputData)

	// Start Agent
	agent := &ArtifactAgent{Port: "27001"}
	agent.Start(inputData)
}
