package server

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"vikingPingvin/locker/server/protobuf"

	"github.com/golang/protobuf/proto"
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

	// TODO: Defer closes connection so subsequent messages are not read
	// BUG ^^

	buffer := make([]byte, 1024)
	n, _ := connection.Read(buffer)
	decodedMessage := &protobuf.FileMeta{}
	if err := proto.Unmarshal(buffer[:n], decodedMessage); err != nil {
		log.Err(err).Msg("Error during unmarshalling")
	}

	//bufReader := bufio.NewReader(connection)
	////connection.Read(readBuffer.Bytes())
	//bufReader.Read()

	//decodedMessage := &protobuf.FileMeta{}
	//proto.Unmarshal(readBuffer.Bytes(), decodedMessage)

	log.Info().
		Str("Artifact Name", decodedMessage.GetFilename()).
		Str("NameSpace", decodedMessage.GetNamespace()).
		Str("Project", decodedMessage.GetProject()).
		Str("hash", fmt.Sprintf("%v", decodedMessage.GetHash())).
		Msg("Artifact Meta info Recieved")

	// Create temp file where the payload will be appended
	createTempFile()
}

func createTempFile() *os.File {
	cwd, err := os.Getwd()
	if err != nil {
		log.Err(err).Str("tempfile", "getwd").Msg("Error getting working directory")
		os.Exit(1)
	}
	outputDir := filepath.Join(cwd, "out")
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, os.ModeDir)
	}

	tmpArtifact, err := ioutil.TempFile(outputDir, "~temp_locker_artifact_")
	if err != nil {
		log.Err(err).Str("tempfile", "ioutil").Msg("Error creating temp file on path")
	}
	defer tmpArtifact.Close()
	return tmpArtifact
}

// ExecuteServer : Entrypoint for Locker server start
func ExecuteServer() {
	server := &ArtifactServer{Address: "localhost", Port: "27001"}
	server.Start()

}
