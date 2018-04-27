package dialect

import "testing"
import (
	protod "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/protobuf"
	"github.com/stretchr/testify/assert"
)

func TestType2SQLType(t *testing.T) {
	assert.Equal(t, type2SQLType(protod.FieldDescriptorProto_TYPE_DOUBLE, "").Name, Double)
	assert.Equal(t, type2SQLType(protod.FieldDescriptorProto_TYPE_FLOAT, "").Name, Float)
	assert.Equal(t, type2SQLType(protod.FieldDescriptorProto_TYPE_SINT64, "").Name, BigInt)
	assert.Equal(t, type2SQLType(protod.FieldDescriptorProto_TYPE_SINT32, "").Name, Int)
	assert.Equal(t, type2SQLType(protod.FieldDescriptorProto_TYPE_BOOL, "").Name, Bool)
	assert.Equal(t, type2SQLType(protod.FieldDescriptorProto_TYPE_STRING, "").Name, Varchar)
	assert.Equal(t, type2SQLType(protod.FieldDescriptorProto_TYPE_BYTES, "").Name, Text)
	assert.Equal(t, type2SQLType(protod.FieldDescriptorProto_TYPE_ENUM, "").Name, Text)
	assert.Equal(t, type2SQLType(protod.FieldDescriptorProto_TYPE_MESSAGE, ".google.protobuf.Timestamp").Name, Timestamp)
}

func TestSQLTypeFromTag(t *testing.T) {
	co := &protobuf.ColumnOptions{
		Type: "DOUBLE",
		Size: 128,
	}

	_, found := sqlTypeFromTag(co)
	assert.Equal(t, found, true)

	co.Type = ""
	_, found = sqlTypeFromTag(co)
	assert.Equal(t, found, false)
}
