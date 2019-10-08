package main

import (
	"context"
	"go-RPC/go_client/proto/hello"
	"io"
	"log"

	"google.golang.org/grpc"
)

const Address = "127.0.0.1:9090" // gRPC 服务地址

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := hello.NewHelloClient(conn)

	// 调用 SayHello 方法
	res, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "Hello World"})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Message)

	stream, err := c.LotsOfReplies(context.Background(), &hello.HelloRequest{Name: "Hello World"})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("stream.Recv: %v", err)
		}

		log.Println("%s", res.Message)
	}

}
