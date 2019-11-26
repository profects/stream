package main

import (
	"context"
	"errors"
	"log"
	"runtime"
	"time"

	proto "github.com/micro/examples/stream/server/proto"
	"github.com/micro/go-micro"
)

type Streamer struct{}

// Server side stream
func (e *Streamer) ServerStream(ctx context.Context, req *proto.Request, stream proto.Streamer_ServerStreamStream) error {
	// log.Printf("Got msg %d", req.Count)
	for i := 0; i < int(req.Count); i++ {
		if err := stream.Send(&proto.Response{Count: int64(i)}); err != nil {
			return err
		}
		return errors.New("ERROR")
	}
	return nil
}

// Bidirectional stream
func (e *Streamer) Stream(ctx context.Context, stream proto.Streamer_StreamStream) error {
	return nil
}

func tick() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		log.Println("Amount of goroutines: ", runtime.NumGoroutine())
	}
}

func main() {
	// new service
	service := micro.NewService(
		micro.Name("go.micro.srv.stream"),
	)

	// Init command line
	service.Init()

	// Register Handler
	proto.RegisterStreamerHandler(service.Server(), new(Streamer))

	go tick()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
