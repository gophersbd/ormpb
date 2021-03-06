package postgres

import (
	"fmt"
	"strings"

	protod "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/dialect"
	"github.com/gophersbd/ormpb/pkg/dialect/postgres/templates"
)

type postgres struct {
}

func init() {
	dialect.RegisterDialect("postgres", &postgres{})
	dialect.RegisterDialect("pg", &postgres{})
}

func (p *postgres) GetUpSQL(message *descriptor.Message) (string, error) {
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

func (p *postgres) GetDownSQL(message *descriptor.Message) (string, error) {
	param := &templates.DownParam{
		MessageName: message.GetName(),
		TableName:   dialect.ToSnake(message.TableOptions.GetName()),
	}

	return templates.ApplyTemplateDown(param)
}

// ColumnSignatureOf returns signature of column
func columnSignatureOf(field *descriptor.Field) string {
	var sqlType, at = dialect.ParseColumnSignature(field, type2SQLType)

	switch sqlType.Name {
	case Integer:
		if at.SetConstraint[dialect.ConstraintAutoIncrement] {
			sqlType.Name = Serial
		}
	case BigInt:
		if at.SetConstraint[dialect.ConstraintAutoIncrement] {
			sqlType.Name = BigSerial
		}
	case Varchar:
		size := sqlType.DefaultLength
		if !(size > 0 && size < 65532) {
			sqlType.Name = Text
			sqlType.DefaultLength = 0
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
		st = dialect.SQLType{Name: Numeric, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_INT64,
		protod.FieldDescriptorProto_TYPE_UINT64,
		protod.FieldDescriptorProto_TYPE_FIXED64,
		protod.FieldDescriptorProto_TYPE_SFIXED64,
		protod.FieldDescriptorProto_TYPE_SINT64:
		st = dialect.SQLType{Name: BigInt, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_INT32,
		protod.FieldDescriptorProto_TYPE_FIXED32,
		protod.FieldDescriptorProto_TYPE_UINT32,
		protod.FieldDescriptorProto_TYPE_SFIXED32,
		protod.FieldDescriptorProto_TYPE_SINT32:
		st = dialect.SQLType{Name: Integer, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_BOOL:
		st = dialect.SQLType{Name: Bool, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_STRING:
		st = dialect.SQLType{Name: Varchar, DefaultLength: 255}
	case protod.FieldDescriptorProto_TYPE_BYTES:
		st = dialect.SQLType{Name: Text, DefaultLength: 0}
	case protod.FieldDescriptorProto_TYPE_MESSAGE:
		if typeName == ".google.protobuf.Timestamp" {
			st = dialect.SQLType{Name: Timestamp, DefaultLength: 0}
		}
	default:
		st = dialect.SQLType{Name: Text, DefaultLength: 0}
	}
	return
}
