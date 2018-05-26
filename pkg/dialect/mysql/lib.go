package mysql

import (
	"fmt"
	"strings"

	protod "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/dialect"
	"github.com/gophersbd/ormpb/pkg/dialect/mysql/templates"
)

type mysql struct {
}

func init() {
	dialect.RegisterDialect("mysql", &mysql{})
}

func (p *mysql) GetUpSQL(message *descriptor.Message) (string, error) {
	param := &templates.UpParam{
		MessageName: message.GetName(),
		TableName:   dialect.ToSnake(message.TableOptions.GetName()),
		Lines:       make([]string, 0),
	}

	for _, f := range message.Fields {
		cn := f.Column.Options.GetName()
		if cn == "" {
			cn = f.GetName()
		}
		columnSignature := columnSignatureOf(f)
		line := fmt.Sprintf("%s %s", dialect.ToSnake(cn), columnSignature)
		param.Lines = append(param.Lines, line)
	}

	return templates.ApplyTemplateUp(param)
}

func (p *mysql) GetDownSQL(message *descriptor.Message) (string, error) {
	param := &templates.DownParam{
		MessageName: message.GetName(),
		TableName:   dialect.ToSnake(message.TableOptions.GetName()),
	}

	return templates.ApplyTemplateDown(param)
}

// columnSignatureOf returns signature of column
func columnSignatureOf(field *descriptor.Field) string {
	var sqlType, at = dialect.ParseColumnSignature(field, type2SQLType)

	switch sqlType.Name {
	case Longtext:
		size := sqlType.DefaultLength
		if size > 0 && size < 65532 {
			sqlType.Name = Varchar
			sqlType.DefaultLength = size
		}
	case Longblob:
		size := sqlType.DefaultLength
		if size > 0 && size < 65532 {
			sqlType.Name = Varbinary
			sqlType.DefaultLength = size
		}
	}

	var additionalType string

	if at.SetConstraint[dialect.ConstraintPrimaryKey] {
		additionalType = additionalType + " " + "PRIMARY KEY"
	}

	if at.SetConstraint[dialect.ConstraintNotNull] {
		additionalType = additionalType + " " + "NOT NULL"
	}

	if at.SetConstraint[dialect.ConstraintUnique] {
		additionalType = additionalType + " " + "UNIQUE"
	}

	if at.SetConstraint[dialect.ConstraintAutoIncrement] {
		additionalType = additionalType + " " + "AUTO_INCREMENT"
	}

	options := field.Column.Options
	d := options.GetDefault()
	if d != "" {
		additionalType = additionalType + " " + fmt.Sprintf("DEFAULT %v", d)
	}

	st := sqlType.Name
	if sqlType.DefaultLength != 0 {
		st = st + fmt.Sprintf("(%d)", sqlType.DefaultLength)
	}

	additionalType = strings.TrimSpace(additionalType)
	if additionalType == "" {
		return st
	}

	return fmt.Sprintf("%v %v", st, additionalType)
}

// type2SQLType converts Proto type to DB Data type
func type2SQLType(field *descriptor.Field) (st dialect.SQLType) {
	filedType := field.FieldDescriptorProto.GetType()
	typeName := field.FieldDescriptorProto.GetTypeName()
	switch filedType {
	case protod.FieldDescriptorProto_TYPE_DOUBLE,
		protod.FieldDescriptorProto_TYPE_FLOAT:
		st = dialect.SQLType{Name: Double, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_INT32,
		protod.FieldDescriptorProto_TYPE_SINT32,
		protod.FieldDescriptorProto_TYPE_SFIXED32:
		st = dialect.SQLType{Name: Int, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_INT64,
		protod.FieldDescriptorProto_TYPE_SINT64,
		protod.FieldDescriptorProto_TYPE_SFIXED64:
		st = dialect.SQLType{Name: Bigint, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_UINT32,
		protod.FieldDescriptorProto_TYPE_FIXED32:
		st = dialect.SQLType{Name: fmt.Sprintf("%s %s", Int, Unsigned), DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_UINT64,
		protod.FieldDescriptorProto_TYPE_FIXED64:
		st = dialect.SQLType{Name: fmt.Sprintf("%s %s", Bigint, Unsigned), DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_BOOL:
		st = dialect.SQLType{Name: Boolean, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_STRING:
		st = dialect.SQLType{Name: Longtext, DefaultLength: 512}
	case protod.FieldDescriptorProto_TYPE_BYTES:
		st = dialect.SQLType{Name: Longblob, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_MESSAGE:
		if typeName == ".google.protobuf.Timestamp" {
			st = dialect.SQLType{Name: Timestamp, DefaultLength: 0}
		}
	default:
		st = dialect.SQLType{Name: Longtext, DefaultLength: 0}
	}
	return
}
