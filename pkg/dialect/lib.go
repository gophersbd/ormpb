package dialect

import (
	"reflect"
	"text/template"
	"unicode"

	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/protobuf"
)

// Constraint for Column
type Constraint string

// Supported Column Constraint
const (
	ConstraintNotNull       Constraint = "NOT_NULL"
	ConstraintAutoIncrement Constraint = "AUTO_INCREMENT"
	ConstraintPrimaryKey    Constraint = "PRIMARY_KEY"
	ConstraintUnique        Constraint = "UNIQUE"
)

// AdditionalType to know which constraint is added for a column
type AdditionalType struct {
	SetConstraint map[Constraint]bool
}

// ParseColumnSignature return SQLType & AdditionalType
func ParseColumnSignature(field *descriptor.Field, type2SQLType func(*descriptor.Field) SQLType) (sqlType SQLType, at AdditionalType) {
	column := field.Column
	sqlType, found := sqlTypeFromTag(column.Options)
	if !found {
		sqlType = type2SQLType(field)

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

// ToSnake converts column name to Snake case
func ToSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}
	return string(out)
}

// GetTemplateFuncMap returns functions to be used in template
func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"not_last": func(x int, a interface{}) bool {
			return x != reflect.ValueOf(a).Len()-1
		},
	}
}

// SQLType for Data type and Data size
type SQLType struct {
	Name          string
	DefaultLength int
}

// sqlTypeFromTag return SQLType from Tag
func sqlTypeFromTag(options *protobuf.ColumnOptions) (st SQLType, set bool) {
	t := options.GetType()
	if t != "" {
		return SQLType{
			t,
			int(options.GetSize()),
		}, true
	}
	return SQLType{}, false
}
