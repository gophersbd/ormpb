// Code generated by protoc-gen-go. DO NOT EDIT.
// source: option.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type TableOptions struct {
	// table_name specifies the table name for the message
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableOptions) Reset()         { *m = TableOptions{} }
func (m *TableOptions) String() string { return proto.CompactTextString(m) }
func (*TableOptions) ProtoMessage()    {}
func (*TableOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_6845bf9f693e8c83, []int{0}
}

func (m *TableOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableOptions.Unmarshal(m, b)
}
func (m *TableOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableOptions.Marshal(b, m, deterministic)
}
func (m *TableOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableOptions.Merge(m, src)
}
func (m *TableOptions) XXX_Size() int {
	return xxx_messageInfo_TableOptions.Size(m)
}
func (m *TableOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TableOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TableOptions proto.InternalMessageInfo

func (m *TableOptions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TableOptions) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ColumnOptions struct {
	// column_name specifies column name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// type specifies column data type
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// size specifies column size, default 255
	Size int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// primary_key specifies column as primary key
	PrimaryKey bool `protobuf:"varint,4,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	// unique specifies column as unique
	Unique bool `protobuf:"varint,5,opt,name=unique,proto3" json:"unique,omitempty"`
	// default specifies column default value
	Default string `protobuf:"bytes,6,opt,name=default,proto3" json:"default,omitempty"`
	// not_null specifies column as NOT NULL
	NotNull bool `protobuf:"varint,7,opt,name=not_null,json=notNull,proto3" json:"not_null,omitempty"`
	// auto_increment specifies column auto incrementable or not
	AutoIncrement        bool     `protobuf:"varint,8,opt,name=auto_increment,json=autoIncrement,proto3" json:"auto_increment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ColumnOptions) Reset()         { *m = ColumnOptions{} }
func (m *ColumnOptions) String() string { return proto.CompactTextString(m) }
func (*ColumnOptions) ProtoMessage()    {}
func (*ColumnOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_6845bf9f693e8c83, []int{1}
}

func (m *ColumnOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnOptions.Unmarshal(m, b)
}
func (m *ColumnOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnOptions.Marshal(b, m, deterministic)
}
func (m *ColumnOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnOptions.Merge(m, src)
}
func (m *ColumnOptions) XXX_Size() int {
	return xxx_messageInfo_ColumnOptions.Size(m)
}
func (m *ColumnOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnOptions proto.InternalMessageInfo

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
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*TableOptions)(nil),
	Field:         10000001,
	Name:          "ormpb.protobuf.table",
	Tag:           "bytes,10000001,opt,name=table",
	Filename:      "option.proto",
}

var E_Column = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*ColumnOptions)(nil),
	Field:         10000003,
	Name:          "ormpb.protobuf.column",
	Tag:           "bytes,10000003,opt,name=column",
	Filename:      "option.proto",
}

func init() {
	proto.RegisterType((*TableOptions)(nil), "ormpb.protobuf.TableOptions")
	proto.RegisterType((*ColumnOptions)(nil), "ormpb.protobuf.ColumnOptions")
	proto.RegisterExtension(E_Table)
	proto.RegisterExtension(E_Column)
}

func init() { proto.RegisterFile("option.proto", fileDescriptor_6845bf9f693e8c83) }

var fileDescriptor_6845bf9f693e8c83 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0xc9, 0xf7, 0xb5, 0x69, 0x3c, 0xfd, 0x59, 0xcc, 0x42, 0x46, 0xb1, 0x34, 0x14, 0x84,
	0xae, 0x52, 0x50, 0x70, 0xd1, 0xa5, 0x82, 0x20, 0xa2, 0x42, 0x10, 0x05, 0x37, 0x25, 0x69, 0x4f,
	0xcb, 0xe0, 0x64, 0x26, 0x4e, 0x66, 0x16, 0x71, 0x27, 0x5e, 0x8f, 0xf7, 0x24, 0x5e, 0x89, 0xcc,
	0xa4, 0x23, 0x56, 0x57, 0xee, 0xde, 0xf3, 0xe6, 0x9c, 0x27, 0xe1, 0x09, 0xf4, 0x64, 0xa9, 0x99,
	0x14, 0x49, 0xa9, 0xa4, 0x96, 0x64, 0x20, 0x55, 0x51, 0xe6, 0xcd, 0x90, 0x9b, 0xd5, 0x7e, 0xbc,
	0x96, 0x72, 0xcd, 0x71, 0xea, 0x8b, 0xe9, 0x12, 0xab, 0x85, 0x62, 0xa5, 0x96, 0xaa, 0x59, 0x1a,
	0x9f, 0x40, 0xef, 0x36, 0xcb, 0x39, 0xde, 0x38, 0x4c, 0x45, 0x08, 0xb4, 0x44, 0x56, 0x20, 0x0d,
	0xe2, 0x60, 0xb2, 0x93, 0xba, 0x6c, 0x3b, 0x5d, 0x97, 0x48, 0xff, 0x35, 0x9d, 0xcd, 0xe3, 0xf7,
	0x00, 0xfa, 0x67, 0x92, 0x9b, 0x42, 0xfc, 0xf1, 0xd2, 0x76, 0x15, 0x7b, 0x46, 0xfa, 0x3f, 0x0e,
	0x26, 0xed, 0xd4, 0x65, 0x32, 0x82, 0x6e, 0xa9, 0x58, 0x91, 0xa9, 0x7a, 0xfe, 0x88, 0x35, 0x6d,
	0xc5, 0xc1, 0x24, 0x4a, 0x61, 0x53, 0x5d, 0x62, 0x4d, 0x76, 0x21, 0x34, 0x82, 0x3d, 0x19, 0xa4,
	0x6d, 0xf7, 0x6c, 0x33, 0x11, 0x0a, 0x9d, 0x25, 0xae, 0x32, 0xc3, 0x35, 0x0d, 0xdd, 0x3b, 0xfc,
	0x48, 0xf6, 0x20, 0x12, 0x52, 0xcf, 0x85, 0xe1, 0x9c, 0x76, 0xdc, 0x4d, 0x47, 0x48, 0x7d, 0x6d,
	0x38, 0x27, 0x87, 0x30, 0xc8, 0x8c, 0x96, 0x73, 0x26, 0x16, 0x0a, 0x0b, 0x14, 0x9a, 0x46, 0x6e,
	0xa1, 0x6f, 0xdb, 0x0b, 0x5f, 0xce, 0xee, 0xa0, 0xad, 0xad, 0x1a, 0x32, 0x4a, 0x1a, 0x8d, 0x5f,
	0x5e, 0x93, 0x2b, 0xac, 0xaa, 0x6c, 0xed, 0xa5, 0xd1, 0x97, 0xb7, 0x0f, 0xfb, 0xc9, 0xdd, 0xa3,
	0x83, 0x64, 0xfb, 0x07, 0x24, 0xdf, 0xdd, 0xa6, 0x0d, 0x6e, 0x76, 0x0f, 0xe1, 0xc2, 0x99, 0x23,
	0xc3, 0x5f, 0xe0, 0x73, 0x86, 0x7c, 0xe9, 0xb1, 0xaf, 0x1e, 0x3b, 0xfc, 0x89, 0xdd, 0x32, 0x9f,
	0x6e, 0x70, 0xa7, 0xf0, 0x10, 0xf9, 0x8d, 0x3c, 0x74, 0xe9, 0xf8, 0x33, 0x00, 0x00, 0xff, 0xff,
	0x21, 0x0e, 0xcc, 0xea, 0x20, 0x02, 0x00, 0x00,
}
