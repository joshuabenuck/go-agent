// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sampleapp.proto

package sampleapp

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ae74b4e52ac4e0, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
}

func init() { proto.RegisterFile("sampleapp.proto", fileDescriptor_38ae74b4e52ac4e0) }

var fileDescriptor_38ae74b4e52ac4e0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4e, 0xcc, 0x2d,
	0xc8, 0x49, 0x4d, 0x2c, 0x28, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe5, 0x62, 0xf7,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe2, 0x62, 0x29, 0x49, 0xad, 0x28, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x8d, 0xb6, 0x33, 0x72, 0x09, 0x06, 0x83, 0xb5, 0x38,
	0x16, 0x14, 0xe4, 0x64, 0x26, 0x27, 0x96, 0x64, 0xe6, 0xe7, 0x09, 0xa9, 0x70, 0xf1, 0xb8, 0xe4,
	0x87, 0xe6, 0x25, 0x16, 0x55, 0x82, 0x09, 0x21, 0x0e, 0x3d, 0xa8, 0x19, 0x52, 0x70, 0x96, 0x12,
	0x83, 0x90, 0x3a, 0x17, 0x2f, 0x54, 0x55, 0x70, 0x49, 0x51, 0x6a, 0x62, 0x2e, 0x76, 0x65, 0x06,
	0x8c, 0x10, 0x85, 0x10, 0x35, 0x78, 0xcc, 0xd3, 0x60, 0x14, 0xd2, 0xe2, 0xe2, 0x83, 0x29, 0xc4,
	0x67, 0xa4, 0x06, 0xa3, 0x01, 0x63, 0x12, 0x1b, 0xd8, 0x7f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe8, 0x8b, 0x56, 0x80, 0xf2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SampleApplicationClient is the client API for SampleApplication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SampleApplicationClient interface {
	DoUnaryUnary(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	DoUnaryStream(ctx context.Context, in *Message, opts ...grpc.CallOption) (SampleApplication_DoUnaryStreamClient, error)
	DoStreamUnary(ctx context.Context, opts ...grpc.CallOption) (SampleApplication_DoStreamUnaryClient, error)
	DoStreamStream(ctx context.Context, opts ...grpc.CallOption) (SampleApplication_DoStreamStreamClient, error)
}

type sampleApplicationClient struct {
	cc *grpc.ClientConn
}

func NewSampleApplicationClient(cc *grpc.ClientConn) SampleApplicationClient {
	return &sampleApplicationClient{cc}
}

func (c *sampleApplicationClient) DoUnaryUnary(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/SampleApplication/DoUnaryUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleApplicationClient) DoUnaryStream(ctx context.Context, in *Message, opts ...grpc.CallOption) (SampleApplication_DoUnaryStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SampleApplication_serviceDesc.Streams[0], "/SampleApplication/DoUnaryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleApplicationDoUnaryStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SampleApplication_DoUnaryStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type sampleApplicationDoUnaryStreamClient struct {
	grpc.ClientStream
}

func (x *sampleApplicationDoUnaryStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sampleApplicationClient) DoStreamUnary(ctx context.Context, opts ...grpc.CallOption) (SampleApplication_DoStreamUnaryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SampleApplication_serviceDesc.Streams[1], "/SampleApplication/DoStreamUnary", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleApplicationDoStreamUnaryClient{stream}
	return x, nil
}

type SampleApplication_DoStreamUnaryClient interface {
	Send(*Message) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type sampleApplicationDoStreamUnaryClient struct {
	grpc.ClientStream
}

func (x *sampleApplicationDoStreamUnaryClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sampleApplicationDoStreamUnaryClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sampleApplicationClient) DoStreamStream(ctx context.Context, opts ...grpc.CallOption) (SampleApplication_DoStreamStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SampleApplication_serviceDesc.Streams[2], "/SampleApplication/DoStreamStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleApplicationDoStreamStreamClient{stream}
	return x, nil
}

type SampleApplication_DoStreamStreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type sampleApplicationDoStreamStreamClient struct {
	grpc.ClientStream
}

func (x *sampleApplicationDoStreamStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sampleApplicationDoStreamStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SampleApplicationServer is the server API for SampleApplication service.
type SampleApplicationServer interface {
	DoUnaryUnary(context.Context, *Message) (*Message, error)
	DoUnaryStream(*Message, SampleApplication_DoUnaryStreamServer) error
	DoStreamUnary(SampleApplication_DoStreamUnaryServer) error
	DoStreamStream(SampleApplication_DoStreamStreamServer) error
}

// UnimplementedSampleApplicationServer can be embedded to have forward compatible implementations.
type UnimplementedSampleApplicationServer struct {
}

func (*UnimplementedSampleApplicationServer) DoUnaryUnary(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoUnaryUnary not implemented")
}
func (*UnimplementedSampleApplicationServer) DoUnaryStream(req *Message, srv SampleApplication_DoUnaryStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DoUnaryStream not implemented")
}
func (*UnimplementedSampleApplicationServer) DoStreamUnary(srv SampleApplication_DoStreamUnaryServer) error {
	return status.Errorf(codes.Unimplemented, "method DoStreamUnary not implemented")
}
func (*UnimplementedSampleApplicationServer) DoStreamStream(srv SampleApplication_DoStreamStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DoStreamStream not implemented")
}

func RegisterSampleApplicationServer(s *grpc.Server, srv SampleApplicationServer) {
	s.RegisterService(&_SampleApplication_serviceDesc, srv)
}

func _SampleApplication_DoUnaryUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleApplicationServer).DoUnaryUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SampleApplication/DoUnaryUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleApplicationServer).DoUnaryUnary(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleApplication_DoUnaryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SampleApplicationServer).DoUnaryStream(m, &sampleApplicationDoUnaryStreamServer{stream})
}

type SampleApplication_DoUnaryStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type sampleApplicationDoUnaryStreamServer struct {
	grpc.ServerStream
}

func (x *sampleApplicationDoUnaryStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _SampleApplication_DoStreamUnary_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SampleApplicationServer).DoStreamUnary(&sampleApplicationDoStreamUnaryServer{stream})
}

type SampleApplication_DoStreamUnaryServer interface {
	SendAndClose(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type sampleApplicationDoStreamUnaryServer struct {
	grpc.ServerStream
}

func (x *sampleApplicationDoStreamUnaryServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sampleApplicationDoStreamUnaryServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SampleApplication_DoStreamStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SampleApplicationServer).DoStreamStream(&sampleApplicationDoStreamStreamServer{stream})
}

type SampleApplication_DoStreamStreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type sampleApplicationDoStreamStreamServer struct {
	grpc.ServerStream
}

func (x *sampleApplicationDoStreamStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sampleApplicationDoStreamStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SampleApplication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SampleApplication",
	HandlerType: (*SampleApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoUnaryUnary",
			Handler:    _SampleApplication_DoUnaryUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoUnaryStream",
			Handler:       _SampleApplication_DoUnaryStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DoStreamUnary",
			Handler:       _SampleApplication_DoStreamUnary_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DoStreamStream",
			Handler:       _SampleApplication_DoStreamStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sampleapp.proto",
}
