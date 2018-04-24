// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/option.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	protobuf/option.proto

It has these top-level messages:
	TableOptions
	ColumnOptions
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TableOptions struct {
	// table_name specifies the table name for the message
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *TableOptions) Reset()                    { *m = TableOptions{} }
func (m *TableOptions) String() string            { return proto.CompactTextString(m) }
func (*TableOptions) ProtoMessage()               {}
func (*TableOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TableOptions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ColumnOptions struct {
	// column_name specifies column name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// type specifies column data type
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// size specifies column size, default 255
	Size int32 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	// primary_key specifies column as primary key
	PrimaryKey bool `protobuf:"varint,4,opt,name=primary_key,json=primaryKey" json:"primary_key,omitempty"`
	// unique specifies column as unique
	Unique bool `protobuf:"varint,5,opt,name=unique" json:"unique,omitempty"`
	// default specifies column default value
	Default string `protobuf:"bytes,6,opt,name=default" json:"default,omitempty"`
	// not_null specifies column as NOT NULL
	NotNull bool `protobuf:"varint,7,opt,name=not_null,json=notNull" json:"not_null,omitempty"`
	// auto_increment specifies column auto incrementable or not
	AutoIncrement bool `protobuf:"varint,8,opt,name=auto_increment,json=autoIncrement" json:"auto_increment,omitempty"`
}

func (m *ColumnOptions) Reset()                    { *m = ColumnOptions{} }
func (m *ColumnOptions) String() string            { return proto.CompactTextString(m) }
func (*ColumnOptions) ProtoMessage()               {}
func (*ColumnOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ColumnOptions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ColumnOptions) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ColumnOptions) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ColumnOptions) GetPrimaryKey() bool {
	if m != nil {
		return m.PrimaryKey
	}
	return false
}

func (m *ColumnOptions) GetUnique() bool {
	if m != nil {
		return m.Unique
	}
	return false
}

func (m *ColumnOptions) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *ColumnOptions) GetNotNull() bool {
	if m != nil {
		return m.NotNull
	}
	return false
}

func (m *ColumnOptions) GetAutoIncrement() bool {
	if m != nil {
		return m.AutoIncrement
	}
	return false
}

var E_Table = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*TableOptions)(nil),
	Field:         10000001,
	Name:          "ormpb.protobuf.table",
	Tag:           "bytes,10000001,opt,name=table",
	Filename:      "protobuf/option.proto",
}

var E_Column = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*ColumnOptions)(nil),
	Field:         10000003,
	Name:          "ormpb.protobuf.column",
	Tag:           "bytes,10000003,opt,name=column",
	Filename:      "protobuf/option.proto",
}

func init() {
	proto.RegisterType((*TableOptions)(nil), "ormpb.protobuf.TableOptions")
	proto.RegisterType((*ColumnOptions)(nil), "ormpb.protobuf.ColumnOptions")
	proto.RegisterExtension(E_Table)
	proto.RegisterExtension(E_Column)
}

func init() { proto.RegisterFile("protobuf/option.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0x86, 0xe9, 0xf7, 0x6d, 0x5d, 0x3d, 0x73, 0xbb, 0x08, 0x28, 0x51, 0x1c, 0x2b, 0x03, 0x61,
	0x57, 0x1d, 0xe8, 0xdd, 0x2e, 0x15, 0x04, 0x11, 0x15, 0x8a, 0x28, 0x78, 0x33, 0xda, 0xed, 0x6c,
	0x04, 0xd3, 0xa4, 0xa6, 0xc9, 0x45, 0xbd, 0x13, 0x7f, 0x8f, 0xff, 0x49, 0xfc, 0x25, 0x92, 0x74,
	0x19, 0x4e, 0xc1, 0xbb, 0x73, 0xde, 0xf3, 0xf6, 0x69, 0x78, 0x60, 0xaf, 0x54, 0x52, 0xcb, 0xdc,
	0x2c, 0x27, 0xb2, 0xd4, 0x4c, 0x8a, 0xc4, 0xed, 0xa4, 0x2f, 0x55, 0x51, 0xe6, 0x89, 0x3f, 0x1e,
	0xc6, 0x2b, 0x29, 0x57, 0x1c, 0x27, 0x9b, 0xf6, 0x02, 0xab, 0xb9, 0x62, 0xa5, 0x96, 0xaa, 0x29,
	0x8d, 0x46, 0xb0, 0x7b, 0x97, 0xe5, 0x1c, 0x6f, 0x1d, 0xa6, 0x22, 0x04, 0x5a, 0x22, 0x2b, 0x90,
	0x06, 0x71, 0x30, 0xde, 0x49, 0xdd, 0x3c, 0xfa, 0x08, 0xa0, 0x77, 0x2e, 0xb9, 0x29, 0xc4, 0x1f,
	0x2d, 0x9b, 0xe9, 0xba, 0x44, 0xfa, 0xaf, 0xc9, 0xec, 0x6c, 0xb3, 0x8a, 0xbd, 0x20, 0xfd, 0x1f,
	0x07, 0xe3, 0x76, 0xea, 0x66, 0x32, 0x84, 0x6e, 0xa9, 0x58, 0x91, 0xa9, 0x7a, 0xf6, 0x84, 0x35,
	0x6d, 0xc5, 0xc1, 0x38, 0x4a, 0x61, 0x1d, 0x5d, 0x61, 0x4d, 0xf6, 0x21, 0x34, 0x82, 0x3d, 0x1b,
	0xa4, 0x6d, 0x77, 0x5b, 0x6f, 0x84, 0x42, 0x67, 0x81, 0xcb, 0xcc, 0x70, 0x4d, 0x43, 0xf7, 0x0f,
	0xbf, 0x92, 0x03, 0x88, 0x84, 0xd4, 0x33, 0x61, 0x38, 0xa7, 0x1d, 0xf7, 0x4d, 0x47, 0x48, 0x7d,
	0x63, 0x38, 0x27, 0xc7, 0xd0, 0xcf, 0x8c, 0x96, 0x33, 0x26, 0xe6, 0x0a, 0x0b, 0x14, 0x9a, 0x46,
	0xae, 0xd0, 0xb3, 0xe9, 0xa5, 0x0f, 0xa7, 0xf7, 0xd0, 0xd6, 0x56, 0x03, 0x19, 0x26, 0x8d, 0xb2,
	0x8d, 0xc3, 0xe4, 0x1a, 0xab, 0x2a, 0x5b, 0x79, 0x41, 0xf4, 0xf5, 0xfd, 0xd3, 0x3e, 0xb9, 0x7b,
	0x72, 0x94, 0x6c, 0xcb, 0x4e, 0xbe, 0x7b, 0x4c, 0x1b, 0xdc, 0xf4, 0x01, 0xc2, 0xb9, 0x33, 0x47,
	0x06, 0xbf, 0xc0, 0x17, 0x0c, 0xf9, 0xc2, 0x63, 0xdf, 0x3c, 0x76, 0xf0, 0x13, 0xbb, 0x65, 0x3e,
	0x5d, 0xe3, 0xce, 0xe0, 0x31, 0xf2, 0x8d, 0x3c, 0x74, 0xd3, 0xe9, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x0f, 0xf8, 0x29, 0x5e, 0x15, 0x02, 0x00, 0x00,
}
