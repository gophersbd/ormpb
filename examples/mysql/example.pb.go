// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ormpb/examples/mysql/example.proto

package importpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/gophersbd/ormpb/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Example struct {
	UserId               int32                `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Point                float64              `protobuf:"fixed64,4,opt,name=point" json:"point,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Example) Reset()         { *m = Example{} }
func (m *Example) String() string { return proto.CompactTextString(m) }
func (*Example) ProtoMessage()    {}
func (*Example) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_12ba9d5c4adf2079, []int{0}
}
func (m *Example) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Example.Unmarshal(m, b)
}
func (m *Example) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Example.Marshal(b, m, deterministic)
}
func (dst *Example) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Example.Merge(dst, src)
}
func (m *Example) XXX_Size() int {
	return xxx_messageInfo_Example.Size(m)
}
func (m *Example) XXX_DiscardUnknown() {
	xxx_messageInfo_Example.DiscardUnknown(m)
}

var xxx_messageInfo_Example proto.InternalMessageInfo

func (m *Example) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Example) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Example) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Example) GetPoint() float64 {
	if m != nil {
		return m.Point
	}
	return 0
}

func (m *Example) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type ExampleAutoIncrement struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExampleAutoIncrement) Reset()         { *m = ExampleAutoIncrement{} }
func (m *ExampleAutoIncrement) String() string { return proto.CompactTextString(m) }
func (*ExampleAutoIncrement) ProtoMessage()    {}
func (*ExampleAutoIncrement) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_12ba9d5c4adf2079, []int{1}
}
func (m *ExampleAutoIncrement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExampleAutoIncrement.Unmarshal(m, b)
}
func (m *ExampleAutoIncrement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExampleAutoIncrement.Marshal(b, m, deterministic)
}
func (dst *ExampleAutoIncrement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleAutoIncrement.Merge(dst, src)
}
func (m *ExampleAutoIncrement) XXX_Size() int {
	return xxx_messageInfo_ExampleAutoIncrement.Size(m)
}
func (m *ExampleAutoIncrement) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleAutoIncrement.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleAutoIncrement proto.InternalMessageInfo

func (m *ExampleAutoIncrement) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*Example)(nil), "ormpb.examples.mysql.Example")
	proto.RegisterType((*ExampleAutoIncrement)(nil), "ormpb.examples.mysql.ExampleAutoIncrement")
}

func init() {
	proto.RegisterFile("ormpb/examples/mysql/example.proto", fileDescriptor_example_12ba9d5c4adf2079)
}

var fileDescriptor_example_12ba9d5c4adf2079 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe5, 0xb6, 0x69, 0x5a, 0xff, 0xbf, 0x18, 0xac, 0x0a, 0x59, 0xe9, 0x40, 0x14, 0xa1,
	0x12, 0x16, 0x47, 0xb4, 0x03, 0xd0, 0x89, 0x46, 0x62, 0xe8, 0x1a, 0x31, 0xb1, 0x44, 0x4e, 0x6a,
	0x2a, 0x4b, 0x71, 0x6c, 0x12, 0x47, 0x82, 0x8d, 0x99, 0xb1, 0x23, 0x4f, 0xd8, 0x97, 0x40, 0x42,
	0xb1, 0x13, 0xc4, 0x80, 0xc4, 0x96, 0x7b, 0xbf, 0x73, 0x4f, 0x74, 0x8e, 0x61, 0x20, 0x2b, 0xa1,
	0xb2, 0x88, 0xbd, 0x50, 0xa1, 0x0a, 0x56, 0x47, 0xe2, 0xb5, 0x7e, 0x2e, 0xfa, 0x91, 0xa8, 0x4a,
	0x6a, 0x89, 0x66, 0x46, 0x43, 0x7a, 0x0d, 0x31, 0x1a, 0xef, 0x6c, 0x2f, 0xe5, 0xbe, 0x60, 0x91,
	0xd1, 0x64, 0xcd, 0x53, 0xa4, 0xb9, 0x60, 0xb5, 0xa6, 0x42, 0xd9, 0x33, 0x6f, 0x6e, 0xad, 0xbf,
	0xb9, 0x54, 0x9a, 0xcb, 0xd2, 0xc2, 0xe0, 0x13, 0x40, 0xf7, 0xde, 0x1a, 0x22, 0x1f, 0xba, 0x4d,
	0xcd, 0xaa, 0x94, 0xef, 0x30, 0xf0, 0x41, 0xe8, 0xc4, 0xee, 0xc7, 0xf1, 0xb0, 0x18, 0xdc, 0x81,
	0x64, 0xdc, 0xee, 0xb7, 0x3b, 0x44, 0xe0, 0xa8, 0xa4, 0x82, 0xe1, 0x81, 0x0f, 0xc2, 0x69, 0xec,
	0xb5, 0x78, 0x6c, 0x57, 0xed, 0xe7, 0x10, 0xbf, 0x01, 0x73, 0xe1, 0x83, 0xc4, 0x2c, 0xd1, 0x39,
	0x74, 0x98, 0xa0, 0xbc, 0xc0, 0x43, 0x73, 0x70, 0x62, 0x68, 0x68, 0x45, 0x37, 0x20, 0xb1, 0x10,
	0x05, 0xd0, 0x51, 0x92, 0x97, 0x1a, 0x8f, 0x7c, 0x10, 0x82, 0xf8, 0x7f, 0x8b, 0xdd, 0xa5, 0x73,
	0x75, 0x4d, 0x56, 0xab, 0xc4, 0x22, 0x74, 0x0b, 0x61, 0x5e, 0x31, 0xaa, 0xd9, 0x2e, 0xa5, 0x1a,
	0x3b, 0x3e, 0x08, 0xff, 0x2d, 0x3d, 0x62, 0xa3, 0x93, 0x3e, 0x1a, 0x79, 0xe8, 0xa3, 0x27, 0xd3,
	0x4e, 0xbd, 0xd1, 0xeb, 0xf9, 0xfb, 0xf1, 0xb0, 0x80, 0x70, 0xd2, 0xf7, 0xd6, 0x8e, 0x2e, 0x72,
	0x4c, 0x7b, 0x41, 0x0e, 0x67, 0x5d, 0xfc, 0x4d, 0xa3, 0xe5, 0xb6, 0xcc, 0x2b, 0x26, 0x58, 0xa9,
	0xff, 0xee, 0x62, 0x7d, 0xd9, 0xfa, 0x60, 0x78, 0xda, 0xd9, 0xa6, 0xb4, 0xd1, 0x32, 0xe5, 0xbd,
	0xc1, 0xcf, 0x9f, 0xc4, 0x17, 0x70, 0x9e, 0x4b, 0x41, 0x7e, 0x7b, 0x3e, 0xa2, 0xb2, 0xc7, 0x09,
	0x17, 0x4a, 0x56, 0x5a, 0x65, 0xd9, 0xd8, 0x24, 0x59, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x82,
	0x5e, 0x0a, 0x31, 0x0e, 0x02, 0x00, 0x00,
}
