package dialect

import (
	"fmt"
	"strings"

	"github.com/gophersbd/ormpb/pkg/descriptor"
)

type postgres struct {
}

func init() {
	RegisterDialect("postgres", &postgres{})
}

func (s *postgres) DataTypeOf(field *descriptor.Field) string {
	var sqlType, at = ParseFieldStructForDialect(field, s)

	switch sqlType.Name {
	case Int:
		if at.SetConstraint[ConstraintAutoIncrement] {
			sqlType.Name = Serial
		} else {
			sqlType.Name = Integer
		}
	case BigInt:
		if at.SetConstraint[ConstraintAutoIncrement] {
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

	if at.SetConstraint[ConstraintPrimaryKey] {
		additionalType = additionalType + " " + "PRIMARY KEY"
	}

	if at.SetConstraint[ConstraintNotNull] {
		additionalType = additionalType + " " + "NOT NULL"
	}

	if at.SetConstraint[ConstraintUnique] {
		additionalType = additionalType + " " + "UNIQUE"
	}

	if field.ColumnOptions != nil {
		d := field.ColumnOptions.GetDefault()
		if d != "" {
			additionalType = additionalType + fmt.Sprintf("DEFAULT %v", d)
		}
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
