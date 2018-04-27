package dialect

import (
	"testing"

	protod "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/protobuf"
	"github.com/stretchr/testify/assert"
)

func TestNewDialect(t *testing.T) {
	RegisterDialect("postgres", &postgres{})

	_, err := NewDialect("pg")
	assert.NotNil(t, err)

	_, err = NewDialect("postgres")
	assert.Nil(t, err)
}

func TestParseColumnSignature(t *testing.T) {
	fType := protod.FieldDescriptorProto_TYPE_FLOAT
	f := &descriptor.Field{
		FieldDescriptorProto: &protod.FieldDescriptorProto{
			Type: &fType,
		},
		Column: &descriptor.Column{
			Tags: map[string]interface{}{},
			Options: &protobuf.ColumnOptions{
				PrimaryKey:    true,
				AutoIncrement: true,
				Unique:        true,
				NotNull:       true,
				Size:          128,
			},
		},
	}

	sqlType, _ := ParseColumnSignature(f)
	assert.Equal(t, type2SQLType(fType, "").Name, sqlType.Name)
}
