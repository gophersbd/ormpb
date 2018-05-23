package postgres

import (
	"testing"

	"github.com/golang/protobuf/proto"
	protod "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/dialect"
	"github.com/gophersbd/ormpb/protobuf"
	"github.com/stretchr/testify/assert"
)

func TestMysql_GetUpSQL(t *testing.T) {
	reg := descriptor.NewRegistry()
	src := `
		file_to_generate: 'example.proto'
		proto_file <
			name: 'example.proto'
			message_type <
				name: 'Example'
				options <
					[ormpb.protobuf.table] < name : "examples", type: "postgres" >
				>
				field <
					name: 'label'
					type: TYPE_STRING
				>
			>
		>
	`
	var req plugin.CodeGeneratorRequest
	err := proto.UnmarshalText(src, &req)
	assert.Nil(t, err)

	err = reg.Load(&req)
	assert.Nil(t, err)

	msg, err := reg.LookupMsg("", ".Example")
	assert.Nil(t, err)

	p := &postgres{}
	p.GetUpSQL(msg)
}

func TestMysql_GetDownSQL(t *testing.T) {
	reg := descriptor.NewRegistry()
	src := `
		file_to_generate: 'example.proto'
		proto_file <
			name: 'example.proto'
			message_type <
				name: 'Example'
				options <
					[ormpb.protobuf.table] < name : "examples", type: "postgres" >
				>
				field <
					name: 'label'
					type: TYPE_STRING
				>
			>
		>
	`
	var req plugin.CodeGeneratorRequest
	err := proto.UnmarshalText(src, &req)
	assert.Nil(t, err)

	err = reg.Load(&req)
	assert.Nil(t, err)

	msg, err := reg.LookupMsg("", ".Example")
	assert.Nil(t, err)

	p := &postgres{}
	_, err = p.GetDownSQL(msg)
	assert.Nil(t, err)
}

func Test_ColumnSignatureOf(t *testing.T) {

	_, err := dialect.NewDialect("pg")
	assert.Nil(t, err)

	_, err = dialect.NewDialect("postgres")
	assert.Nil(t, err)

	fType := protod.FieldDescriptorProto_TYPE_FLOAT
	f := &descriptor.Field{
		FieldDescriptorProto: &protod.FieldDescriptorProto{
			Type: &fType,
		},
		Column: &descriptor.Column{
			Tags: map[string]interface{}{},
			Options: &protobuf.ColumnOptions{
				PrimaryKey:    true,
				AutoIncrement: false,
				Unique:        true,
				NotNull:       true,
				Size:          128,
			},
		},
	}

	signature := columnSignatureOf(f)
	assert.Equal(t, "NUMERIC(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	fType = protod.FieldDescriptorProto_TYPE_INT32
	f.FieldDescriptorProto.Type = &fType
	signature = columnSignatureOf(f)
	assert.Equal(t, "INTEGER(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	f.Column.Options.AutoIncrement = true
	signature = columnSignatureOf(f)
	assert.Equal(t, "SERIAL(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	fType = protod.FieldDescriptorProto_TYPE_INT64
	f.FieldDescriptorProto.Type = &fType
	signature = columnSignatureOf(f)
	assert.Equal(t, "BIGSERIAL(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	f.Column.Options.AutoIncrement = false
	signature = columnSignatureOf(f)
	assert.Equal(t, "BIGINT(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	fType = protod.FieldDescriptorProto_TYPE_STRING
	f.FieldDescriptorProto.Type = &fType
	signature = columnSignatureOf(f)
	assert.Equal(t, "VARCHAR(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	f.Column.Options.Size = 65533
	signature = columnSignatureOf(f)
	assert.Equal(t, "TEXT PRIMARY KEY NOT NULL UNIQUE", signature)

	f.Column.Options.Default = "TEST"
	f.Column.Options.Unique = false
	signature = columnSignatureOf(f)
	assert.Equal(t, "TEXT PRIMARY KEY NOT NULL DEFAULT TEST", signature)

	f.Column.Options = nil
	signature = columnSignatureOf(f)
	assert.Equal(t, "VARCHAR(255)", signature)
}

func TestType2SQLType(t *testing.T) {
	fType := protod.FieldDescriptorProto_TYPE_DOUBLE
	f := &descriptor.Field{
		FieldDescriptorProto: &protod.FieldDescriptorProto{
			Type: &fType,
		},
	}
	assert.Equal(t, type2SQLType(f).Name, Numeric)

	fType = protod.FieldDescriptorProto_TYPE_FLOAT
	assert.Equal(t, type2SQLType(f).Name, Numeric)
	fType = protod.FieldDescriptorProto_TYPE_SINT64
	assert.Equal(t, type2SQLType(f).Name, BigInt)
	fType = protod.FieldDescriptorProto_TYPE_SINT32
	assert.Equal(t, type2SQLType(f).Name, Integer)
	fType = protod.FieldDescriptorProto_TYPE_BOOL
	assert.Equal(t, type2SQLType(f).Name, Bool)
	fType = protod.FieldDescriptorProto_TYPE_STRING
	assert.Equal(t, type2SQLType(f).Name, Varchar)
	fType = protod.FieldDescriptorProto_TYPE_BYTES
	assert.Equal(t, type2SQLType(f).Name, Text)
	fType = protod.FieldDescriptorProto_TYPE_ENUM
	assert.Equal(t, type2SQLType(f).Name, Text)
	fType = protod.FieldDescriptorProto_TYPE_MESSAGE
	tName := ".google.protobuf.Timestamp"
	f.TypeName = &tName
	assert.Equal(t, type2SQLType(f).Name, Timestamp)
}
