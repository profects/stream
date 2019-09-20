package main

import (
	"context"
	"fmt"
	"time"

	proto "github.com/micro/examples/stream/server/proto"
	"github.com/micro/go-micro"
)

func serverStream(cl proto.StreamerService, streamIndex int64) {
	// send request to stream count of 10
	stream, err := cl.ServerStream(context.Background(), &proto.Request{Count: streamIndex})
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// server side stream
	// receive messages for a 10 count
	for j := 0; j < 40; j++ {
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println("recv err", err)
			return
		}
		fmt.Printf("got msg %v => %v\n", j, rsp.Count)
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

	// bidirectional stream
	//bidirectional(cl)

	// server side stream
	for i := 0; i < 1000; i++ {
		time.Sleep((100 * time.Millisecond))
		serverStream(cl, int64(i))
	}
}
