package templates

import (
	"bytes"
	"text/template"

	"github.com/gophersbd/ormpb/pkg/dialect"
)

// DownParam holds information to be used in template
type DownParam struct {
	MessageName string
	TableName   string
}

// ApplyTemplateDown Execute template to create migration SQl
func ApplyTemplateDown(param *DownParam) (string, error) {
	w := bytes.NewBuffer(nil)

	helperTemplate := template.New("down")
	helperTemplate = helperTemplate.Funcs(dialect.GetTemplateFuncMap())
	migrationUpTemplate := template.Must(helperTemplate.Parse(migrationDownTemplate))

	if err := migrationUpTemplate.Execute(w, param); err != nil {
		return "", err
	}
	return w.String(), nil
}

var migrationDownTemplate = `/* Generated for {{ .MessageName }} */
DROP TABLE IF EXISTS {{ .TableName }};
`
