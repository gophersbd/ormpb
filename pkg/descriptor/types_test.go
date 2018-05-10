package descriptor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoPackage_Standard(t *testing.T) {
	gp := &GoPackage{
		Path: "test.a",
	}
	assert.False(t, gp.Standard())

	gp.Path = "test/a"
	assert.True(t, gp.Standard())
}
