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

// SQLType for Data type and Data size
type SQLType struct {
	Name          string
	DefaultLength int
}

// type2SQLType converts Proto type to DB Data type
func type2SQLType(t descriptor.FieldDescriptorProto_Type) (st SQLType) {
	switch t {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		st = SQLType{Double, 0}
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		st = SQLType{Float, 0}
	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT64:
		st = SQLType{BigInt, 0}
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		st = SQLType{Int, 0}
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		st = SQLType{Bool, 0}
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		st = SQLType{Varchar, 255}
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		st = SQLType{Text, 0}
	default:
		st = SQLType{Text, 0}
	}
	return
}

// sqlTypeFromTag return SQLType from Tag
func sqlTypeFromTag(options *protobuf.ColumnOptions) (st SQLType, set bool) {
	if options == nil {
		return SQLType{}, false
	}

	t := options.GetType()
	if t != "" {
		return SQLType{
			options.GetType(),
			int(options.GetSize()),
		}, true
	}

	return SQLType{}, false
}
