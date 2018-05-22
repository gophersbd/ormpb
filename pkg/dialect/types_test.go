package dialect

import (
	"testing"

	"github.com/gophersbd/ormpb/protobuf"
	"github.com/stretchr/testify/assert"
)

func TestSQLTypeFromTag(t *testing.T) {
	co := &protobuf.ColumnOptions{
		Type: "DOUBLE",
		Size: 128,
	}

	_, found := sqlTypeFromTag(co)
	assert.Equal(t, found, true)

	co.Type = ""
	_, found = sqlTypeFromTag(co)
	assert.Equal(t, found, false)
}
