// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: stream.proto

/*
Package stream is a generated protocol buffer package.

It is generated from these files:
	stream.proto

It has these top-level messages:
	Request
	Response
*/
package stream

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	chunk "git.profects.com/profects/utils/chunk"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option
var _ chunk.Chunk

// Client API for Streamer service

type StreamerService interface {
	RequestStream(ctx context.Context, opts ...client.CallOption) (Streamer_RequestStreamService, error)
	ServerStream(ctx context.Context, in *Request, opts ...client.CallOption) (Streamer_ServerStreamService, error)
	RequestAndServerStream(ctx context.Context, opts ...client.CallOption) (Streamer_RequestAndServerStreamService, error)
	ChunkRequestStream(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	ChunkServerStream(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	ChunkRequestAndServerStream(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type streamerService struct {
	c    client.Client
	name string
}

func NewStreamerService(name string, c client.Client) StreamerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "streamer"
	}
	return &streamerService{
		c:    c,
		name: name,
	}
}

func (c *streamerService) RequestStream(ctx context.Context, opts ...client.CallOption) (Streamer_RequestStreamService, error) {
	req := c.c.NewRequest(c.name, "Streamer.RequestStream", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &streamerServiceRequestStream{stream}, nil
}

type Streamer_RequestStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Request) error
	Recv() (*Response, error)
}

type streamerServiceRequestStream struct {
	stream client.Stream
}

func (x *streamerServiceRequestStream) Close() error {
	return x.stream.Close()
}

func (x *streamerServiceRequestStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerServiceRequestStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerServiceRequestStream) Send(m *Request) error {
	return x.stream.Send(m)
}

func (x *streamerServiceRequestStream) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamerService) ChunkRequestStream(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Streamer.ChunkRequestStream", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := chunk.Send(stream.Send, in); err != nil {
		return nil, err
	}
	receiveStream := &streamerServiceChunkRequestStream{stream}
	defer receiveStream.Close()
	blob, err := chunk.Receive(receiveStream.RecvMsg)
	if err != nil {
		return nil, err
	}
	var rsp Response
	err = proto.Unmarshal(blob, &rsp)
	return &rsp, err
}

type Streamer_ChunkRequestStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Request) error
	Recv() (*Response, error)
}

type streamerServiceChunkRequestStream struct {
	stream client.Stream
}

func (x *streamerServiceChunkRequestStream) Close() error {
	return x.stream.Close()
}

func (x *streamerServiceChunkRequestStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerServiceChunkRequestStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerServiceChunkRequestStream) Send(m *Request) error {
	return x.stream.Send(m)
}

func (x *streamerServiceChunkRequestStream) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamerService) ServerStream(ctx context.Context, in *Request, opts ...client.CallOption) (Streamer_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "Streamer.ServerStream", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &streamerServiceServerStream{stream}, nil
}

type Streamer_ServerStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Response, error)
}

type streamerServiceServerStream struct {
	stream client.Stream
}

func (x *streamerServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *streamerServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerServiceServerStream) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamerService) ChunkServerStream(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Streamer.ChunkServerStream", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	receiveStream := &streamerServiceChunkServerStream{stream}
	defer receiveStream.Close()
	blob, err := chunk.Receive(receiveStream.RecvMsg)
	if err != nil {
		return nil, err
	}
	var rsp Response
	err = proto.Unmarshal(blob, &rsp)
	return &rsp, err
}

type Streamer_ChunkServerStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Request) error
	Recv() (*Response, error)
}

type streamerServiceChunkServerStream struct {
	stream client.Stream
}

func (x *streamerServiceChunkServerStream) Close() error {
	return x.stream.Close()
}

func (x *streamerServiceChunkServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerServiceChunkServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerServiceChunkServerStream) Send(m *Request) error {
	return x.stream.Send(m)
}

func (x *streamerServiceChunkServerStream) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamerService) RequestAndServerStream(ctx context.Context, opts ...client.CallOption) (Streamer_RequestAndServerStreamService, error) {
	req := c.c.NewRequest(c.name, "Streamer.RequestAndServerStream", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &streamerServiceRequestAndServerStream{stream}, nil
}

type Streamer_RequestAndServerStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Request) error
	Recv() (*Response, error)
}

type streamerServiceRequestAndServerStream struct {
	stream client.Stream
}

func (x *streamerServiceRequestAndServerStream) Close() error {
	return x.stream.Close()
}

func (x *streamerServiceRequestAndServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerServiceRequestAndServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerServiceRequestAndServerStream) Send(m *Request) error {
	return x.stream.Send(m)
}

func (x *streamerServiceRequestAndServerStream) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamerService) ChunkRequestAndServerStream(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Streamer.ChunkRequestAndServerStream", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := chunk.Send(stream.Send, in); err != nil {
		return nil, err
	}
	receiveStream := &streamerServiceChunkRequestAndServerStream{stream}
	defer receiveStream.Close()
	blob, err := chunk.Receive(receiveStream.RecvMsg)
	if err != nil {
		return nil, err
	}
	var rsp Response
	err = proto.Unmarshal(blob, &rsp)
	return &rsp, err
}

type Streamer_ChunkRequestAndServerStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Request) error
	Recv() (*Response, error)
}

type streamerServiceChunkRequestAndServerStream struct {
	stream client.Stream
}

func (x *streamerServiceChunkRequestAndServerStream) Close() error {
	return x.stream.Close()
}

func (x *streamerServiceChunkRequestAndServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerServiceChunkRequestAndServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerServiceChunkRequestAndServerStream) Send(m *Request) error {
	return x.stream.Send(m)
}

func (x *streamerServiceChunkRequestAndServerStream) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Streamer service

type StreamerHandler interface {
	RequestStream(context.Context, Streamer_RequestStreamStream) error
	ServerStream(context.Context, *Request, Streamer_ServerStreamStream) error
	RequestAndServerStream(context.Context, Streamer_RequestAndServerStreamStream) error
	ChunkRequestStream(context.Context, *Request, *Response) error
	ChunkServerStream(context.Context, *Request, *Response) error
	ChunkRequestAndServerStream(context.Context, *Request, *Response) error
}

func RegisterStreamerHandler(s server.Server, hdlr StreamerHandler, opts ...server.HandlerOption) error {
	type streamer interface {
		RequestStream(ctx context.Context, stream server.Stream) error
		ChunkRequestStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		ChunkServerStream(ctx context.Context, stream server.Stream) error
		RequestAndServerStream(ctx context.Context, stream server.Stream) error
		ChunkRequestAndServerStream(ctx context.Context, stream server.Stream) error
	}
	type Streamer struct {
		streamer
	}
	h := &streamerHandler{hdlr}
	return s.Handle(s.NewHandler(&Streamer{h}, opts...))
}

type streamerHandler struct {
	StreamerHandler
}

func (h *streamerHandler) RequestStream(ctx context.Context, stream server.Stream) error {
	return h.StreamerHandler.RequestStream(ctx, &streamerRequestStreamStream{stream})
}

type Streamer_RequestStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
	Recv() (*Request, error)
}

type streamerRequestStreamStream struct {
	stream server.Stream
}

func (x *streamerRequestStreamStream) Close() error {
	return x.stream.Close()
}

func (x *streamerRequestStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerRequestStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerRequestStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (x *streamerRequestStreamStream) Recv() (*Request, error) {
	m := new(Request)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *streamerHandler) ChunkRequestStream(ctx context.Context, stream server.Stream) error {
	defer stream.Close()
	blob, err := chunk.Receive(stream.Recv)
	if err != nil {
		return err
	}
	m := new(Request)
	if err = proto.Unmarshal(blob, m); err != nil {
		return err
	}
	rsp := new(Response)
	if err := h.StreamerHandler.ChunkRequestStream(ctx, m, rsp); err != nil {
		return err
	}
	return chunk.Send(stream.Send, rsp)
}

type Streamer_ChunkRequestStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
	Recv() (*Request, error)
}

type streamerChunkRequestStreamStream struct {
	stream server.Stream
}

func (x *streamerChunkRequestStreamStream) Close() error {
	return x.stream.Close()
}

func (x *streamerChunkRequestStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerChunkRequestStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerChunkRequestStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (x *streamerChunkRequestStreamStream) Recv() (*Request, error) {
	m := new(Request)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *streamerHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(Request)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.StreamerHandler.ServerStream(ctx, m, &streamerServerStreamStream{stream})
}

type Streamer_ServerStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
}

type streamerServerStreamStream struct {
	stream server.Stream
}

func (x *streamerServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *streamerServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerServerStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (h *streamerHandler) ChunkServerStream(ctx context.Context, stream server.Stream) error {
	defer stream.Close()
	m := new(Request)
	if err := stream.Recv(m); err != nil {
		return err
	}
	rsp := new(Response)
	if err := h.StreamerHandler.ChunkServerStream(ctx, m, rsp); err != nil {
		return err
	}
	return chunk.Send(stream.Send, rsp)
}

type Streamer_ChunkServerStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
}

type streamerChunkServerStreamStream struct {
	stream server.Stream
}

func (x *streamerChunkServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *streamerChunkServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerChunkServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerChunkServerStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (h *streamerHandler) RequestAndServerStream(ctx context.Context, stream server.Stream) error {
	return h.StreamerHandler.RequestAndServerStream(ctx, &streamerRequestAndServerStreamStream{stream})
}

type Streamer_RequestAndServerStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
	Recv() (*Request, error)
}

type streamerRequestAndServerStreamStream struct {
	stream server.Stream
}

func (x *streamerRequestAndServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *streamerRequestAndServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerRequestAndServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerRequestAndServerStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (x *streamerRequestAndServerStreamStream) Recv() (*Request, error) {
	m := new(Request)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *streamerHandler) ChunkRequestAndServerStream(ctx context.Context, stream server.Stream) error {
	defer stream.Close()
	blob, err := chunk.Receive(stream.Recv)
	if err != nil {
		return err
	}
	m := new(Request)
	if err = proto.Unmarshal(blob, m); err != nil {
		return err
	}
	rsp := new(Response)
	if err := h.StreamerHandler.ChunkRequestAndServerStream(ctx, m, rsp); err != nil {
		return err
	}
	return chunk.Send(stream.Send, rsp)
}

type Streamer_ChunkRequestAndServerStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
	Recv() (*Request, error)
}

type streamerChunkRequestAndServerStreamStream struct {
	stream server.Stream
}

func (x *streamerChunkRequestAndServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *streamerChunkRequestAndServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerChunkRequestAndServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerChunkRequestAndServerStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (x *streamerChunkRequestAndServerStreamStream) Recv() (*Request, error) {
	m := new(Request)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
