package mysql

import (
	"fmt"

	protod "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/dialect"
)

// Supported Data type
const (
	Numeric   = "NUMERIC"
	Int       = "INT"
	Varchar   = "VARCHAR"
	Varbinary = "VARBINARY"
	Bigint    = "BIGINT"
	Unsigned  = "UNSIGNED"
	Longtext  = "LONGTEXT"
	Longblob  = "LONGBLOB"
	Boolean   = "BOOLEAN"
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
	case protod.FieldDescriptorProto_TYPE_INT32,
		protod.FieldDescriptorProto_TYPE_SINT32,
		protod.FieldDescriptorProto_TYPE_SFIXED32:
		st = dialect.SQLType{Name: Int, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_INT64,
		protod.FieldDescriptorProto_TYPE_SINT64,
		protod.FieldDescriptorProto_TYPE_SFIXED64:
		st = dialect.SQLType{Name: Bigint, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_UINT32,
		protod.FieldDescriptorProto_TYPE_FIXED32:
		st = dialect.SQLType{Name: fmt.Sprintf("%s %s", Int, Unsigned), DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_UINT64,
		protod.FieldDescriptorProto_TYPE_FIXED64:
		st = dialect.SQLType{Name: fmt.Sprintf("%s %s", Bigint, Unsigned), DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_BOOL:
		st = dialect.SQLType{Name: Boolean, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_STRING:
		st = dialect.SQLType{Name: Longtext, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_BYTES:
		st = dialect.SQLType{Name: Longblob, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_MESSAGE:
		if typeName == ".google.protobuf.Timestamp" {
			st = dialect.SQLType{Name: Timestamp, DefaultLength: 0}
		}
	default:
		st = dialect.SQLType{Name: Longtext, DefaultLength: 0}
	}
	return
}
