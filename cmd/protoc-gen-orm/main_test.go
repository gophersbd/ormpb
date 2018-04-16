package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainNotNil(t *testing.T) {
	assert.NotNil(t, main)
}
