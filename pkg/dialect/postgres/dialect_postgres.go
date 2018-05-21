package postgres

import (
	"fmt"
	"strings"

	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/dialect"
)

type postgres struct {
}

func init() {
	dialect.RegisterDialect("postgres", &postgres{})
	dialect.RegisterDialect("pg", &postgres{})
}

// ColumnSignatureOf returns signature of column
func (s *postgres) ColumnSignatureOf(field *descriptor.Field) string {
	var sqlType, at = parseColumnSignature(field)

	switch sqlType.Name {
	case Int:
		if at.setConstraint[dialect.ConstraintAutoIncrement] {
			sqlType.Name = Serial
		} else {
			sqlType.Name = Integer
		}
	case BigInt:
		if at.setConstraint[dialect.ConstraintAutoIncrement] {
			sqlType.Name = BigSerial
		} else {
			sqlType.Name = BigInt
		}
	case Float, Double:
		sqlType.Name = Numeric
	case Varchar:
		size := sqlType.DefaultLength
		if !(size > 0 && size < 65532) {
			sqlType.Name = Text
			sqlType.DefaultLength = 0
		}
	}

	var additionalType string

	if at.setConstraint[dialect.ConstraintPrimaryKey] {
		additionalType = additionalType + " " + "PRIMARY KEY"
	}

	if at.setConstraint[dialect.ConstraintNotNull] {
		additionalType = additionalType + " " + "NOT NULL"
	}

	if at.setConstraint[dialect.ConstraintUnique] {
		additionalType = additionalType + " " + "UNIQUE"
	}

	options := field.Column.Options
	d := options.GetDefault()
	if d != "" {
		additionalType = additionalType + fmt.Sprintf("DEFAULT %v", d)
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

// additionalType to know which constraint in added for a column
type additionalType struct {
	setConstraint map[dialect.Constraint]bool
}

// parseColumnSignature return SQLType & additionalType
func parseColumnSignature(field *descriptor.Field) (sqlType dialect.SQLType, at additionalType) {
	column := field.Column
	sqlType, found := dialect.SQLTypeFromTag(column.Options)
	if !found {
		sqlType = type2SQLType(field.FieldDescriptorProto.GetType(), field.FieldDescriptorProto.GetTypeName())

		size := column.Options.GetSize()
		if size != 0 {
			sqlType.DefaultLength = int(size)
		}
	}

	at = additionalType{
		setConstraint: make(map[dialect.Constraint]bool),
	}

	options := column.Options
	if options.GetNotNull() {
		at.setConstraint[dialect.ConstraintNotNull] = true
	}
	if options.GetAutoIncrement() {
		at.setConstraint[dialect.ConstraintAutoIncrement] = true
	}
	if options.GetPrimaryKey() {
		at.setConstraint[dialect.ConstraintPrimaryKey] = true
	}
	if options.GetUnique() {
		at.setConstraint[dialect.ConstraintUnique] = true
	}

	return
}
