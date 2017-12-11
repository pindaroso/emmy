// Code generated by protoc-gen-go.
// source: services.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Protocol service

type ProtocolClient interface {
	Run(ctx context.Context, opts ...grpc.CallOption) (Protocol_RunClient, error)
}

type protocolClient struct {
	cc *grpc.ClientConn
}

func NewProtocolClient(cc *grpc.ClientConn) ProtocolClient {
	return &protocolClient{cc}
}

func (c *protocolClient) Run(ctx context.Context, opts ...grpc.CallOption) (Protocol_RunClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Protocol_serviceDesc.Streams[0], c.cc, "/protobuf.Protocol/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &protocolRunClient{stream}
	return x, nil
}

type Protocol_RunClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type protocolRunClient struct {
	grpc.ClientStream
}

func (x *protocolRunClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *protocolRunClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Protocol service

type ProtocolServer interface {
	Run(Protocol_RunServer) error
}

func RegisterProtocolServer(s *grpc.Server, srv ProtocolServer) {
	s.RegisterService(&_Protocol_serviceDesc, srv)
}

func _Protocol_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProtocolServer).Run(&protocolRunServer{stream})
}

type Protocol_RunServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type protocolRunServer struct {
	grpc.ServerStream
}

func (x *protocolRunServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *protocolRunServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Protocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Protocol",
	HandlerType: (*ProtocolServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _Protocol_Run_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}

// Client API for Info service

type InfoClient interface {
	GetServiceInfo(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*ServiceInfo, error)
}

type infoClient struct {
	cc *grpc.ClientConn
}

func NewInfoClient(cc *grpc.ClientConn) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) GetServiceInfo(ctx context.Context, in *EmptyMsg, opts ...grpc.CallOption) (*ServiceInfo, error) {
	out := new(ServiceInfo)
	err := grpc.Invoke(ctx, "/protobuf.Info/GetServiceInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Info service

type InfoServer interface {
	GetServiceInfo(context.Context, *EmptyMsg) (*ServiceInfo, error)
}

func RegisterInfoServer(s *grpc.Server, srv InfoServer) {
	s.RegisterService(&_Info_serviceDesc, srv)
}

func _Info_GetServiceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).GetServiceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Info/GetServiceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).GetServiceInfo(ctx, req.(*EmptyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Info_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceInfo",
			Handler:    _Info_GetServiceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

func init() { proto.RegisterFile("services.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5,
	0x69, 0x52, 0x7c, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x19, 0x23, 0x5b, 0x2e, 0x8e, 0x00,
	0x10, 0x23, 0x39, 0x3f, 0x47, 0xc8, 0x90, 0x8b, 0x39, 0xa8, 0x34, 0x4f, 0x48, 0x50, 0x0f, 0xa6,
	0x5a, 0xcf, 0x17, 0xa2, 0x58, 0x0a, 0x53, 0x48, 0x89, 0x41, 0x83, 0xd1, 0x80, 0xd1, 0xc8, 0x95,
	0x8b, 0xc5, 0x33, 0x2f, 0x2d, 0x5f, 0xc8, 0x96, 0x8b, 0xcf, 0x3d, 0xb5, 0x24, 0x18, 0x62, 0x2b,
	0x58, 0x44, 0x08, 0xa1, 0xc5, 0x35, 0xb7, 0xa0, 0xa4, 0xd2, 0xb7, 0x38, 0x5d, 0x4a, 0x14, 0x21,
	0x86, 0xa4, 0x54, 0x89, 0x21, 0x89, 0x0d, 0x2c, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0x83, 0xf8, 0x71, 0xb8, 0x00, 0x00, 0x00,
}