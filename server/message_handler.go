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

func CreateMessage_FileInfo(id int32, messageType pb.MessageType, payload string) (message *pb.FileInfo, err error) {

	message = &pb.FileInfo{
		Id:      id,
		MsgType: messageType,
		Payload: payload,
	}
	return message, err
}

// CreateMessage_FileMeta Create ProtoBuf message.
//  id
//  msgType
//  fileName
//  fileHash
func CreateMessage_FileMeta(id int32, msgType pb.MessageType, namesSpace string, projectName string, fileName string, fileHash []byte) (message *pb.FileMeta, err error) {

	message = &pb.FileMeta{
		Id:        id,
		MsgType:   msgType,
		Namespace: namesSpace,
		Project:   projectName,
		Filename:  fileName,
		Hash:      fileHash,
	}

	return message, err
}

func CreateMessage_FilePackage(id int32, msgType pb.MessageType, payload []byte) (message *pb.FilePackage, err error) {

	message = &pb.FilePackage{
		Id:      id,
		MsgType: msgType,
		Payload: payload,
	}

	return message, err
}
