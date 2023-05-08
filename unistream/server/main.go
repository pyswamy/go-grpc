package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "swaammi.com/go-grpc/unistream/proto"
)

type Server struct {
	pb.UnistreamServiceServer
}

var addr string = "0.0.0.0:50055"

func main() {

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterUnistreamServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
