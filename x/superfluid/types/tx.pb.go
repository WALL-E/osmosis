// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/superfluid/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	_ "google.golang.org/protobuf/types/known/durationpb"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("osmosis/superfluid/tx.proto", fileDescriptor_55b645f187d22814) }

var fileDescriptor_55b645f187d22814 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0x2f, 0x2e, 0x2d, 0x48, 0x2d, 0x4a, 0xcb, 0x29, 0xcd, 0x4c, 0xd1, 0x2f,
	0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0x4a, 0xea, 0x21, 0x24, 0xa5, 0x44,
	0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xd2, 0xfa, 0x20, 0x16, 0x44, 0xa5, 0x94, 0x5c, 0x7a, 0x7e, 0x7e,
	0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x52, 0x5a, 0x94, 0x58, 0x92, 0x99,
	0x9f, 0x07, 0x93, 0x4f, 0x06, 0x1b, 0xa5, 0x9f, 0x94, 0x58, 0x9c, 0xaa, 0x5f, 0x66, 0x98, 0x94,
	0x5a, 0x92, 0x68, 0xa8, 0x9f, 0x9c, 0x9f, 0x09, 0x93, 0x57, 0xc6, 0xe2, 0x0c, 0x04, 0x13, 0xa2,
	0xc8, 0x88, 0x95, 0x8b, 0xd9, 0xb7, 0x38, 0xdd, 0xc9, 0xe7, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b,
	0x8f, 0xe5, 0x18, 0xa2, 0x8c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0xa1, 0x06, 0xea, 0xe6, 0x24, 0x26, 0x15, 0xc3, 0x38, 0xfa, 0x15, 0x28, 0xde, 0xac, 0x2c, 0x48,
	0x2d, 0x4e, 0x62, 0x03, 0x9b, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x4d, 0x77, 0x81,
	0x09, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.superfluid.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "osmosis/superfluid/tx.proto",
}
