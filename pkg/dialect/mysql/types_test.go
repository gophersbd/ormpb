package mysql

import (
	"fmt"
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
	fType = protod.FieldDescriptorProto_TYPE_SFIXED32
	assert.Equal(t, Type2SQLType(f).Name, Int)
	fType = protod.FieldDescriptorProto_TYPE_SFIXED64
	assert.Equal(t, Type2SQLType(f).Name, Bigint)
	fType = protod.FieldDescriptorProto_TYPE_FIXED32
	assert.Equal(t, Type2SQLType(f).Name, fmt.Sprintf("%s %s", Int, Unsigned))
	fType = protod.FieldDescriptorProto_TYPE_FIXED64
	assert.Equal(t, Type2SQLType(f).Name, fmt.Sprintf("%s %s", Bigint, Unsigned))
	fType = protod.FieldDescriptorProto_TYPE_BOOL
	assert.Equal(t, Type2SQLType(f).Name, Boolean)
	fType = protod.FieldDescriptorProto_TYPE_STRING
	assert.Equal(t, Type2SQLType(f).Name, Longtext)
	fType = protod.FieldDescriptorProto_TYPE_BYTES
	assert.Equal(t, Type2SQLType(f).Name, Longblob)
	fType = protod.FieldDescriptorProto_TYPE_MESSAGE
	tName := ".google.protobuf.Timestamp"
	f.TypeName = &tName
	assert.Equal(t, Type2SQLType(f).Name, Timestamp)
	fType = protod.FieldDescriptorProto_TYPE_ENUM
	assert.Equal(t, Type2SQLType(f).Name, Longtext)
}
