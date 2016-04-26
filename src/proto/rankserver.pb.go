// Code generated by protoc-gen-go.
// source: rankserver.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	rankserver.proto

It has these top-level messages:
	Ranking
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto1.ProtoPackageIsVersion1

type Ranking struct {
}

func (m *Ranking) Reset()                    { *m = Ranking{} }
func (m *Ranking) String() string            { return proto1.CompactTextString(m) }
func (*Ranking) ProtoMessage()               {}
func (*Ranking) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Ranking_Nil struct {
}

func (m *Ranking_Nil) Reset()                    { *m = Ranking_Nil{} }
func (m *Ranking_Nil) String() string            { return proto1.CompactTextString(m) }
func (*Ranking_Nil) ProtoMessage()               {}
func (*Ranking_Nil) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Ranking_SetId struct {
	SetId uint64 `protobuf:"varint,1,opt,name=SetId" json:"SetId,omitempty"`
}

func (m *Ranking_SetId) Reset()                    { *m = Ranking_SetId{} }
func (m *Ranking_SetId) String() string            { return proto1.CompactTextString(m) }
func (*Ranking_SetId) ProtoMessage()               {}
func (*Ranking_SetId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Ranking_DeleteUserRequest struct {
	SetId  uint64 `protobuf:"varint,1,opt,name=SetId" json:"SetId,omitempty"`
	UserId int32  `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *Ranking_DeleteUserRequest) Reset()                    { *m = Ranking_DeleteUserRequest{} }
func (m *Ranking_DeleteUserRequest) String() string            { return proto1.CompactTextString(m) }
func (*Ranking_DeleteUserRequest) ProtoMessage()               {}
func (*Ranking_DeleteUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

type Ranking_Change struct {
	UserId int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	Score  int32  `protobuf:"varint,2,opt,name=Score" json:"Score,omitempty"`
	SetId  uint64 `protobuf:"varint,3,opt,name=SetId" json:"SetId,omitempty"`
}

func (m *Ranking_Change) Reset()                    { *m = Ranking_Change{} }
func (m *Ranking_Change) String() string            { return proto1.CompactTextString(m) }
func (*Ranking_Change) ProtoMessage()               {}
func (*Ranking_Change) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

type Ranking_Range struct {
	A     int32  `protobuf:"varint,1,opt,name=A" json:"A,omitempty"`
	B     int32  `protobuf:"varint,2,opt,name=B" json:"B,omitempty"`
	SetId uint64 `protobuf:"varint,3,opt,name=SetId" json:"SetId,omitempty"`
}

func (m *Ranking_Range) Reset()                    { *m = Ranking_Range{} }
func (m *Ranking_Range) String() string            { return proto1.CompactTextString(m) }
func (*Ranking_Range) ProtoMessage()               {}
func (*Ranking_Range) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 4} }

type Ranking_RankList struct {
	UserIds []int32 `protobuf:"varint,1,rep,packed,name=UserIds" json:"UserIds,omitempty"`
	Scores  []int32 `protobuf:"varint,2,rep,packed,name=Scores" json:"Scores,omitempty"`
}

func (m *Ranking_RankList) Reset()                    { *m = Ranking_RankList{} }
func (m *Ranking_RankList) String() string            { return proto1.CompactTextString(m) }
func (*Ranking_RankList) ProtoMessage()               {}
func (*Ranking_RankList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 5} }

type Ranking_Users struct {
	UserIds []int32 `protobuf:"varint,1,rep,packed,name=UserIds" json:"UserIds,omitempty"`
	SetId   uint64  `protobuf:"varint,2,opt,name=SetId" json:"SetId,omitempty"`
}

func (m *Ranking_Users) Reset()                    { *m = Ranking_Users{} }
func (m *Ranking_Users) String() string            { return proto1.CompactTextString(m) }
func (*Ranking_Users) ProtoMessage()               {}
func (*Ranking_Users) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 6} }

type Ranking_UserList struct {
	Ranks  []int32 `protobuf:"varint,1,rep,packed,name=Ranks" json:"Ranks,omitempty"`
	Scores []int32 `protobuf:"varint,2,rep,packed,name=Scores" json:"Scores,omitempty"`
}

func (m *Ranking_UserList) Reset()                    { *m = Ranking_UserList{} }
func (m *Ranking_UserList) String() string            { return proto1.CompactTextString(m) }
func (*Ranking_UserList) ProtoMessage()               {}
func (*Ranking_UserList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 7} }

func init() {
	proto1.RegisterType((*Ranking)(nil), "proto.Ranking")
	proto1.RegisterType((*Ranking_Nil)(nil), "proto.Ranking.Nil")
	proto1.RegisterType((*Ranking_SetId)(nil), "proto.Ranking.SetId")
	proto1.RegisterType((*Ranking_DeleteUserRequest)(nil), "proto.Ranking.DeleteUserRequest")
	proto1.RegisterType((*Ranking_Change)(nil), "proto.Ranking.Change")
	proto1.RegisterType((*Ranking_Range)(nil), "proto.Ranking.Range")
	proto1.RegisterType((*Ranking_RankList)(nil), "proto.Ranking.RankList")
	proto1.RegisterType((*Ranking_Users)(nil), "proto.Ranking.Users")
	proto1.RegisterType((*Ranking_UserList)(nil), "proto.Ranking.UserList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for RankingService service

type RankingServiceClient interface {
	RankChange(ctx context.Context, in *Ranking_Change, opts ...grpc.CallOption) (*Ranking_Nil, error)
	DeleteSet(ctx context.Context, in *Ranking_SetId, opts ...grpc.CallOption) (*Ranking_Nil, error)
	DeleteUser(ctx context.Context, in *Ranking_DeleteUserRequest, opts ...grpc.CallOption) (*Ranking_Nil, error)
	QueryRankRange(ctx context.Context, in *Ranking_Range, opts ...grpc.CallOption) (*Ranking_RankList, error)
	QueryUsers(ctx context.Context, in *Ranking_Users, opts ...grpc.CallOption) (*Ranking_UserList, error)
}

type rankingServiceClient struct {
	cc *grpc.ClientConn
}

func NewRankingServiceClient(cc *grpc.ClientConn) RankingServiceClient {
	return &rankingServiceClient{cc}
}

func (c *rankingServiceClient) RankChange(ctx context.Context, in *Ranking_Change, opts ...grpc.CallOption) (*Ranking_Nil, error) {
	out := new(Ranking_Nil)
	err := grpc.Invoke(ctx, "/proto.RankingService/RankChange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankingServiceClient) DeleteSet(ctx context.Context, in *Ranking_SetId, opts ...grpc.CallOption) (*Ranking_Nil, error) {
	out := new(Ranking_Nil)
	err := grpc.Invoke(ctx, "/proto.RankingService/DeleteSet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankingServiceClient) DeleteUser(ctx context.Context, in *Ranking_DeleteUserRequest, opts ...grpc.CallOption) (*Ranking_Nil, error) {
	out := new(Ranking_Nil)
	err := grpc.Invoke(ctx, "/proto.RankingService/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankingServiceClient) QueryRankRange(ctx context.Context, in *Ranking_Range, opts ...grpc.CallOption) (*Ranking_RankList, error) {
	out := new(Ranking_RankList)
	err := grpc.Invoke(ctx, "/proto.RankingService/QueryRankRange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankingServiceClient) QueryUsers(ctx context.Context, in *Ranking_Users, opts ...grpc.CallOption) (*Ranking_UserList, error) {
	out := new(Ranking_UserList)
	err := grpc.Invoke(ctx, "/proto.RankingService/QueryUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RankingService service

type RankingServiceServer interface {
	RankChange(context.Context, *Ranking_Change) (*Ranking_Nil, error)
	DeleteSet(context.Context, *Ranking_SetId) (*Ranking_Nil, error)
	DeleteUser(context.Context, *Ranking_DeleteUserRequest) (*Ranking_Nil, error)
	QueryRankRange(context.Context, *Ranking_Range) (*Ranking_RankList, error)
	QueryUsers(context.Context, *Ranking_Users) (*Ranking_UserList, error)
}

func RegisterRankingServiceServer(s *grpc.Server, srv RankingServiceServer) {
	s.RegisterService(&_RankingService_serviceDesc, srv)
}

func _RankingService_RankChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ranking_Change)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankingServiceServer).RankChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RankingService/RankChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankingServiceServer).RankChange(ctx, req.(*Ranking_Change))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankingService_DeleteSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ranking_SetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankingServiceServer).DeleteSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RankingService/DeleteSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankingServiceServer).DeleteSet(ctx, req.(*Ranking_SetId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankingService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ranking_DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankingServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RankingService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankingServiceServer).DeleteUser(ctx, req.(*Ranking_DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankingService_QueryRankRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ranking_Range)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankingServiceServer).QueryRankRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RankingService/QueryRankRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankingServiceServer).QueryRankRange(ctx, req.(*Ranking_Range))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankingService_QueryUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ranking_Users)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankingServiceServer).QueryUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RankingService/QueryUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankingServiceServer).QueryUsers(ctx, req.(*Ranking_Users))
	}
	return interceptor(ctx, in, info, handler)
}

var _RankingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RankingService",
	HandlerType: (*RankingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RankChange",
			Handler:    _RankingService_RankChange_Handler,
		},
		{
			MethodName: "DeleteSet",
			Handler:    _RankingService_DeleteSet_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _RankingService_DeleteUser_Handler,
		},
		{
			MethodName: "QueryRankRange",
			Handler:    _RankingService_QueryRankRange_Handler,
		},
		{
			MethodName: "QueryUsers",
			Handler:    _RankingService_QueryUsers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x51, 0xdd, 0x4a, 0x33, 0x31,
	0x14, 0x64, 0xb7, 0x5f, 0xb6, 0xed, 0xc0, 0x57, 0xda, 0xf8, 0x4b, 0xae, 0x8a, 0x57, 0x05, 0xa5,
	0x60, 0x8b, 0x7a, 0xe1, 0x85, 0x58, 0xbd, 0x11, 0x44, 0xb0, 0xc5, 0x07, 0xa8, 0xf5, 0x50, 0x97,
	0x96, 0xae, 0x66, 0xb7, 0x82, 0xaf, 0xe7, 0x4b, 0xf8, 0x3a, 0x26, 0x27, 0xd1, 0x4a, 0x76, 0xc1,
	0xab, 0x4d, 0x26, 0x33, 0x67, 0xce, 0xce, 0xa0, 0xad, 0xa7, 0xab, 0x45, 0x4e, 0xfa, 0x8d, 0x74,
	0xff, 0x45, 0x67, 0x45, 0x26, 0x05, 0x7f, 0x0e, 0x3e, 0x63, 0xd4, 0xc7, 0xe6, 0x2d, 0x5d, 0xcd,
	0x95, 0x40, 0xed, 0x2e, 0x5d, 0xaa, 0x5d, 0x88, 0x09, 0x15, 0x37, 0x4f, 0xf2, 0xbf, 0x3f, 0xec,
	0x47, 0xdd, 0xa8, 0xf7, 0x4f, 0x0d, 0xd0, 0xb9, 0xa6, 0x25, 0x15, 0xf4, 0x60, 0x06, 0x8d, 0xe9,
	0x75, 0x4d, 0x79, 0x11, 0x70, 0x64, 0x0b, 0x89, 0x7d, 0x35, 0xf7, 0xd8, 0xdc, 0x85, 0x3a, 0x45,
	0x72, 0xf5, 0x3c, 0x5d, 0xcd, 0xe9, 0xd7, 0x8b, 0x65, 0x0a, 0x16, 0xce, 0x32, 0x4d, 0x8e, 0xb8,
	0x99, 0x53, 0x63, 0xaf, 0x23, 0x88, 0x31, 0xcb, 0x9a, 0x88, 0x2e, 0xbd, 0xc2, 0x1c, 0x47, 0xd5,
	0xec, 0x21, 0x1a, 0xf6, 0x1f, 0x6e, 0x53, 0xb3, 0xd0, 0x16, 0xea, 0xce, 0x27, 0x37, 0xb2, 0x5a,
	0x4f, 0x8c, 0xe2, 0x76, 0x24, 0x25, 0x12, 0x36, 0xcb, 0x8d, 0xde, 0x63, 0xea, 0x10, 0xc2, 0x12,
	0xf3, 0x6a, 0xc5, 0x8f, 0x43, 0xcc, 0x0e, 0xc7, 0x68, 0x58, 0x0e, 0x3b, 0x74, 0x78, 0xb7, 0xc5,
	0x1f, 0xf3, 0x07, 0x1f, 0x31, 0x5a, 0x3e, 0xd9, 0x89, 0x09, 0x3e, 0x9d, 0x91, 0x3c, 0x03, 0x2c,
	0xe2, 0x13, 0xd9, 0x71, 0x4d, 0xf4, 0x3d, 0xa9, 0xef, 0x60, 0x25, 0x03, 0xd8, 0x54, 0x22, 0x4f,
	0xd0, 0x74, 0xd1, 0x9b, 0x9d, 0xe4, 0x76, 0x40, 0xe0, 0x3d, 0x2b, 0x65, 0x23, 0x60, 0xd3, 0x98,
	0xec, 0x06, 0x8c, 0x52, 0x99, 0x95, 0x33, 0x2e, 0xd0, 0xba, 0x5f, 0x93, 0x7e, 0xb7, 0x98, 0xab,
	0x24, 0xf4, 0x67, 0x54, 0xed, 0x95, 0x51, 0x57, 0xc8, 0x39, 0xc0, 0x03, 0x5c, 0xd8, 0xa1, 0x98,
	0xd1, 0x92, 0xf8, 0x3b, 0xeb, 0xc7, 0x84, 0xf1, 0xe1, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x20, 0x0f, 0x43, 0xc0, 0x02, 0x00, 0x00,
}
