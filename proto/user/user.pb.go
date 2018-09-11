// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package customer is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserRequest
	UserResponse
	UserFilter
*/
package customer

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserRequest struct {
	Id   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserResponse struct {
	Id      int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type UserFilter struct {
	Keyword string `protobuf:"bytes,1,opt,name=keyword" json:"keyword,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *UserFilter) Reset()                    { *m = UserFilter{} }
func (m *UserFilter) String() string            { return proto.CompactTextString(m) }
func (*UserFilter) ProtoMessage()               {}
func (*UserFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *UserFilter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "customer.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "customer.UserResponse")
	proto.RegisterType((*UserFilter)(nil), "customer.UserFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	// Get all Users with filter - A server-to-client streaming RPC.
	GetUsers(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (User_GetUsersClient, error)
	// Create a new User - A simple RPC
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUsers(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (User_GetUsersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_User_serviceDesc.Streams[0], c.cc, "/customer.User/GetUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type User_GetUsersClient interface {
	Recv() (*UserRequest, error)
	grpc.ClientStream
}

type userGetUsersClient struct {
	grpc.ClientStream
}

func (x *userGetUsersClient) Recv() (*UserRequest, error) {
	m := new(UserRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/customer.User/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	// Get all Users with filter - A server-to-client streaming RPC.
	GetUsers(*UserFilter, User_GetUsersServer) error
	// Create a new User - A simple RPC
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServer).GetUsers(m, &userGetUsersServer{stream})
}

type User_GetUsersServer interface {
	Send(*UserRequest) error
	grpc.ServerStream
}

type userGetUsersServer struct {
	grpc.ServerStream
}

func (x *userGetUsersServer) Send(m *UserRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsers",
			Handler:       _User_GetUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4a, 0xc6, 0x30,
	0x14, 0xc5, 0x9b, 0x50, 0xb5, 0xbd, 0x8a, 0x43, 0x50, 0x29, 0x9d, 0x4a, 0xa6, 0x4e, 0xc5, 0x3f,
	0x20, 0x82, 0x38, 0x09, 0xba, 0x07, 0x7c, 0x80, 0xda, 0xde, 0x21, 0x68, 0x9b, 0x9a, 0x9b, 0x20,
	0xae, 0x3e, 0xb9, 0x24, 0x69, 0x11, 0xfd, 0xbe, 0xed, 0x9e, 0xc3, 0xf9, 0x9d, 0x70, 0x02, 0xe0,
	0x09, 0x6d, 0xb7, 0x58, 0xe3, 0x8c, 0x28, 0x06, 0x4f, 0xce, 0x4c, 0x68, 0xe5, 0x15, 0x1c, 0xbf,
	0x10, 0x5a, 0x85, 0x1f, 0x1e, 0xc9, 0x89, 0x53, 0xe0, 0x7a, 0xac, 0x58, 0xc3, 0xda, 0x03, 0xc5,
	0xf5, 0x28, 0x04, 0xe4, 0x73, 0x3f, 0x61, 0xc5, 0x1b, 0xd6, 0x96, 0x2a, 0xde, 0xf2, 0x0e, 0x4e,
	0x12, 0x42, 0x8b, 0x99, 0x09, 0x77, 0x98, 0x0a, 0x8e, 0xc8, 0x0f, 0x03, 0x12, 0x45, 0xac, 0x50,
	0x9b, 0x94, 0xb7, 0x00, 0x81, 0x7c, 0xd2, 0xef, 0x0e, 0x6d, 0xc8, 0xbd, 0xe1, 0xd7, 0xa7, 0xb1,
	0x09, 0x2e, 0xd5, 0x26, 0xd7, 0xc6, 0xf4, 0x26, 0xd7, 0xe3, 0xf5, 0x37, 0x83, 0x3c, 0x80, 0xe2,
	0x1e, 0x8a, 0x67, 0x74, 0xe1, 0x24, 0x71, 0xd6, 0x6d, 0x23, 0xba, 0xdf, 0xd2, 0xfa, 0xfc, 0xaf,
	0xbb, 0xee, 0x92, 0xd9, 0x25, 0x13, 0x0f, 0x00, 0x8f, 0x16, 0x7b, 0x87, 0xb1, 0x6a, 0x7f, 0xb0,
	0xbe, 0xf8, 0x6f, 0xa7, 0x91, 0x32, 0x7b, 0x3d, 0x8c, 0x5f, 0x77, 0xf3, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x3b, 0x8d, 0x9f, 0x5d, 0x48, 0x01, 0x00, 0x00,
}