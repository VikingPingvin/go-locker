package locker

import (
	"fmt"
	"net"
	"sync"

	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

// InputArgPath Relative or absolute path of files for Cobra CLI
//var (
//	InputArgPath      string
//	InputArgNamespace string
//	InputArgConsume   string
//)

// InputData Represents a single artifact
type InputData struct {
	FilePath  string
	FileName  string
	NameSpace string
	Project   string
	JobID     string
	FileHash  []byte
	ID        xid.ID
}

// AgentConfig holds configuration values
type AgentConfig struct {
	ServerIP       string
	ServerPort     string
	SendConcurrent bool
	LogPath        string
	ArgPath        string
	ArgNamespace   string
	ArgConsume     string
}

// LockerConfig a
var LockerConfig *AgentConfig

type Agent interface {
	Start() bool
	Stop() bool
}

type ArtifactAgent struct {
	Configuration AgentConfig
}

func (a ArtifactAgent) Start(inputDataArray []*InputData) bool {
	var sendConcurrent = a.Configuration.SendConcurrent

	var wg sync.WaitGroup
	for _, singleInputData := range inputDataArray {
		if sendConcurrent {
			wg.Add(1)
			go sendArtifactToServer(singleInputData, &a.Configuration, &wg)
		} else {
			sendArtifactToServer(singleInputData, &a.Configuration, &wg)
		}
	}

	wg.Wait()
	return true
}

func sendArtifactToServer(artifact *InputData, agentConfig *AgentConfig, wg *sync.WaitGroup) {
	serverAddr := fmt.Sprintf("%s:%s", agentConfig.ServerIP, agentConfig.ServerPort)
	connection, err := net.Dial("tcp", serverAddr)
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
	agent := &ArtifactAgent{Configuration: *LockerConfig}
	agent.Start(inputData)
}
