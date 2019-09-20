package main

import (
	"context"
	"fmt"

	proto "github.com/micro/examples/stream/server/proto"
	"github.com/micro/go-micro"
)

func serverStream(cl proto.StreamerService) {
	// send request to stream count of 10
	stream, err := cl.ServerStream(context.Background(), &proto.Request{Count: int64(10)})
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// server side stream
	// receive messages for a 10 count
	for j := 0; j < 10; j++ {
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println("recv err", err)
			return
		}
		fmt.Printf("got msg %v\n", rsp.Count)
	}

	// close the stream
	if err := stream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
}

func main() {
	service := micro.NewService()
	service.Init()

	// create client
	cl := proto.NewStreamerService("go.micro.srv.stream", service.Client())

	// server side stream
	for i := 0; i < 1000; i++ {
		serverStream(cl)
	}
}
