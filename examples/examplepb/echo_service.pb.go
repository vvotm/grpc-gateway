// Code generated by protoc-gen-go.
// source: examples/examplepb/echo_service.proto
// DO NOT EDIT!

/*
Package examplepb is a generated protocol buffer package.

Echo Service

Echo Service API consists of a single service which returns
a message.

It is generated from these files:
	examples/examplepb/echo_service.proto
	examples/examplepb/a_bit_of_everything.proto
	examples/examplepb/stream.proto
	examples/examplepb/flow_combination.proto

It has these top-level messages:
	SimpleMessage
	ABitOfEverything
	EmptyProto
	NonEmptyProto
	UnaryProto
	NestedProto
	SingleNestedProto
*/
package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// SimpleMessage represents a simple message sent to the Echo service.
type SimpleMessage struct {
	// Id represents the message identifier.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *SimpleMessage) Reset()                    { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string            { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()               {}
func (*SimpleMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*SimpleMessage)(nil), "gengo.grpc.gateway.examples.examplepb.SimpleMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for EchoService service

type EchoServiceClient interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.EchoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.EchoService/EchoBody", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(context.Context, *SimpleMessage) (*SimpleMessage, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengo.grpc.gateway.examples.examplepb.EchoService/EchoBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoBody(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gengo.grpc.gateway.examples.examplepb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
		{
			MethodName: "EchoBody",
			Handler:    _EchoService_EchoBody_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("examples/examplepb/echo_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x87, 0x32, 0x0a, 0x92, 0xf4, 0x53, 0x93, 0x33, 0xf2, 0xe3, 0x8b,
	0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0xd3, 0x53,
	0xf3, 0xd2, 0xf3, 0xf5, 0xd2, 0x8b, 0x0a, 0x92, 0xf5, 0xd2, 0x13, 0x4b, 0x52, 0xcb, 0x13, 0x2b,
	0xf5, 0x60, 0x3a, 0xf5, 0xe0, 0x3a, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13,
	0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0x86,
	0x28, 0xc9, 0x73, 0xf1, 0x06, 0x67, 0x82, 0x54, 0xfa, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a,
	0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x01, 0x59, 0x46, 0x7b,
	0x98, 0xb8, 0xb8, 0x5d, 0x81, 0x96, 0x07, 0x43, 0xec, 0x16, 0x9a, 0xc8, 0xc8, 0xc5, 0x02, 0xe2,
	0x0b, 0x99, 0xe8, 0x11, 0x65, 0xbf, 0x1e, 0x8a, 0xf1, 0x52, 0x64, 0xe9, 0x52, 0x92, 0x6d, 0xba,
	0xfc, 0x64, 0x32, 0x93, 0xb8, 0x92, 0xa8, 0x7e, 0x99, 0x21, 0x2c, 0x50, 0xc0, 0x41, 0xa2, 0x5f,
	0x9d, 0x99, 0x52, 0x2b, 0x34, 0x83, 0x91, 0x8b, 0x03, 0xe4, 0x26, 0xa7, 0xfc, 0x94, 0x4a, 0xba,
	0xba, 0x4b, 0x01, 0xec, 0x2e, 0x29, 0x4c, 0x77, 0xc5, 0x27, 0x01, 0x9d, 0x62, 0xc5, 0xa8, 0xe5,
	0xc4, 0x1d, 0xc5, 0x09, 0xd7, 0x9c, 0xc4, 0x06, 0x0e, 0x73, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2a, 0x5f, 0x67, 0x79, 0xe1, 0x01, 0x00, 0x00,
}
