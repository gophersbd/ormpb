package dialect

import (
	"testing"

	protod "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/protobuf"
	"github.com/stretchr/testify/assert"
)

func TestToSnake(t *testing.T) {
	assert.Equal(t, ToSnake("Name"), "name")
	assert.Equal(t, ToSnake("NameCheck"), "name_check")
	assert.Equal(t, ToSnake("aCheck"), "a_check")
}

func TestGetTemplateFuncMap(t *testing.T) {
	funcMap := GetTemplateFuncMap()
	notLast, found := funcMap["not_last"]
	assert.True(t, found)

	assert.True(t, notLast.(func(x int, a interface{}) bool)(0, []int{0, 1, 2}))
	assert.False(t, notLast.(func(x int, a interface{}) bool)(2, []int{0, 1, 2}))
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
