package server

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

type connectionType struct {
	// Send or Recieve
	intent int
}

type Server interface {
	Start() bool
	Stop() bool
}

type ArtifactServer struct {
	Address string
	Port    string
}

func (s ArtifactServer) Start() bool {
	listenAddr := fmt.Sprintf("%s:%s", s.Address, s.Port)
	server, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Info().Msg("Error listening on port!")
		os.Exit(1)
	}

	defer server.Close()
	log.Info().Msg("Locker server started! Listening for connections...")
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
	log.Info().Msg("Locker client connected.")
	defer connection.Close()

	intentSize := INTENT_STREAM_SIZE
	intentBuffer := make([]byte, intentSize)

	connection.Read(intentBuffer)
	fmt.Println(intentBuffer)
	intent := binary.BigEndian.Uint16(intentBuffer)

	fmt.Println(intent)
	fmt.Println(strconv.ParseInt(string(intent), 10, 16))
}

// ExecuteServer... Entrypoint for Locker server start
func ExecuteServer() {
	server := &ArtifactServer{Address: "localhost", Port: "27001"}
	server.Start()
}
