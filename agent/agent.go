package agent

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"vikingPingvin/locker/server"
	"vikingPingvin/locker/server/protobuf"

	"github.com/golang/protobuf/proto"
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

// ExecuteAgent : Entrypoint for Locker agent start
func ExecuteAgent() {
	agent := &ArtifactAgent{Port: "27001"}
	agent.Start()
}
