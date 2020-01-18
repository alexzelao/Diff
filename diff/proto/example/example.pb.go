// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mytest/proto/example/example.proto

/*
Package cq_srv_test_service is a generated protocol buffer package.

It is generated from these files:
	mytest/proto/example/example.proto

It has these top-level messages:
	Message
	Request
	Response
	DiffRequest
	DiffResponse
	StreamingRequest
	StreamingResponse
	Ping
	Pong
*/
package cq_srv_test_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Message struct {
	Say string `protobuf:"bytes,1,opt,name=say" json:"say,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type DiffRequest struct {
	Original        string `protobuf:"bytes,1,opt,name=original" json:"original,omitempty"`
	Revised         string `protobuf:"bytes,2,opt,name=revised" json:"revised,omitempty"`
	Bincardoriginal string `protobuf:"bytes,3,opt,name=bincardoriginal" json:"bincardoriginal,omitempty"`
	Bincardrevised  string `protobuf:"bytes,4,opt,name=bincardrevised" json:"bincardrevised,omitempty"`
	Locale          string `protobuf:"bytes,5,opt,name=locale" json:"locale,omitempty"`
}

func (m *DiffRequest) Reset()                    { *m = DiffRequest{} }
func (m *DiffRequest) String() string            { return proto.CompactTextString(m) }
func (*DiffRequest) ProtoMessage()               {}
func (*DiffRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DiffRequest) GetOriginal() string {
	if m != nil {
		return m.Original
	}
	return ""
}

func (m *DiffRequest) GetRevised() string {
	if m != nil {
		return m.Revised
	}
	return ""
}

func (m *DiffRequest) GetBincardoriginal() string {
	if m != nil {
		return m.Bincardoriginal
	}
	return ""
}

func (m *DiffRequest) GetBincardrevised() string {
	if m != nil {
		return m.Bincardrevised
	}
	return ""
}

func (m *DiffRequest) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

type DiffResponse struct {
	Symbols string                   `protobuf:"bytes,1,opt,name=symbols" json:"symbols,omitempty"`
	Added   []*DiffResponse_Position `protobuf:"bytes,2,rep,name=added" json:"added,omitempty"`
	Deleted []*DiffResponse_Position `protobuf:"bytes,3,rep,name=deleted" json:"deleted,omitempty"`
}

func (m *DiffResponse) Reset()                    { *m = DiffResponse{} }
func (m *DiffResponse) String() string            { return proto.CompactTextString(m) }
func (*DiffResponse) ProtoMessage()               {}
func (*DiffResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DiffResponse) GetSymbols() string {
	if m != nil {
		return m.Symbols
	}
	return ""
}

func (m *DiffResponse) GetAdded() []*DiffResponse_Position {
	if m != nil {
		return m.Added
	}
	return nil
}

func (m *DiffResponse) GetDeleted() []*DiffResponse_Position {
	if m != nil {
		return m.Deleted
	}
	return nil
}

type DiffResponse_Position struct {
	Pos int64 `protobuf:"varint,1,opt,name=pos" json:"pos,omitempty"`
	Len int64 `protobuf:"varint,2,opt,name=len" json:"len,omitempty"`
}

func (m *DiffResponse_Position) Reset()                    { *m = DiffResponse_Position{} }
func (m *DiffResponse_Position) String() string            { return proto.CompactTextString(m) }
func (*DiffResponse_Position) ProtoMessage()               {}
func (*DiffResponse_Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *DiffResponse_Position) GetPos() int64 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func (m *DiffResponse_Position) GetLen() int64 {
	if m != nil {
		return m.Len
	}
	return 0
}

type StreamingRequest struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *StreamingRequest) Reset()                    { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()               {}
func (*StreamingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *StreamingResponse) Reset()                    { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()               {}
func (*StreamingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke int64 `protobuf:"varint,1,opt,name=stroke" json:"stroke,omitempty"`
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke int64 `protobuf:"varint,1,opt,name=stroke" json:"stroke,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "cq.srv.test.service.Message")
	proto.RegisterType((*Request)(nil), "cq.srv.test.service.Request")
	proto.RegisterType((*Response)(nil), "cq.srv.test.service.Response")
	proto.RegisterType((*DiffRequest)(nil), "cq.srv.test.service.diffRequest")
	proto.RegisterType((*DiffResponse)(nil), "cq.srv.test.service.diffResponse")
	proto.RegisterType((*DiffResponse_Position)(nil), "cq.srv.test.service.diffResponse.Position")
	proto.RegisterType((*StreamingRequest)(nil), "cq.srv.test.service.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "cq.srv.test.service.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "cq.srv.test.service.Ping")
	proto.RegisterType((*Pong)(nil), "cq.srv.test.service.Pong")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Example service

type ExampleClient interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Example_StreamClient, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Example_PingPongClient, error)
	Diff(ctx context.Context, in *DiffRequest, opts ...client.CallOption) (*DiffResponse, error)
}

