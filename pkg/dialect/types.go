package dialect

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/protobuf"
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
)

// sqlType for Data type and Data size
type sqlType struct {
	Name          string
	DefaultLength int
}

// type2SQLType converts Proto type to DB Data type
func type2SQLType(t descriptor.FieldDescriptorProto_Type) (st sqlType) {
	switch t {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		st = sqlType{Double, 0}
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		st = sqlType{Float, 0}
	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT64:
		st = sqlType{BigInt, 0}
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		st = sqlType{Int, 0}
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		st = sqlType{Bool, 0}
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		st = sqlType{Varchar, 255}
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		st = sqlType{Text, 0}
	default:
		st = sqlType{Text, 0}
	}
	return
}

// sqlTypeFromTag return sqlType from Tag
func sqlTypeFromTag(options *protobuf.ColumnOptions) (st sqlType, set bool) {
	if options == nil {
		return sqlType{}, false
	}

	t := options.GetType()
	if t != "" {
		return sqlType{
			options.GetType(),
			int(options.GetSize()),
		}, true
	}

	return sqlType{}, false
}
