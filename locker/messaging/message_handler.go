package messaging

import (
	"bytes"
	"encoding/binary"
	"net"
	"vikingPingvin/locker/locker/messaging/protobuf"
	pb "vikingPingvin/locker/locker/messaging/protobuf"

	"google.golang.org/protobuf/proto"
)

func CreateMessage_ServerACk(id []byte, messageType pb.MessageType, isSuccess bool) (protoMessage *pb.LockerMessage, err error) {

	protoMessage = &pb.LockerMessage{
		Message: &pb.LockerMessage_Ack{
			Ack: &pb.ServerAck{
				Id:            id,
				MsgType:       messageType,
				ServerSuccess: isSuccess,
			},
		},
	}

	return protoMessage, err
}

// CreateMessage_FileMeta Create ProtoBuf message.
//  id
//  msgType
//  fileName
//  fileHash
func CreateMessage_FileMeta(id []byte, msgType pb.MessageType, namesSpace string, projectName string, jobID string, fileName string, fileHash []byte) (protoMessage *pb.LockerMessage, err error) {

	protoMessage = &pb.LockerMessage{
		Message: &pb.LockerMessage_Meta{
			Meta: &pb.FileMeta{
				Id:        id,
				MsgType:   msgType,
				Namespace: namesSpace,
				Project:   projectName,
				JobID:     jobID,
				Filename:  fileName,
				Hash:      fileHash,
			},
		},
	}

	return protoMessage, err
}

func CreateMessage_FilePackage(id []byte, msgType pb.MessageType, payload []byte, isFinalPayload bool) (protoMessage *pb.LockerMessage, err error) {

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

// SendProtoBufMessage Generic protobuf sender that accepts an interface to the defined messages
func SendProtoBufMessage(connection net.Conn, message *protobuf.LockerMessage) {
	sizePrefix := make([]byte, 4)
	dataToSend, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}

	// Write Proto Packet Data
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, dataToSend)

	// Prepend 4 bytes of Proto Packet Size
	binary.BigEndian.PutUint32(sizePrefix, uint32(len(buffer.Bytes())))
	connection.Write(sizePrefix)

	// Send Packet
	connection.Write(buffer.Bytes())
}
