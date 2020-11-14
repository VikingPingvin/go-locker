package agent

import (
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

	testString := strconv.FormatInt(server.INTENT_SERVER_RECIEVE, 2)

	fmt.Println(server.INTENT_SERVER_RECIEVE)
	fmt.Println(testString)
	connection.Write([]byte(testString))
	log.Info().Msg("Test string sent...")

	return true
}
