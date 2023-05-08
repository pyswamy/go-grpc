package main

import (
	"encoding/json"
	"fmt"
	"log"

	pb "swaammi.com/go-grpc/unistream/proto"
)

func (S *Server) SendMessage(in *pb.UniStreamRequest, stream pb.UnistreamService_SendMessageServer) error {
	log.Printf("SendMessage function invoked with: %v\n", in)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.Data, i)
		byteData, _ := json.Marshal(res)

		stream.Send(&pb.UniStreamResponse{
			Response: byteData,
		})
	}
	return nil
}
