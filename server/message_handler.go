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
func CreateMessage_FileMeta(id int32, msgType pb.MessageType, namesSpace string, projectName string, fileName string, fileHash []byte) (protoMessage *pb.LockerMessage, err error) {

	message := &pb.FileMeta{
		Id:        id,
		MsgType:   msgType,
		Namespace: namesSpace,
		Project:   projectName,
		Filename:  fileName,
		Hash:      fileHash,
	}

	test := &pb.LockerMessage_Meta{
		Meta: message,
	}

	testfull := &pb.LockerMessage{
		Message: test,
	}

	return testfull, err
	//protoMessage = &pb.LockerMessage{
	//	Message: &message,
	//}
	//return protoMessage, err
}

func CreateMessage_FilePackage(id int32, msgType pb.MessageType, payload []byte, isFinalPayload bool) (protoMessage *pb.LockerMessage, err error) {

	//message = &pb.FilePackage{
	//	Id:      id,
	//	MsgType: msgType,
	//	Payload: payload,
	//}

	protoMessage = &pb.LockerMessage{
		Message: &pb.LockerMessage_Package{
			Package: &pb.FilePackage{
				Id:           id,
				MsgType:      msgType,
				Payload:      payload,
				IsTerminated: isFinalPayload,
			},
		},
	}

	return protoMessage, err
}
