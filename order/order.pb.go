// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order/order.proto

package order

import (
	fmt "fmt"
	product "github.com/drhelius/grpc-demo-proto/product"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Order struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Date                 int64              `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
	Products             []*product.Product `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Order) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Order) GetProducts() []*product.Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type CreateOrderReq struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderReq) Reset()         { *m = CreateOrderReq{} }
func (m *CreateOrderReq) String() string { return proto.CompactTextString(m) }
func (*CreateOrderReq) ProtoMessage()    {}
func (*CreateOrderReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{1}
}

func (m *CreateOrderReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderReq.Unmarshal(m, b)
}
func (m *CreateOrderReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderReq.Marshal(b, m, deterministic)
}
func (m *CreateOrderReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderReq.Merge(m, src)
}
func (m *CreateOrderReq) XXX_Size() int {
	return xxx_messageInfo_CreateOrderReq.Size(m)
}
func (m *CreateOrderReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderReq proto.InternalMessageInfo

func (m *CreateOrderReq) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type CreateOrderResp struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderResp) Reset()         { *m = CreateOrderResp{} }
func (m *CreateOrderResp) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResp) ProtoMessage()    {}
func (*CreateOrderResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{2}
}

func (m *CreateOrderResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResp.Unmarshal(m, b)
}
func (m *CreateOrderResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResp.Marshal(b, m, deterministic)
}
func (m *CreateOrderResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResp.Merge(m, src)
}
func (m *CreateOrderResp) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResp.Size(m)
}
func (m *CreateOrderResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResp proto.InternalMessageInfo

func (m *CreateOrderResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadOrderReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadOrderReq) Reset()         { *m = ReadOrderReq{} }
func (m *ReadOrderReq) String() string { return proto.CompactTextString(m) }
func (*ReadOrderReq) ProtoMessage()    {}
func (*ReadOrderReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{3}
}

func (m *ReadOrderReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadOrderReq.Unmarshal(m, b)
}
func (m *ReadOrderReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadOrderReq.Marshal(b, m, deterministic)
}
func (m *ReadOrderReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadOrderReq.Merge(m, src)
}
func (m *ReadOrderReq) XXX_Size() int {
	return xxx_messageInfo_ReadOrderReq.Size(m)
}
func (m *ReadOrderReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadOrderReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadOrderReq proto.InternalMessageInfo

func (m *ReadOrderReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadOrderResp struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadOrderResp) Reset()         { *m = ReadOrderResp{} }
func (m *ReadOrderResp) String() string { return proto.CompactTextString(m) }
func (*ReadOrderResp) ProtoMessage()    {}
func (*ReadOrderResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa47a2077d8980ed, []int{4}
}

func (m *ReadOrderResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadOrderResp.Unmarshal(m, b)
}
func (m *ReadOrderResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadOrderResp.Marshal(b, m, deterministic)
}
func (m *ReadOrderResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadOrderResp.Merge(m, src)
}
func (m *ReadOrderResp) XXX_Size() int {
	return xxx_messageInfo_ReadOrderResp.Size(m)
}
func (m *ReadOrderResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadOrderResp.DiscardUnknown(m)
}

var xxx_messageInfo_ReadOrderResp proto.InternalMessageInfo

func (m *ReadOrderResp) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "order.Order")
	proto.RegisterType((*CreateOrderReq)(nil), "order.CreateOrderReq")
	proto.RegisterType((*CreateOrderResp)(nil), "order.CreateOrderResp")
	proto.RegisterType((*ReadOrderReq)(nil), "order.ReadOrderReq")
	proto.RegisterType((*ReadOrderResp)(nil), "order.ReadOrderResp")
}

func init() {
	proto.RegisterFile("order/order.proto", fileDescriptor_fa47a2077d8980ed)
}

