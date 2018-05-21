package dialect

import (
	"testing"

	protod "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/protobuf"
	"github.com/stretchr/testify/assert"
)

type fake struct {
}

func (s *fake) ColumnSignatureOf(field *descriptor.Field) string {
	return ""
}

func init() {
	RegisterDialect("fake", &fake{})
	RegisterDialect("f", &fake{})
}

func TestNewDialect(t *testing.T) {
	RegisterDialect("fake", &fake{})

	_, err := NewDialect("f")
	assert.Nil(t, err)

	_, err = NewDialect("fake")
	assert.Nil(t, err)

	_, err = NewDialect("mdb")
	assert.NotNil(t, err)
}

// Type2SQLType converts Proto type to DB Data type
func Type2SQLType(field *descriptor.Field) (st SQLType) {
	filedType := field.FieldDescriptorProto.GetType()
	switch filedType {
	case protod.FieldDescriptorProto_TYPE_FLOAT:
		st = SQLType{Name: "FLOAT", DefaultLength: 0}
	}
	return
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

	sqlType, _ := ParseColumnSignature(f, Type2SQLType)
	assert.Equal(t, Type2SQLType(f).Name, sqlType.Name)
}
