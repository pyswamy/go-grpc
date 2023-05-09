package main

import (
	"encoding/json"
	"log"

	pb "swaammi.com/go-grpc/unistream/proto"
)

// Define System struct
type System struct {
	IpAddress string
	Port      int32
	HostName  string
}

func (S *Server) SendMessage(in *pb.UniStreamRequest, stream pb.UnistreamService_SendMessageServer) error {
	log.Printf("SendMessage function invoked with: %v\n", in)
	// Create a new instance of System struct
	sysStruct := &System{
		IpAddress: "192.168.1.1",
		Port:      8080,
		HostName:  "localhost",
	}

	// Convert your struct to a byte array
	structBytes, err := json.Marshal(sysStruct)
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		for _, b := range structBytes {
			err := stream.Send(&pb.UniStreamResponse{
				Response: []byte{b},
			})
			if err != nil {
				return err
			}
		}

	}

	return nil
}
