package mysql

import (
	"fmt"
	"strings"

	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/dialect"
)

type mysql struct {
}

func init() {
	dialect.RegisterDialect("mysql", &mysql{})
}

// ColumnSignatureOf returns signature of column
func (s *mysql) ColumnSignatureOf(field *descriptor.Field) string {
	var sqlType, at = dialect.ParseColumnSignature(field, Type2SQLType)

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