type exampleClient struct {
	c           client.Client
	serviceName string
}

func NewExampleClient(serviceName string, c client.Client) ExampleClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "cq.srv.test.service"
	}
	return &exampleClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *exampleClient) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Example.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleClient) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Example_StreamClient, error) {
	req := c.c.NewRequest(c.serviceName, "Example.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &exampleStreamClient{stream}, nil
}

type Example_StreamClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type exampleStreamClient struct {
	stream client.Streamer
}

func (x *exampleStreamClient) Close() error {
	return x.stream.Close()
}

func (x *exampleStreamClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *exampleStreamClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *exampleStreamClient) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleClient) PingPong(ctx context.Context, opts ...client.CallOption) (Example_PingPongClient, error) {
	req := c.c.NewRequest(c.serviceName, "Example.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &examplePingPongClient{stream}, nil
}

type Example_PingPongClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type examplePingPongClient struct {
	stream client.Streamer
}

func (x *examplePingPongClient) Close() error {
	return x.stream.Close()
}

func (x *examplePingPongClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *examplePingPongClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *examplePingPongClient) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *examplePingPongClient) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleClient) Diff(ctx context.Context, in *DiffRequest, opts ...client.CallOption) (*DiffResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Example.Diff", in)
	out := new(DiffResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Example_StreamStream) error
	PingPong(context.Context, Example_PingPongStream) error
	Diff(context.Context, *DiffRequest, *DiffResponse) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Example{hdlr}, opts...))
}

type Example struct {
	ExampleHandler
}

func (h *Example) Call(ctx context.Context, in *Request, out *Response) error {
	return h.ExampleHandler.Call(ctx, in, out)
}

func (h *Example) Stream(ctx context.Context, stream server.Streamer) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ExampleHandler.Stream(ctx, m, &exampleStreamStream{stream})
}

type Example_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type exampleStreamStream struct {
	stream server.Streamer
}

func (x *exampleStreamStream) Close() error {
	return x.stream.Close()
}

func (x *exampleStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *exampleStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *exampleStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *Example) PingPong(ctx context.Context, stream server.Streamer) error {
	return h.ExampleHandler.PingPong(ctx, &examplePingPongStream{stream})
}

type Example_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type examplePingPongStream struct {
	stream server.Streamer
}

func (x *examplePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *examplePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *examplePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *examplePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *examplePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *Example) Diff(ctx context.Context, in *DiffRequest, out *DiffResponse) error {
	return h.ExampleHandler.Diff(ctx, in, out)
}

