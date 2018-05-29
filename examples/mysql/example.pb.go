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
	return fileDescriptor_example_eebe85aef3454405, []int{0}
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
	return fileDescriptor_example_eebe85aef3454405, []int{1}
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
	proto.RegisterFile("ormpb/examples/mysql/example.proto", fileDescriptor_example_eebe85aef3454405)
}

var fileDescriptor_example_eebe85aef3454405 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4b, 0xfb, 0x30,
	0x1c, 0xc5, 0xc9, 0xb6, 0xae, 0x5b, 0x7e, 0x3f, 0x3c, 0xc4, 0x21, 0x65, 0x3b, 0x58, 0x8a, 0x6c,
	0xf5, 0x92, 0xe2, 0x76, 0x50, 0x77, 0x72, 0x05, 0x0f, 0xbb, 0x16, 0x4f, 0x5e, 0x4a, 0xda, 0xc5,
	0x11, 0x68, 0x9a, 0xd8, 0xa6, 0xa0, 0x37, 0xcf, 0x1e, 0x77, 0xf4, 0x2f, 0xdc, 0x3f, 0x21, 0x48,
	0x93, 0x46, 0x44, 0x3c, 0x25, 0xdf, 0xbc, 0x97, 0x4f, 0x78, 0x2f, 0x30, 0x10, 0x15, 0x97, 0x59,
	0x44, 0x5f, 0x08, 0x97, 0x05, 0xad, 0x23, 0xfe, 0x5a, 0x3f, 0x17, 0x76, 0xc4, 0xb2, 0x12, 0x4a,
	0xa0, 0x89, 0xf6, 0x60, 0xeb, 0xc1, 0xda, 0x33, 0x3d, 0xdf, 0x0b, 0xb1, 0x2f, 0x68, 0xa4, 0x3d,
	0x59, 0xf3, 0x14, 0x29, 0xc6, 0x69, 0xad, 0x08, 0x97, 0xe6, 0xda, 0x74, 0x66, 0xd0, 0xdf, 0xba,
	0x90, 0x8a, 0x89, 0xd2, 0x88, 0xc1, 0x27, 0x80, 0xee, 0xbd, 0x01, 0xa2, 0x05, 0x74, 0x9b, 0x9a,
	0x56, 0x29, 0xdb, 0x79, 0xc0, 0x07, 0xa1, 0x13, 0x9f, 0x7c, 0x1c, 0x0f, 0xf3, 0xde, 0x1d, 0xd0,
	0x8b, 0x0f, 0x92, 0x61, 0x2b, 0x6f, 0x77, 0x68, 0x01, 0x07, 0x25, 0xe1, 0xd4, 0xeb, 0xf9, 0x20,
	0x1c, 0xc7, 0xa7, 0xad, 0x3c, 0x34, 0x47, 0xed, 0xb6, 0xef, 0xbd, 0x81, 0x44, 0x4f, 0xe8, 0x02,
	0x3a, 0x94, 0x13, 0x56, 0x78, 0x7d, 0xed, 0x34, 0xbc, 0xd0, 0xf0, 0x6e, 0x40, 0x62, 0x44, 0x14,
	0x40, 0x47, 0x0a, 0x56, 0x2a, 0x6f, 0xe0, 0x83, 0x10, 0xc4, 0xff, 0x5b, 0xd9, 0x5d, 0x3a, 0x57,
	0xd7, 0x78, 0xb5, 0x4a, 0x8c, 0x84, 0x6e, 0x21, 0xcc, 0x2b, 0x4a, 0x14, 0xdd, 0xa5, 0x44, 0x79,
	0x8e, 0x0f, 0xc2, 0x7f, 0xcb, 0x29, 0x36, 0xd1, 0xb1, 0x8d, 0x86, 0x1f, 0x6c, 0xf4, 0x64, 0xdc,
	0xb9, 0x37, 0x6a, 0x3d, 0x7b, 0x3f, 0x1e, 0xe6, 0x10, 0x8e, 0x6c, 0x6f, 0xed, 0xe8, 0x22, 0x47,
	0xb7, 0x17, 0xe4, 0x70, 0xd2, 0xc5, 0xdf, 0x34, 0x4a, 0x6c, 0xcb, 0xbc, 0xa2, 0x9c, 0x96, 0x0a,
	0xf9, 0xbf, 0xbb, 0x70, 0xbb, 0x2e, 0x6c, 0x09, 0xeb, 0xcb, 0x96, 0xe3, 0xc1, 0xb3, 0x0e, 0x9b,
	0x92, 0x46, 0x89, 0x94, 0x59, 0xc0, 0xcf, 0x47, 0xe2, 0x05, 0x9c, 0xe5, 0x82, 0xe3, 0xbf, 0xbe,
	0x0f, 0xcb, 0xec, 0x71, 0xc4, 0xb8, 0x14, 0x95, 0x92, 0x59, 0x36, 0xd4, 0x49, 0x56, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xc6, 0xba, 0x2a, 0xc3, 0x0e, 0x02, 0x00, 0x00,
}