var fileDescriptor_fa47a2077d8980ed = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x4e, 0xf3, 0x30,
	0x10, 0x55, 0xd2, 0x1f, 0x7d, 0x9d, 0xf6, 0x2b, 0x65, 0xa0, 0x55, 0x14, 0x21, 0x54, 0xbc, 0xea,
	0x02, 0xc5, 0x22, 0x65, 0xc5, 0x96, 0x15, 0x2b, 0x20, 0x9c, 0xc0, 0xd4, 0x56, 0x65, 0x89, 0xc6,
	0xc6, 0x36, 0xdd, 0x20, 0x36, 0x5c, 0x81, 0x6b, 0x70, 0x1b, 0xae, 0xc0, 0x41, 0x50, 0xec, 0x10,
	0xb5, 0x85, 0x05, 0x9b, 0x78, 0xf2, 0xde, 0xcc, 0x7b, 0x6f, 0x34, 0xb0, 0xaf, 0x0c, 0x17, 0x86,
	0xfa, 0x6f, 0xa6, 0x8d, 0x72, 0x0a, 0x3b, 0xfe, 0x27, 0x3d, 0x5a, 0x2a, 0xb5, 0x7c, 0x10, 0x94,
	0x69, 0x49, 0x59, 0x59, 0x2a, 0xc7, 0x9c, 0x54, 0xa5, 0x0d, 0x4d, 0xe9, 0x58, 0x1b, 0xc5, 0x9f,
	0x16, 0x8e, 0xd6, 0x6f, 0x80, 0xc9, 0x0a, 0x3a, 0xd7, 0xd5, 0x34, 0x0e, 0x21, 0x96, 0x3c, 0x89,
	0xa6, 0xd1, 0xac, 0x57, 0xc4, 0x92, 0x23, 0x42, 0xbb, 0x64, 0x2b, 0x91, 0xc4, 0x1e, 0xf1, 0x75,
	0x85, 0x71, 0xe6, 0x44, 0xd2, 0x9a, 0x46, 0xb3, 0x56, 0xe1, 0x6b, 0x3c, 0x85, 0x7f, 0xb5, 0xa2,
	0x4d, 0xda, 0xd3, 0xd6, 0xac, 0x9f, 0x8f, 0xb2, 0x6f, 0x8b, 0x9b, 0xf0, 0x16, 0x4d, 0x07, 0x39,
	0x87, 0xe1, 0xa5, 0x11, 0xcc, 0x09, 0x6f, 0x5a, 0x88, 0x47, 0x24, 0x10, 0xe2, 0x7b, 0xeb, 0x7e,
	0x3e, 0xc8, 0xc2, 0x66, 0x81, 0x0f, 0x14, 0x39, 0x81, 0xbd, 0xad, 0x29, 0xab, 0x77, 0xe3, 0x92,
	0x63, 0x18, 0x14, 0x82, 0xf1, 0x46, 0x76, 0x97, 0x9f, 0xc3, 0xff, 0x0d, 0xde, 0xea, 0xbf, 0xf8,
	0xe6, 0xef, 0x11, 0x0c, 0x3c, 0x70, 0x27, 0xcc, 0x5a, 0x2e, 0x04, 0xde, 0x42, 0x37, 0x04, 0xc1,
	0x71, 0xdd, 0xbf, 0xbd, 0x4d, 0x3a, 0xf9, 0x0d, 0xb6, 0x9a, 0x24, 0xaf, 0x1f, 0x9f, 0x6f, 0x31,
	0x92, 0x1e, 0x5d, 0x9f, 0x85, 0xdb, 0x5d, 0x04, 0x0f, 0xbc, 0x82, 0x76, 0x15, 0x0c, 0x0f, 0xea,
	0xc9, 0xcd, 0x2d, 0xd2, 0xc3, 0x9f, 0xa0, 0xd5, 0x64, 0xe2, 0xc5, 0x46, 0x38, 0x6c, 0xc4, 0xe8,
	0xb3, 0xe4, 0x2f, 0xf7, 0x5d, 0x7f, 0xd2, 0xf9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x5d,
	0xf5, 0x18, 0x23, 0x02, 0x00, 0x00,
}