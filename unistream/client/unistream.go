package main

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"

	pb "swaammi.com/go-grpc/unistream/proto"
)

func doUnistream(c pb.UnistreamServiceClient) {
	log.Println("doUnistream was invoked")

	req := &pb.UniStreamRequest{}

	// Make the gRPC call
	stream, err := c.SendMessage(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling SendMessage: %v\n", err)
	}

	// Create a buffer to hold the received bytes
	var buffer bytes.Buffer

	// Read the bytes from the response stream and write them to the buffer
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		_, err = buffer.Write(msg.Response)
		if err != nil {
			log.Fatalf("Failed to write response to buffer: %v", err)
		}

		log.Printf("SendMessage: %s\n", msg.Response)
	}

	// Write the buffer contents to a JSON file
	err = ioutil.WriteFile("output.json", buffer.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

}
