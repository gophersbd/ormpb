package postgres

import (
	protod "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/dialect"
)

// Supported Data type
const (
	BigInt    = "BIGINT"
	Varchar   = "VARCHAR"
	Text      = "TEXT"
	Numeric   = "NUMERIC"
	Bool      = "BOOL"
	Boolean   = "BOOLEAN"
	Serial    = "SERIAL"
	BigSerial = "BIGSERIAL"
	Integer   = "INTEGER"
	Timestamp = "TIMESTAMP"
)

// Type2SQLType converts Proto type to DB Data type
func Type2SQLType(field *descriptor.Field) (st dialect.SQLType) {
	filedType := field.FieldDescriptorProto.GetType()
	typeName := field.FieldDescriptorProto.GetTypeName()
	switch filedType {
	case protod.FieldDescriptorProto_TYPE_DOUBLE,
		protod.FieldDescriptorProto_TYPE_FLOAT:
		st = dialect.SQLType{Name: Numeric, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_INT64,
		protod.FieldDescriptorProto_TYPE_UINT64,
		protod.FieldDescriptorProto_TYPE_FIXED64,
		protod.FieldDescriptorProto_TYPE_SFIXED64,
		protod.FieldDescriptorProto_TYPE_SINT64:
		st = dialect.SQLType{Name: BigInt, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_INT32,
		protod.FieldDescriptorProto_TYPE_FIXED32,
		protod.FieldDescriptorProto_TYPE_UINT32,
		protod.FieldDescriptorProto_TYPE_SFIXED32,
		protod.FieldDescriptorProto_TYPE_SINT32:
		st = dialect.SQLType{Name: Integer, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_BOOL:
		st = dialect.SQLType{Name: Bool, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_STRING:
		st = dialect.SQLType{Name: Varchar, DefaultLength: 255}
	case protod.FieldDescriptorProto_TYPE_BYTES:
		st = dialect.SQLType{Name: Text, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_MESSAGE:
		if typeName == ".google.protobuf.Timestamp" {
			st = dialect.SQLType{Name: Timestamp, DefaultLength: 0}
		}
	default:
		st = dialect.SQLType{Name: Text, DefaultLength: 0}
	}
	return
}
