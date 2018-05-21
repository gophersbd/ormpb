package postgres

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/pkg/dialect"
)

// Supported Data type
const (
	Int       = "INT"
	BigInt    = "BIGINT"
	Varchar   = "VARCHAR"
	Text      = "TEXT"
	Float     = "FLOAT"
	Double    = "DOUBLE"
	Numeric   = "NUMERIC"
	Bool      = "BOOL"
	Boolean   = "BOOLEAN"
	Serial    = "SERIAL"
	BigSerial = "BIGSERIAL"
	Integer   = "INTEGER"
	Timestamp = "TIMESTAMP"
)

// type2SQLType converts Proto type to DB Data type
func type2SQLType(filedType descriptor.FieldDescriptorProto_Type, typeName string) (st dialect.SQLType) {
	switch filedType {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		st = dialect.SQLType{Name: Double, DefaultLength: 0}
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		st = dialect.SQLType{Name: Float, DefaultLength: 0}
	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT64:
		st = dialect.SQLType{Name: BigInt, DefaultLength: 0}
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		st = dialect.SQLType{Name: Int, DefaultLength: 0}
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		st = dialect.SQLType{Name: Bool, DefaultLength: 0}
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		st = dialect.SQLType{Name: Varchar, DefaultLength: 255}
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		st = dialect.SQLType{Name: Text, DefaultLength: 0}
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		if typeName == ".google.protobuf.Timestamp" {
			st = dialect.SQLType{Name: Timestamp, DefaultLength: 0}
		}
	default:
		st = dialect.SQLType{Name: Text, DefaultLength: 0}
	}
	return
}
