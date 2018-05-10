package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintInterface(t *testing.T) {
	assert.Equal(t, printInterface(5), `"5"`)
	assert.Equal(t, printInterface(5.2), `"5.2"`)
	assert.Equal(t, printInterface(true), `"true"`)
	assert.Equal(t, printInterface("check"), `"check"`)
}
