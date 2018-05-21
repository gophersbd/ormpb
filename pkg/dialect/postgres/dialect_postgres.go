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
	var sqlType, at = dialect.ParseColumnSignature(field, Type2SQLType)

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