func init() { proto.RegisterFile("mytest/proto/example/example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x6d, 0x96, 0xb4, 0x29, 0x77, 0x08, 0x86, 0x41, 0x28, 0x84, 0x0d, 0x0d, 0x4b, 0x4c, 0x81,
	0x87, 0x6c, 0x1a, 0x3f, 0x80, 0xc4, 0xc6, 0xdb, 0xa4, 0x29, 0x3c, 0xf2, 0xe4, 0x26, 0xb7, 0x91,
	0x85, 0x63, 0x77, 0xb1, 0x57, 0xd1, 0x8f, 0xe2, 0x2f, 0xf8, 0x10, 0x3e, 0x05, 0xd9, 0xb1, 0xa7,
	0xa9, 0x4a, 0x99, 0xf6, 0x94, 0x7b, 0xef, 0x39, 0xe7, 0xea, 0xf8, 0xd8, 0x01, 0xda, 0x6d, 0x0c,
	0x6a, 0x73, 0xba, 0xea, 0x95, 0x51, 0xa7, 0xf8, 0x8b, 0x75, 0x2b, 0x81, 0xe1, 0x5b, 0xba, 0x29,
	0x79, 0x59, 0xdf, 0x94, 0xba, 0x5f, 0x97, 0x96, 0x58, 0x6a, 0xec, 0xd7, 0xbc, 0x46, 0xfa, 0x16,
	0xd2, 0x2b, 0xd4, 0x9a, 0xb5, 0x48, 0x0e, 0x20, 0xd6, 0x6c, 0x93, 0x45, 0xc7, 0x51, 0xf1, 0xa4,
	0xb2, 0x25, 0x3d, 0x82, 0xb4, 0xc2, 0x9b, 0x5b, 0xd4, 0x86, 0x10, 0x48, 0x24, 0xeb, 0xd0, 0xa3,
	0xae, 0xa6, 0x87, 0x30, 0xaf, 0x50, 0xaf, 0x94, 0xd4, 0x4e, 0xdc, 0xe9, 0x36, 0x88, 0x3b, 0xdd,
	0xd2, 0xdf, 0x11, 0xec, 0x37, 0x7c, 0xb9, 0x0c, 0x1b, 0x72, 0x98, 0xab, 0x9e, 0xb7, 0x5c, 0x32,
	0xe1, 0x69, 0x77, 0x3d, 0xc9, 0x20, 0xed, 0x71, 0xcd, 0x35, 0x36, 0xd9, 0x9e, 0x83, 0x42, 0x4b,
	0x0a, 0x78, 0xbe, 0xe0, 0xb2, 0x66, 0x7d, 0x73, 0x27, 0x8e, 0x1d, 0x63, 0x7b, 0x4c, 0x4e, 0xe0,
	0x99, 0x1f, 0x85, 0x55, 0x89, 0x23, 0x6e, 0x4d, 0xc9, 0x6b, 0x98, 0x09, 0x55, 0x33, 0x81, 0xd9,
	0xd4, 0xe1, 0xbe, 0xa3, 0x7f, 0x23, 0x78, 0x3a, 0xf8, 0xf5, 0x47, 0xca, 0x20, 0xd5, 0x9b, 0x6e,
	0xa1, 0x84, 0xf6, 0x7e, 0x43, 0x4b, 0xbe, 0xc0, 0x94, 0x35, 0x8d, 0x33, 0x1b, 0x17, 0xfb, 0xe7,
	0x9f, 0xca, 0x91, 0x64, 0xcb, 0xfb, 0xbb, 0xca, 0x6b, 0xa5, 0xb9, 0xe1, 0x4a, 0x56, 0x83, 0x90,
	0x5c, 0x40, 0xda, 0xa0, 0x40, 0x83, 0x4d, 0x16, 0x3f, 0x7a, 0x47, 0x90, 0xe6, 0x25, 0xcc, 0xc3,
	0xd0, 0x5e, 0xc0, 0x4a, 0x0d, 0x4e, 0xe3, 0xca, 0x96, 0x76, 0x22, 0x50, 0xba, 0x40, 0xe3, 0xca,
	0x96, 0xb4, 0x80, 0x83, 0xef, 0xa6, 0x47, 0xd6, 0x71, 0xd9, 0x86, 0x6b, 0x79, 0x05, 0xd3, 0x5a,
	0xdd, 0x4a, 0xe3, 0x95, 0x43, 0x43, 0x3f, 0xc2, 0x8b, 0x7b, 0x4c, 0x1f, 0xc8, 0x38, 0xf5, 0x1d,
	0x24, 0xd7, 0x5c, 0xb6, 0x36, 0x57, 0x6d, 0x7a, 0xf5, 0x13, 0x3d, 0xec, 0x3b, 0x87, 0xab, 0xdd,
	0xf8, 0xf9, 0x9f, 0x3d, 0x48, 0x2f, 0x87, 0x87, 0x4a, 0x2e, 0x21, 0xf9, 0xca, 0x84, 0x20, 0x87,
	0xa3, 0x69, 0x78, 0xcb, 0xf9, 0xd1, 0x0e, 0x74, 0xb0, 0x49, 0x27, 0xe4, 0x07, 0xcc, 0x06, 0xf7,
	0xe4, 0xc3, 0x28, 0x75, 0x3b, 0x84, 0xfc, 0xe4, 0x21, 0x5a, 0x58, 0x7d, 0x16, 0x91, 0x6f, 0x30,
	0xb7, 0xe7, 0x75, 0x67, 0x7a, 0x33, 0xaa, 0xb3, 0x70, 0xbe, 0x03, 0x52, 0xb2, 0xa5, 0x93, 0x22,
	0x3a, 0x8b, 0xc8, 0x15, 0x24, 0x17, 0x7c, 0xb9, 0x24, 0xc7, 0xff, 0xb9, 0xf9, 0xc1, 0xdd, 0xfb,
	0x07, 0xdf, 0x06, 0x9d, 0x2c, 0x66, 0xee, 0x27, 0xff, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x78,
	0x1e, 0xbf, 0x4d, 0x0a, 0x04, 0x00, 0x00,
}
