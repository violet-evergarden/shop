// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

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

type TestReq struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestReq) Reset()         { *m = TestReq{} }
func (m *TestReq) String() string { return proto.CompactTextString(m) }
func (*TestReq) ProtoMessage()    {}
func (*TestReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *TestReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestReq.Unmarshal(m, b)
}
func (m *TestReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestReq.Marshal(b, m, deterministic)
}
func (m *TestReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestReq.Merge(m, src)
}
func (m *TestReq) XXX_Size() int {
	return xxx_messageInfo_TestReq.Size(m)
}
func (m *TestReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TestReq.DiscardUnknown(m)
}

var xxx_messageInfo_TestReq proto.InternalMessageInfo

func (m *TestReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TestRsp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRsp) Reset()         { *m = TestRsp{} }
func (m *TestRsp) String() string { return proto.CompactTextString(m) }
func (*TestRsp) ProtoMessage()    {}
func (*TestRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *TestRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRsp.Unmarshal(m, b)
}
func (m *TestRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRsp.Marshal(b, m, deterministic)
}
func (m *TestRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRsp.Merge(m, src)
}
func (m *TestRsp) XXX_Size() int {
	return xxx_messageInfo_TestRsp.Size(m)
}
func (m *TestRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRsp.DiscardUnknown(m)
}

var xxx_messageInfo_TestRsp proto.InternalMessageInfo

func (m *TestRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*TestReq)(nil), "proto.TestReq")
	proto.RegisterType((*TestRsp)(nil), "proto.TestRsp")
}

func init() {
	proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e)
}

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xca, 0x5c, 0xec, 0x21, 0xa9,
	0xc5, 0x25, 0x41, 0xa9, 0x85, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0x2e, 0x5c, 0x51, 0x71, 0x01, 0x6e, 0x45, 0x46,
	0xfa, 0x5c, 0x2c, 0x20, 0x45, 0x42, 0xea, 0x5c, 0xcc, 0x8e, 0xc5, 0xd9, 0x42, 0x7c, 0x10, 0x7b,
	0xf4, 0xa0, 0xa6, 0x4b, 0xa1, 0xf0, 0x8b, 0x0b, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x02, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0x76, 0x30, 0x76, 0x96, 0x00, 0x00, 0x00,
}
