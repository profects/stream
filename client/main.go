package main

import (
	"context"
	"log"
	"runtime"
	"time"

	proto "github.com/micro/examples/stream/server/proto"
	"github.com/micro/go-micro"
)

func serverStream(cl proto.StreamerService) {
	// send request to stream count of 10
	stream, err := cl.ServerStream(context.Background(), &proto.Request{Count: int64(10)})
	if err != nil {
		//log.Println("err in client:", err)
		return
	}
	defer stream.Close()

	// server side stream
	// receive messages for a 10 count
	for j := 0; j < 10; j++ {
		_, err = stream.Recv()
		if err != nil {
			//log.Println("recv err", err)
			return
		}
		//log.Printf("got msg %v\n", rsp.Count)
	}
}

func tick() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		log.Println("Amount of goroutines: ", runtime.NumGoroutine())
	}
}

func main() {
	service := micro.NewService()
	service.Init()

	// create client
	cl := proto.NewStreamerService("go.micro.srv.stream", service.Client())

	go tick()

	// server side stream
	for i := 0; i < 100000; i++ {
		serverStream(cl)
	}
}
