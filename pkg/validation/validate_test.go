package validation

import (
	"testing"

	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/protobuf"
)

func TestValidateTableOptions(t *testing.T) {
	m := &descriptor.Message{}
	err := ValidateTableOptions(m)
	if err == nil {
		t.Error("Validating TableOptions should fail")
	}

	m.TableOptions = &protobuf.TableOptions{}
	err = ValidateTableOptions(m)
	if err == nil {
		t.Error("Validating TableOptions should fail. Missing name in TableOptions")
	}

	m.Fields = []*descriptor.Field{
		{
			Name: "name",
		},
	}
	m.TableOptions = &protobuf.TableOptions{
		Name: "examples",
	}
	err = ValidateTableOptions(m)
	if err != nil {
		t.Error("Validating TableOptions should not fail.")
	}
}
