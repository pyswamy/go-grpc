package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "swaammi.com/go-grpc/unistream/proto"
)

var addr string = "localhost:50055"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Failed to connect : %v\n", err)
	}
	defer conn.Close()

	c := pb.NewUnistreamServiceClient(conn)

	defer conn.Close()

	stream, err := c.SendMessage(context.Background(), &pb.UniStreamRequest{})
	if err != nil {
		log.Fatalf("error opening stream: %v", err)
	}
	fmt.Println("stream: ")
	// read the stream of StreamResponse messages and concatenate the bytes data
	var buf bytes.Buffer
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error reading stream: %v", err)
		}
		fmt.Println("stream final data: ", string(res.Response))
		buf.Write(res.Response)
	}

}
