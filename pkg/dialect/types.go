package dialect

import (
	"github.com/gophersbd/ormpb/protobuf"
)

// SQLType for Data type and Data size
type SQLType struct {
	Name          string
	DefaultLength int
}

// sqlTypeFromTag return SQLType from Tag
func sqlTypeFromTag(options *protobuf.ColumnOptions) (st SQLType, set bool) {
	t := options.GetType()
	if t != "" {
		return SQLType{
			t,
			int(options.GetSize()),
		}, true
	}
	return SQLType{}, false
}
