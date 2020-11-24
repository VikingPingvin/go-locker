package server

import (
	pb "vikingPingvin/locker/server/protobuf"
)

func CreateMessage_ServerACk(id int32, messageType pb.MessageType, isSuccess bool) (protoMessage *pb.LockerMessage, err error) {

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
func CreateMessage_FileMeta(id int32, msgType pb.MessageType, namesSpace string, projectName string, fileName string, fileHash []byte) (protoMessage *pb.LockerMessage, err error) {

	protoMessage = &pb.LockerMessage{
		Message: &pb.LockerMessage_Meta{
			Meta: &pb.FileMeta{
				Id:        id,
				MsgType:   msgType,
				Namespace: namesSpace,
				Project:   projectName,
				Filename:  fileName,
				Hash:      fileHash,
			},
		},
	}

	return protoMessage, err
}

func CreateMessage_FilePackage(id int32, msgType pb.MessageType, payload []byte, isFinalPayload bool) (protoMessage *pb.LockerMessage, err error) {

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
