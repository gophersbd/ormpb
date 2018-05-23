package dialect

import (
	"testing"

	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/stretchr/testify/assert"
)

type fake struct {
}

func (s *fake) ColumnSignatureOf(field *descriptor.Field) string {
	return ""
}
func (s *fake) GetUpSQL(message *descriptor.Message) (string, error) {
	return "", nil
}
func (s *fake) GetDownSQL(message *descriptor.Message) (string, error) {
	return "", nil
}

func init() {
	RegisterDialect("fake", &fake{})
	RegisterDialect("f", &fake{})
}

func TestNewDialect(t *testing.T) {
	RegisterDialect("fake", &fake{})

	_, err := NewDialect("f")
	assert.Nil(t, err)

	_, err = NewDialect("fake")
	assert.Nil(t, err)

	_, err = NewDialect("mdb")
	assert.NotNil(t, err)
}
