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

// ArtifactData Represents a single artifact
type ArtifactData struct {
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
	Agent struct {
		ServerIP       string `yaml:"server_ip" env:"LOCKER_SERVER_IP" env-default:"127.0.0.1"`
		ServerPort     string `yaml:"server_port" env:"LOCKER_SERVER_PORT" env-default:"27001"`
		SendConcurrent bool   `yaml:"send_concurrent" env:"LOCKER_AGENT_CONCURRENT" env-default:"true"`
		LogPath        string `yaml:"log_path" env:"LOCKER_AGENT_LOG" env-default:"./locker-agent.log"`
		ArgPath        string
		ArgNamespace   string
		ArgConsume     string
	} `yaml:"agentconfig"`
}

// LockerAgentConfig a
var LockerAgentConfig *AgentConfig

type Agent interface {
	Start() bool
	Stop() bool
}

type ArtifactAgent struct {
	Configuration AgentConfig
}

func (a ArtifactAgent) Start(inputDataArray []*ArtifactData) bool {
	var sendConcurrent = a.Configuration.Agent.SendConcurrent

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

func sendArtifactToServer(artifact *ArtifactData, agentConfig *AgentConfig, wg *sync.WaitGroup) {
	serverAddr := fmt.Sprintf("%s:%s", agentConfig.Agent.ServerIP, agentConfig.Agent.ServerPort)
	connection, err := net.Dial("tcp", serverAddr)
	if err != nil {
		// TODO: if wg done is called first, panic is not executed.
		// Panic first and WG is not decremented -> goroutine error
		log.Panic().Err(err).Msg("Can't establish connection to the server!!!")
		wg.Done()
	}
	defer func() {
		connection.Close()
		wg.Done()
	}()

	log.Info().Msg("Agent connected to Locker Server...")

	// Send Metadata message
	parseAndSendMetaData(connection, artifact)

	// Send Payload message(s)
	parseAndSendPayload(connection, artifact)

	// Listen for ACK from server
	listenForACK(connection, artifact)
}

// ExecuteAgent : Entrypoint for Locker agent start
func ExecuteAgent() {
	// Handle input flags
	inputData := parseInputArguments()

	// Start Agent
	agent := &ArtifactAgent{Configuration: *LockerAgentConfig}
	agent.Start(inputData)
}
