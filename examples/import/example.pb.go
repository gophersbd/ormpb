// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ormpb/examples/import/example.proto

/*
Package importpb is a generated protocol buffer package.

It is generated from these files:
	ormpb/examples/import/example.proto

It has these top-level messages:
	Example
*/
package importpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Example) Reset()                    { *m = Example{} }
func (m *Example) String() string            { return proto.CompactTextString(m) }
func (*Example) ProtoMessage()               {}
func (*Example) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Example) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Example)(nil), "ormpb.examples.import.Example")
}

func init() { proto.RegisterFile("ormpb/examples/import/example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0x2f, 0xca, 0x2d,
	0x48, 0xd2, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0xcf, 0xcc, 0x2d, 0xc8, 0x2f,
	0x2a, 0x81, 0xf1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0xc1, 0x8a, 0xf4, 0x60, 0x8a,
	0xf4, 0x20, 0x8a, 0xa4, 0xa4, 0x21, 0x7a, 0xc1, 0x6a, 0x92, 0x4a, 0xd3, 0xf4, 0xf3, 0x0b, 0x4a,
	0x32, 0xf3, 0xf3, 0x20, 0x7a, 0x94, 0x9c, 0xb9, 0xd8, 0x5d, 0x21, 0xea, 0x85, 0xd4, 0xb9, 0x58,
	0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0x84, 0x67, 0xbd, 0x98, 0xa4,
	0xc6, 0x06, 0x11, 0x02, 0x31, 0x99, 0x25, 0x1a, 0x18, 0x83, 0xc0, 0x3c, 0x2b, 0xfe, 0xae, 0x17,
	0x93, 0xd4, 0xb8, 0xb8, 0x38, 0x60, 0x16, 0x39, 0x69, 0x70, 0xc9, 0x24, 0xe7, 0xe7, 0xea, 0x61,
	0xb5, 0x5e, 0xaf, 0x20, 0x29, 0x8a, 0x03, 0xc2, 0x2c, 0x48, 0x4a, 0x62, 0x03, 0xdb, 0x6a, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x19, 0xe5, 0xd7, 0xdc, 0xd0, 0x00, 0x00, 0x00,
}
