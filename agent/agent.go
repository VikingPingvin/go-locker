package agent

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"vikingPingvin/locker/server"

	"github.com/rs/zerolog/log"
)

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

	testString := strconv.FormatInt((int64(server.INTENT_SERVER_RECIEVE)), 2)

	fmt.Printf("Sending source: %d\n", server.INTENT_SERVER_RECIEVE)
	fmt.Printf("Sending converted: %s\n", testString)

	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, server.INTENT_SERVER_SEND)
	connection.Write(buffer.Bytes())

	fmt.Printf("Buffer Bytes: %d\n", buffer.Bytes())
	log.Info().Msg("Test string sent...")

	return true
}

// ExecuteAgent... Entrypoint for Locker agent start
func ExecuteAgent() {
	agent := &ArtifactAgent{Port: "27001"}
	agent.Start()
}
