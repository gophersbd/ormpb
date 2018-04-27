package migration

import (
	"bytes"
	"reflect"
	"text/template"
	"unicode"

	"github.com/gophersbd/ormpb/pkg/descriptor"
)

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

var fns = template.FuncMap{
	"not_last": func(x int, a interface{}) bool {
		return x != reflect.ValueOf(a).Len()-1
	},
}

type param struct {
	*descriptor.Message
}

func applyTemplateUp(p param) (string, error) {
	w := bytes.NewBuffer(nil)

	helperTemplate := template.New("migration")
	helperTemplate = helperTemplate.Funcs(template.FuncMap{"toSnake": ToSnake})
	helperTemplate = helperTemplate.Funcs(fns)

	migrationUpTemplate := template.Must(helperTemplate.Parse(migrationUpTemplate))

	if err := migrationUpTemplate.Execute(w, p); err != nil {
		return "", err
	}
	return w.String(), nil
}

func applyTemplateDown(p param) (string, error) {
	w := bytes.NewBuffer(nil)

	helperTemplate := template.New("migration")
	helperTemplate = helperTemplate.Funcs(fns)

	migrationDownTemplate := template.Must(helperTemplate.Parse(migrationDownTemplate))

	if err := migrationDownTemplate.Execute(w, p); err != nil {
		return "", err
	}
	return w.String(), nil
}

var (
	migrationUpTemplate = `
CREATE TABLE {{ .TableOptions.GetName }}(
{{- range $i, $f := .Fields }}
	{{ $f.ColumnName | toSnake }} {{ $f.ColumnConstraint }}{{if not_last $i $.Fields}},{{end}}
{{- end }}
);
`
	migrationDownTemplate = `
DROP TABLE IF EXISTS {{ .TableOptions.GetName }};
`
)