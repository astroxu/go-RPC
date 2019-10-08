package main

import (
	"go-RPC/go_server/controller/hello_controller"
	"go-RPC/go_server/proto/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

const Address = "127.0.0.1:9090"

func main() {
	lis, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	hello.RegisterHelloServer(s, &hello_controller.HelloController{})

	log.Println("Listen on " + Address)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
