package validation

import (
	"testing"

	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/protobuf"
	"github.com/stretchr/testify/assert"
)

func TestValidateTableOptions(t *testing.T) {
	m := &descriptor.Message{}
	err := ValidateTableOptions(m)
	assert.NotNil(t, err)

	m.TableOptions = &protobuf.TableOptions{}
	err = ValidateTableOptions(m)
	assert.NotNil(t, err)

	m.TableOptions.Name = "examples"
	err = ValidateTableOptions(m)
	assert.NotNil(t, err)

	m.TableOptions.Type = "postgres"

	m.Fields = []*descriptor.Field{
		{
			Name: "name",
		},
	}
	m.TableOptions = &protobuf.TableOptions{
		Name: "examples",
		Type: "postgres",
	}
	err = ValidateTableOptions(m)
	assert.Nil(t, err)
}

func TestValidateColumnOptions(t *testing.T) {
	assert.Nil(t, ValidateColumnOptions(nil))
}
