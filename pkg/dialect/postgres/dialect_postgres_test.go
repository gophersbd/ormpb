package postgres

import (
	"testing"

	protod "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/dialect"
	"github.com/gophersbd/ormpb/protobuf"
	"github.com/stretchr/testify/assert"
)

func TestPostgres_ColumnSignatureOf(t *testing.T) {

	_, err := dialect.NewDialect("pg")
	assert.Nil(t, err)

	d, err := dialect.NewDialect("postgres")
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

	signature := d.ColumnSignatureOf(f)
	assert.Equal(t, "NUMERIC(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	fType = protod.FieldDescriptorProto_TYPE_INT32
	f.FieldDescriptorProto.Type = &fType
	signature = d.ColumnSignatureOf(f)
	assert.Equal(t, "INTEGER(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	f.Column.Options.AutoIncrement = true
	signature = d.ColumnSignatureOf(f)
	assert.Equal(t, "SERIAL(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	fType = protod.FieldDescriptorProto_TYPE_INT64
	f.FieldDescriptorProto.Type = &fType
	signature = d.ColumnSignatureOf(f)
	assert.Equal(t, "BIGSERIAL(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	f.Column.Options.AutoIncrement = false
	signature = d.ColumnSignatureOf(f)
	assert.Equal(t, "BIGINT(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	fType = protod.FieldDescriptorProto_TYPE_STRING
	f.FieldDescriptorProto.Type = &fType
	signature = d.ColumnSignatureOf(f)
	assert.Equal(t, "VARCHAR(128) PRIMARY KEY NOT NULL UNIQUE", signature)

	f.Column.Options.Size = 65533
	signature = d.ColumnSignatureOf(f)
	assert.Equal(t, "TEXT PRIMARY KEY NOT NULL UNIQUE", signature)

	f.Column.Options.Default = "TEST"
	f.Column.Options.Unique = false
	signature = d.ColumnSignatureOf(f)
	assert.Equal(t, "TEXT PRIMARY KEY NOT NULL DEFAULT TEST", signature)

	f.Column.Options = nil
	signature = d.ColumnSignatureOf(f)
	assert.Equal(t, "VARCHAR(255)", signature)
}
