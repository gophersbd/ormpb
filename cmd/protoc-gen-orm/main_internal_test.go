package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	assert.Equal(t, "Hello I am going to be ORM. Please Code me up", Dummy())
}
