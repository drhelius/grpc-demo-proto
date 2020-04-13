// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateUserReq struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserReq) Reset()         { *m = CreateUserReq{} }
func (m *CreateUserReq) String() string { return proto.CompactTextString(m) }
func (*CreateUserReq) ProtoMessage()    {}
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{1}
}

func (m *CreateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReq.Unmarshal(m, b)
}
func (m *CreateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReq.Marshal(b, m, deterministic)
}
func (m *CreateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReq.Merge(m, src)
}
func (m *CreateUserReq) XXX_Size() int {
	return xxx_messageInfo_CreateUserReq.Size(m)
}
func (m *CreateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReq proto.InternalMessageInfo

func (m *CreateUserReq) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResp struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResp) Reset()         { *m = CreateUserResp{} }
func (m *CreateUserResp) String() string { return proto.CompactTextString(m) }
func (*CreateUserResp) ProtoMessage()    {}
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{2}
}

func (m *CreateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResp.Unmarshal(m, b)
}
func (m *CreateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResp.Marshal(b, m, deterministic)
}
func (m *CreateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResp.Merge(m, src)
}
func (m *CreateUserResp) XXX_Size() int {
	return xxx_messageInfo_CreateUserResp.Size(m)
}
func (m *CreateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResp proto.InternalMessageInfo

func (m *CreateUserResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadUserReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadUserReq) Reset()         { *m = ReadUserReq{} }
func (m *ReadUserReq) String() string { return proto.CompactTextString(m) }
func (*ReadUserReq) ProtoMessage()    {}
func (*ReadUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{3}
}

func (m *ReadUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadUserReq.Unmarshal(m, b)
}
func (m *ReadUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadUserReq.Marshal(b, m, deterministic)
}
func (m *ReadUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadUserReq.Merge(m, src)
}
func (m *ReadUserReq) XXX_Size() int {
	return xxx_messageInfo_ReadUserReq.Size(m)
}
func (m *ReadUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadUserReq proto.InternalMessageInfo

func (m *ReadUserReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadUserResp struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadUserResp) Reset()         { *m = ReadUserResp{} }
func (m *ReadUserResp) String() string { return proto.CompactTextString(m) }
func (*ReadUserResp) ProtoMessage()    {}
func (*ReadUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{4}
}

func (m *ReadUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadUserResp.Unmarshal(m, b)
}
func (m *ReadUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadUserResp.Marshal(b, m, deterministic)
}
func (m *ReadUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadUserResp.Merge(m, src)
}
func (m *ReadUserResp) XXX_Size() int {
	return xxx_messageInfo_ReadUserResp.Size(m)
}
func (m *ReadUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_ReadUserResp proto.InternalMessageInfo

func (m *ReadUserResp) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*CreateUserReq)(nil), "user.CreateUserReq")
	proto.RegisterType((*CreateUserResp)(nil), "user.CreateUserResp")
	proto.RegisterType((*ReadUserReq)(nil), "user.ReadUserReq")
	proto.RegisterType((*ReadUserResp)(nil), "user.ReadUserResp")
}

func init() {
	proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74)
}

var fileDescriptor_ed89022014131a74 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2d, 0x4e, 0x2d,
	0xd2, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x94, 0x4c, 0x7a,
	0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62,
	0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x8d, 0x92, 0x03, 0x17, 0x4b, 0x68, 0x71, 0x6a, 0x91, 0x10,
	0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x53, 0x66, 0x8a, 0x90,
	0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58, 0x04, 0xcc, 0x16, 0x12, 0xe1, 0x62,
	0x4d, 0xcd, 0x4d, 0xcc, 0xcc, 0x91, 0x60, 0x06, 0x0b, 0x42, 0x38, 0x4a, 0xfa, 0x5c, 0xbc, 0xce,
	0x45, 0xa9, 0x89, 0x25, 0xa9, 0x20, 0x73, 0x82, 0x52, 0x0b, 0x85, 0xe4, 0xb8, 0xc0, 0x16, 0x83,
	0x0d, 0xe3, 0x36, 0xe2, 0xd2, 0x03, 0xbb, 0x08, 0x2c, 0x09, 0x16, 0x57, 0x52, 0xe0, 0xe2, 0x43,
	0xd6, 0x50, 0x5c, 0x80, 0x6e, 0xb9, 0x92, 0x2c, 0x17, 0x77, 0x50, 0x6a, 0x62, 0x0a, 0xcc, 0x40,
	0x74, 0x69, 0x3d, 0x2e, 0x1e, 0x84, 0x74, 0x71, 0x01, 0x21, 0x0b, 0x8d, 0x16, 0x30, 0x72, 0x71,
	0x83, 0xb8, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xde, 0x5c, 0x6c, 0x10, 0x07, 0x08,
	0x09, 0x43, 0xd4, 0xa2, 0xb8, 0x5f, 0x4a, 0x04, 0x53, 0xb0, 0xb8, 0x40, 0x49, 0xac, 0xe9, 0xf2,
	0x93, 0xc9, 0x4c, 0x02, 0x4a, 0x1c, 0xfa, 0x65, 0x86, 0xe0, 0x40, 0xb6, 0x02, 0x1b, 0x2e, 0xe4,
	0xc2, 0xc5, 0x02, 0x72, 0x8c, 0x90, 0x20, 0x44, 0x17, 0x92, 0xbb, 0xa5, 0x84, 0xd0, 0x85, 0x8a,
	0x0b, 0x94, 0x44, 0xc1, 0xc6, 0xf0, 0x0b, 0xf1, 0xc2, 0x8c, 0xd1, 0xaf, 0xce, 0x4c, 0xa9, 0x4d,
	0x62, 0x03, 0xc7, 0x86, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x06, 0xf4, 0x84, 0xbf, 0xc4, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error)
	Read(ctx context.Context, in *ReadUserReq, opts ...grpc.CallOption) (*ReadUserResp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error) {
	out := new(CreateUserResp)
	err := c.cc.Invoke(ctx, "/user.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Read(ctx context.Context, in *ReadUserReq, opts ...grpc.CallOption) (*ReadUserResp, error) {
	out := new(ReadUserResp)
	err := c.cc.Invoke(ctx, "/user.UserService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *CreateUserReq) (*CreateUserResp, error)
	Read(context.Context, *ReadUserReq) (*ReadUserResp, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Create(ctx context.Context, req *CreateUserReq) (*CreateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) Read(ctx context.Context, req *ReadUserReq) (*ReadUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Read(ctx, req.(*ReadUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _UserService_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
