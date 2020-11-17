package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	pb "vikingPingvin/locker/server/protobuf"

	"google.golang.org/protobuf/proto"
)

func testMessage() {
	fmt.Println("Testing protobuf")

	testMsg := &pb.TestMessage{
		Id:      10,
		String_: "Test string",
		Uint:    123,
		Enum:    pb.MessageType_PACKAGE,
	}

	data, err := proto.Marshal(testMsg)
	if err != nil {
		log.Fatal("Marshalling error", err)
	}

	fmt.Println(data)

	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, data)
	fmt.Printf("Buffer Bytes: %d\n", buffer.Bytes())
}

func CreateMessage_FileInfo(id int32, messageType pb.MessageType, infoContent string) (message *pb.FileInfo, err error) {

	message = &pb.FileInfo{
		Id:      id,
		MsgType: messageType,
		Content: infoContent,
	}

	return message, err

}
