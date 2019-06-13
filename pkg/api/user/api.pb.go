// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/alanchchen/go-project-skeleton/pkg/api/user/api.proto

package user // import "github.com/alanchchen/go-project-skeleton/pkg/api/user"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_01807ff94bab253b, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserRequest) Reset()         { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()    {}
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_01807ff94bab253b, []int{1}
}
func (m *AddUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserRequest.Unmarshal(m, b)
}
func (m *AddUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserRequest.Marshal(b, m, deterministic)
}
func (dst *AddUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserRequest.Merge(dst, src)
}
func (m *AddUserRequest) XXX_Size() int {
	return xxx_messageInfo_AddUserRequest.Size(m)
}
func (m *AddUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserRequest proto.InternalMessageInfo

func (m *AddUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FindUserByIdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUserByIdRequest) Reset()         { *m = FindUserByIdRequest{} }
func (m *FindUserByIdRequest) String() string { return proto.CompactTextString(m) }
func (*FindUserByIdRequest) ProtoMessage()    {}
func (*FindUserByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_01807ff94bab253b, []int{2}
}
func (m *FindUserByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUserByIdRequest.Unmarshal(m, b)
}
func (m *FindUserByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUserByIdRequest.Marshal(b, m, deterministic)
}
func (dst *FindUserByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUserByIdRequest.Merge(dst, src)
}
func (m *FindUserByIdRequest) XXX_Size() int {
	return xxx_messageInfo_FindUserByIdRequest.Size(m)
}
func (m *FindUserByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUserByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindUserByIdRequest proto.InternalMessageInfo

func (m *FindUserByIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FindUserByNameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUserByNameRequest) Reset()         { *m = FindUserByNameRequest{} }
func (m *FindUserByNameRequest) String() string { return proto.CompactTextString(m) }
func (*FindUserByNameRequest) ProtoMessage()    {}
func (*FindUserByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_01807ff94bab253b, []int{3}
}
func (m *FindUserByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUserByNameRequest.Unmarshal(m, b)
}
func (m *FindUserByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUserByNameRequest.Marshal(b, m, deterministic)
}
func (dst *FindUserByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUserByNameRequest.Merge(dst, src)
}
func (m *FindUserByNameRequest) XXX_Size() int {
	return xxx_messageInfo_FindUserByNameRequest.Size(m)
}
func (m *FindUserByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUserByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindUserByNameRequest proto.InternalMessageInfo

func (m *FindUserByNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_01807ff94bab253b, []int{4}
}
func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (dst *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(dst, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*AddUserRequest)(nil), "user.AddUserRequest")
	proto.RegisterType((*FindUserByIdRequest)(nil), "user.FindUserByIdRequest")
	proto.RegisterType((*FindUserByNameRequest)(nil), "user.FindUserByNameRequest")
	proto.RegisterType((*Users)(nil), "user.Users")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*Users, error)
	FindUserById(ctx context.Context, in *FindUserByIdRequest, opts ...grpc.CallOption) (*Users, error)
	FindUserByName(ctx context.Context, in *FindUserByNameRequest, opts ...grpc.CallOption) (*Users, error)
	ListUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Users, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/user.Service/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FindUserById(ctx context.Context, in *FindUserByIdRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/user.Service/FindUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FindUserByName(ctx context.Context, in *FindUserByNameRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/user.Service/FindUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/user.Service/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	AddUser(context.Context, *AddUserRequest) (*Users, error)
	FindUserById(context.Context, *FindUserByIdRequest) (*Users, error)
	FindUserByName(context.Context, *FindUserByNameRequest) (*Users, error)
	ListUsers(context.Context, *empty.Empty) (*Users, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Service/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FindUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FindUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Service/FindUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FindUserById(ctx, req.(*FindUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FindUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FindUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Service/FindUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FindUserByName(ctx, req.(*FindUserByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Service/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListUsers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Service_AddUser_Handler,
		},
		{
			MethodName: "FindUserById",
			Handler:    _Service_FindUserById_Handler,
		},
		{
			MethodName: "FindUserByName",
			Handler:    _Service_FindUserByName_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Service_ListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/alanchchen/go-project-skeleton/pkg/api/user/api.proto",
}

func init() {
	proto.RegisterFile("github.com/alanchchen/go-project-skeleton/pkg/api/user/api.proto", fileDescriptor_api_01807ff94bab253b)
}

var fileDescriptor_api_01807ff94bab253b = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xdd, 0xfe, 0xb1, 0x74, 0x2a, 0x3d, 0xac, 0x7f, 0xa8, 0xed, 0x25, 0x04, 0x85, 0xaa,
	0x74, 0x17, 0xea, 0x45, 0xa8, 0x82, 0x16, 0x14, 0x04, 0xf1, 0x50, 0xf1, 0xe2, 0x2d, 0x4d, 0xc6,
	0x74, 0x6d, 0x93, 0x8d, 0xd9, 0x8d, 0xd0, 0x0f, 0xec, 0xf7, 0x90, 0xcd, 0x5a, 0x9b, 0x06, 0xf1,
	0xe0, 0x25, 0xcc, 0x64, 0x7e, 0x6f, 0xf2, 0x1e, 0x13, 0xb8, 0x0e, 0x85, 0x9e, 0x65, 0x53, 0xe6,
	0xcb, 0x88, 0x7b, 0x0b, 0x2f, 0xf6, 0x67, 0xfe, 0x0c, 0x63, 0x1e, 0xca, 0x41, 0x92, 0xca, 0x37,
	0xf4, 0xf5, 0x40, 0xcd, 0x71, 0x81, 0x5a, 0xc6, 0x3c, 0x99, 0x87, 0xdc, 0x4b, 0x04, 0xcf, 0x14,
	0xa6, 0xa6, 0x60, 0x49, 0x2a, 0xb5, 0xa4, 0x35, 0xd3, 0x77, 0x7b, 0xa1, 0x94, 0xe1, 0x02, 0x79,
	0xfe, 0x6e, 0x9a, 0xbd, 0x72, 0x8c, 0x12, 0xbd, 0xb4, 0x88, 0x7b, 0x0a, 0xb5, 0x67, 0x85, 0x29,
	0x6d, 0x43, 0x45, 0x04, 0x1d, 0xe2, 0x90, 0x7e, 0x75, 0x52, 0x11, 0x01, 0xa5, 0x50, 0x8b, 0xbd,
	0x08, 0x3b, 0x15, 0x87, 0xf4, 0x9b, 0x93, 0xbc, 0x76, 0x8f, 0xa0, 0x7d, 0x13, 0x04, 0x06, 0x9f,
	0xe0, 0x7b, 0x86, 0x4a, 0xff, 0x50, 0xa4, 0x40, 0x1d, 0xc3, 0xee, 0x9d, 0x88, 0x73, 0x6c, 0xbc,
	0xbc, 0x0f, 0x56, 0x68, 0xe9, 0x03, 0xee, 0x19, 0xec, 0xaf, 0xb1, 0x47, 0x2f, 0xc2, 0xbf, 0x76,
	0x9e, 0x40, 0xdd, 0x80, 0x8a, 0x3a, 0x50, 0x37, 0x99, 0x54, 0x87, 0x38, 0xd5, 0x7e, 0x6b, 0x08,
	0xcc, 0x74, 0x2c, 0xb7, 0x64, 0x07, 0xc3, 0x4f, 0x02, 0x8d, 0x27, 0x4c, 0x3f, 0x84, 0x8f, 0x94,
	0x41, 0xe3, 0xdb, 0x30, 0xdd, 0xb3, 0xe4, 0xa6, 0xff, 0x6e, 0x6b, 0xad, 0x57, 0xee, 0x16, 0xbd,
	0x80, 0x9d, 0xa2, 0x75, 0x7a, 0x68, 0xc7, 0xbf, 0xc4, 0x29, 0x2b, 0x2f, 0xa1, 0xbd, 0x99, 0x86,
	0xf6, 0xca, 0xda, 0x42, 0xc6, 0xb2, 0x7a, 0x08, 0xcd, 0x07, 0xa1, 0xb4, 0x8d, 0x78, 0xc0, 0xec,
	0xbd, 0xd8, 0xea, 0x5e, 0xec, 0xd6, 0xdc, 0xab, 0xa4, 0x19, 0x5f, 0xbd, 0x8c, 0xfe, 0xf7, 0x7f,
	0x8c, 0xcc, 0x63, 0xba, 0x9d, 0x6f, 0x3f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x55, 0xa9, 0x39,
	0x3e, 0x65, 0x02, 0x00, 0x00,
}
