package templates

import (
	"bytes"
	"text/template"

	"github.com/gophersbd/ormpb/pkg/dialect"
)

// UpParam holds information to be used in template
type UpParam struct {
	MessageName string
	TableName   string
	Lines       []string
}

// ApplyTemplateUp Execute template to create migration SQl
func ApplyTemplateUp(param *UpParam) (string, error) {
	w := bytes.NewBuffer(nil)

	helperTemplate := template.New("up")
	helperTemplate = helperTemplate.Funcs(dialect.GetTemplateFuncMap())
	migrationUpTemplate := template.Must(helperTemplate.Parse(migrationUpTemplate))

	if err := migrationUpTemplate.Execute(w, param); err != nil {
		return "", err
	}
	return w.String(), nil
}

var migrationUpTemplate = `/* Generated for {{ .MessageName }} */
CREATE TABLE {{ .TableName }} (
{{- range $i, $f := .Lines }}
	{{ $f }}{{if not_last $i $.Lines}},{{end}}
{{- end }}
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
`
