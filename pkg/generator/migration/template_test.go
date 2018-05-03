package migration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSnake(t *testing.T) {
	assert.Equal(t, ToSnake("Name"), "name")
	assert.Equal(t, ToSnake("NameCheck"), "name_check")
	assert.Equal(t, ToSnake("aCheck"), "a_check")
}
