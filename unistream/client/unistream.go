package main

import (
	"context"
	"io"
	"log"

	pb "swaammi.com/go-grpc/unistream/proto"
)

func doUnistream(c pb.UnistreamServiceClient) {
	log.Println("doUnistream was invoked")

	req := &pb.UniStreamRequest{}

	stream, err := c.SendMessage(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling SendMessage: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("SendMessage: %s\n", msg.Response)
	}

}
