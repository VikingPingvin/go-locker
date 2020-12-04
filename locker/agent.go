package locker

import (
	"net"
	"sync"

	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

// InputArgPath Relative or absolute path of files for Cobra CLI
var (
	InputArgPath      string
	InputArgNamespace string
	InputArgConsume   string
)

// InputData Populated at the start of the program
type InputData struct {
	FilePath  string
	FileName  string
	NameSpace string
	Project   string
	JobID     string
	FileHash  []byte
	ID        xid.ID
}

type Agent interface {
	Start() bool
	Stop() bool
}

type ArtifactAgent struct {
	Port string
}

func (a ArtifactAgent) Start(inputDataArray []*InputData) bool {
	// TODO: move sendConcurrent to config file
	const sendConcurrent = true

	var wg sync.WaitGroup
	for _, singleInputData := range inputDataArray {
		if sendConcurrent {
			wg.Add(1)
			go sendArtifactToServer(singleInputData, &wg)
		} else {
			sendArtifactToServer(singleInputData, &wg)
		}
	}

	wg.Wait()
	return true
}

func sendArtifactToServer(artifact *InputData, wg *sync.WaitGroup) {
	connection, err := net.Dial("tcp", "localhost:27001")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	log.Info().Msg("Agent connected to Locker Server...")

	// Send Metadata message
	parseAndSendMetaData(connection, artifact)

	// Send Payload message(s)
	parseAndSendPayload(connection, artifact)

	// Listen for ACK from server
	listenForACK(connection, artifact)

	wg.Done()
}

// ExecuteAgent : Entrypoint for Locker agent start
func ExecuteAgent() {
	// Handle input flags
	inputData := parseInputArguments()

	// Start Agent
	agent := &ArtifactAgent{Port: "27001"}
	agent.Start(inputData)
}
