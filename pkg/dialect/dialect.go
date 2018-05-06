package dialect

import (
	"fmt"
	"reflect"

	"sync"

	"github.com/gophersbd/ormpb/pkg/descriptor"
)

// Dialect interface contains behaviors that differ across SQL database
type Dialect interface {
	// ColumnSignatureOf return column's signature (data type & constraint)
	ColumnSignatureOf(field *descriptor.Field) string
}

var dialectsMap struct {
	v map[string]Dialect
	sync.Mutex
}

func init() {
	dialectsMap.v = make(map[string]Dialect)
}

// RegisterDialect register new dialect
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap.Lock()
	dialectsMap.v[name] = dialect
	dialectsMap.Unlock()
}

// NewDialect return registered Dialect
func NewDialect(name string) (Dialect, error) {
	value, ok := dialectsMap.v[name]
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

// ParseColumnSignature return SQLType & AdditionalType
func ParseColumnSignature(field *descriptor.Field) (sqlType SQLType, at AdditionalType) {
	column := field.Column
	sqlType, found := sqlTypeFromTag(column.Options)
	if !found {
		sqlType = type2SQLType(field.FieldDescriptorProto.GetType(), field.FieldDescriptorProto.GetTypeName())

		size := column.Options.GetSize()
		if size != 0 {
			sqlType.DefaultLength = int(size)
		}
	}

	at = AdditionalType{
		SetConstraint: make(map[Constraint]bool),
	}

	options := column.Options
	if options.GetNotNull() {
		at.SetConstraint[ConstraintNotNull] = true
	}
	if options.GetAutoIncrement() {
		at.SetConstraint[ConstraintAutoIncrement] = true
	}
	if options.GetPrimaryKey() {
		at.SetConstraint[ConstraintPrimaryKey] = true
	}
	if options.GetUnique() {
		at.SetConstraint[ConstraintUnique] = true
	}

	return
}
