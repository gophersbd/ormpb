package postgres

import (
	"testing"

	protod "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/stretchr/testify/assert"
)

func TestType2SQLType(t *testing.T) {
	fType := protod.FieldDescriptorProto_TYPE_DOUBLE
	f := &descriptor.Field{
		FieldDescriptorProto: &protod.FieldDescriptorProto{
			Type: &fType,
		},
	}
	assert.Equal(t, Type2SQLType(f).Name, Numeric)

	fType = protod.FieldDescriptorProto_TYPE_FLOAT
	assert.Equal(t, Type2SQLType(f).Name, Numeric)
	fType = protod.FieldDescriptorProto_TYPE_SINT64
	assert.Equal(t, Type2SQLType(f).Name, BigInt)
	fType = protod.FieldDescriptorProto_TYPE_SINT32
	assert.Equal(t, Type2SQLType(f).Name, Integer)
	fType = protod.FieldDescriptorProto_TYPE_BOOL
	assert.Equal(t, Type2SQLType(f).Name, Bool)
	fType = protod.FieldDescriptorProto_TYPE_STRING
	assert.Equal(t, Type2SQLType(f).Name, Varchar)
	fType = protod.FieldDescriptorProto_TYPE_BYTES
	assert.Equal(t, Type2SQLType(f).Name, Text)
	fType = protod.FieldDescriptorProto_TYPE_ENUM
	assert.Equal(t, Type2SQLType(f).Name, Text)
	fType = protod.FieldDescriptorProto_TYPE_MESSAGE
	tName := ".google.protobuf.Timestamp"
	f.TypeName = &tName
	assert.Equal(t, Type2SQLType(f).Name, Timestamp)
}
