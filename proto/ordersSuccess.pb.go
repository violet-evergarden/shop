// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ordersSuccess.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UsersOrdersSuccessReq struct {
	PlatformOrderNo      uint64   `protobuf:"varint,1,opt,name=platformOrderNo,proto3" json:"platformOrderNo,omitempty"`
	Status               uint64   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersOrdersSuccessReq) Reset()         { *m = UsersOrdersSuccessReq{} }
func (m *UsersOrdersSuccessReq) String() string { return proto.CompactTextString(m) }
func (*UsersOrdersSuccessReq) ProtoMessage()    {}
func (*UsersOrdersSuccessReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8245efb4bf621cf8, []int{0}
}

func (m *UsersOrdersSuccessReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersOrdersSuccessReq.Unmarshal(m, b)
}
func (m *UsersOrdersSuccessReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersOrdersSuccessReq.Marshal(b, m, deterministic)
}
func (m *UsersOrdersSuccessReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersOrdersSuccessReq.Merge(m, src)
}
func (m *UsersOrdersSuccessReq) XXX_Size() int {
	return xxx_messageInfo_UsersOrdersSuccessReq.Size(m)
}
func (m *UsersOrdersSuccessReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersOrdersSuccessReq.DiscardUnknown(m)
}

var xxx_messageInfo_UsersOrdersSuccessReq proto.InternalMessageInfo

func (m *UsersOrdersSuccessReq) GetPlatformOrderNo() uint64 {
	if m != nil {
		return m.PlatformOrderNo
	}
	return 0
}

func (m *UsersOrdersSuccessReq) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type UsersOrdersSuccessRsp struct {
	Code                 uint64   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersOrdersSuccessRsp) Reset()         { *m = UsersOrdersSuccessRsp{} }
func (m *UsersOrdersSuccessRsp) String() string { return proto.CompactTextString(m) }
func (*UsersOrdersSuccessRsp) ProtoMessage()    {}
func (*UsersOrdersSuccessRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8245efb4bf621cf8, []int{1}
}

func (m *UsersOrdersSuccessRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersOrdersSuccessRsp.Unmarshal(m, b)
}
func (m *UsersOrdersSuccessRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersOrdersSuccessRsp.Marshal(b, m, deterministic)
}
func (m *UsersOrdersSuccessRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersOrdersSuccessRsp.Merge(m, src)
}
func (m *UsersOrdersSuccessRsp) XXX_Size() int {
	return xxx_messageInfo_UsersOrdersSuccessRsp.Size(m)
}
func (m *UsersOrdersSuccessRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersOrdersSuccessRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UsersOrdersSuccessRsp proto.InternalMessageInfo

func (m *UsersOrdersSuccessRsp) GetCode() uint64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UsersOrdersSuccessRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UsersOrdersSuccessReq)(nil), "proto.UsersOrdersSuccessReq")
	proto.RegisterType((*UsersOrdersSuccessRsp)(nil), "proto.UsersOrdersSuccessRsp")
}

func init() {
	proto.RegisterFile("ordersSuccess.proto", fileDescriptor_8245efb4bf621cf8)
}

var fileDescriptor_8245efb4bf621cf8 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x2f, 0x4a, 0x49,
	0x2d, 0x2a, 0x0e, 0x2e, 0x4d, 0x4e, 0x4e, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x53, 0x4a, 0x91, 0x5c, 0xa2, 0xa1, 0xc5, 0xa9, 0x45, 0xc5, 0xfe, 0xc8, 0x4a, 0x82,
	0x52, 0x0b, 0x85, 0x34, 0xb8, 0xf8, 0x0b, 0x72, 0x12, 0x4b, 0xd2, 0xf2, 0x8b, 0x72, 0xc1, 0x72,
	0x7e, 0xf9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0xe8, 0xc2, 0x42, 0x62, 0x5c, 0x6c, 0xc5,
	0x25, 0x89, 0x25, 0xa5, 0xc5, 0x12, 0x4c, 0x60, 0x05, 0x50, 0x9e, 0x92, 0x2b, 0x56, 0xa3, 0x8b,
	0x0b, 0x84, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53, 0x52, 0xa1, 0xe6, 0x81, 0xd9, 0x42, 0x12, 0x5c,
	0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x60, 0x53, 0x38, 0x83, 0x60, 0x5c, 0xa3, 0x68,
	0x2e, 0x21, 0x4c, 0x63, 0x84, 0x5c, 0xb9, 0x58, 0xc1, 0x02, 0x42, 0x32, 0x10, 0xff, 0xe8, 0x61,
	0xf5, 0x85, 0x14, 0x1e, 0xd9, 0xe2, 0x02, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xb4, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xd1, 0x66, 0x1d, 0x5b, 0x23, 0x01, 0x00, 0x00,
}