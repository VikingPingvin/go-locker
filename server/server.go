package server

import (
	"fmt"
	"net"
	"os"

	"github.com/rs/zerolog/log"
)

type Server interface {
	Start() bool
	Stop() bool
}

type artifactServer struct {
	address string
	port    string
}

func (s artifactServer) Start() bool {
	listenAddr := fmt.Sprintf("%s:%s", s.address, s.port)
	server, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Info().Msg("Error listening on port!")
		os.Exit(1)
	}

	defer server.Close()
	log.Info().Msg("Server started! Listening for connections...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {

}
