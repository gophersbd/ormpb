package dialect

import (
	"fmt"
	"reflect"

	"github.com/gophersbd/ormpb/pkg/descriptor"
)

// Dialect interface contains behaviors that differ across SQL database
type Dialect interface {
	// DataTypeOf return data's sql type
	DataTypeOf(field *descriptor.Field) string
}

var dialectsMap = map[string]Dialect{}

// RegisterDialect register new dialect
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

// NewDialect return registered Dialect
func NewDialect(name string) (Dialect, error) {
	value, ok := dialectsMap[name]
	if !ok {
		return nil, fmt.Errorf("dialect not fount for %s", name)
	}
	dialect := reflect.New(reflect.TypeOf(value).Elem()).Interface().(Dialect)
	return dialect, nil
}

// Constraint for Column
type Constraint string

// Supported Column Constraint
const (
	ConstraintNotNull       Constraint = "NOT_NULL"
	ConstraintAutoIncrement Constraint = "AUTO_INCREMENT"
	ConstraintPrimaryKey    Constraint = "PRIMARY_KEY"
	ConstraintUnique        Constraint = "UNIQUE"
)

// AdditionalType to know which Constraint in added for a column
type AdditionalType struct {
	SetConstraint map[Constraint]bool
}

// ParseFieldStructForDialect return SQLType & AdditionalType
var ParseFieldStructForDialect = func(field *descriptor.Field, dialect Dialect) (sqlType SQLType, at AdditionalType) {
	sqlType, found := sqlTypeFromTag(field.ColumnOptions)
	if !found {
		sqlType = type2SQLType(*field.FieldDescriptorProto.Type)

		if field.ColumnOptions != nil {
			size := field.ColumnOptions.GetSize()
			if size != 0 {
				sqlType.DefaultLength = int(size)
			}
		}
	}

	at = AdditionalType{
		SetConstraint: make(map[Constraint]bool),
	}

	if field.ColumnOptions != nil {
		if field.ColumnOptions.GetNotNull() {
			at.SetConstraint[ConstraintNotNull] = true
		}
		if field.ColumnOptions.GetAutoIncrement() {
			at.SetConstraint[ConstraintAutoIncrement] = true
		}
		if field.ColumnOptions.GetPrimaryKey() {
			at.SetConstraint[ConstraintPrimaryKey] = true
		}
		if field.ColumnOptions.GetUnique() {
			at.SetConstraint[ConstraintUnique] = true
		}
	}

	return
}
