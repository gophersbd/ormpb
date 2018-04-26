package validator

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/gophersbd/ormpb/pkg/descriptor"
)

// ValidateTableOptions validates TableOptions of a message
func ValidateTableOptions(m *descriptor.Message) error {
	glog.V(1).Infof("Validating message %s", m.GetName())
	if m.TableOptions == nil {
		return fmt.Errorf(`TableOptions not found in message "%s"`, m.GetName())
	}
	if m.TableOptions.GetName() == "" {
		return fmt.Errorf(`option "(ormpb.protobuf.table).name" is not set in message "%s"`, m.GetName())
	}

	for _, f := range m.Fields {
		if err := ValidateColumnOptions(f); err != nil {
			return err
		}
	}
	return nil
}

// ValidateColumnOptions validates ColumnOptions of a field
func ValidateColumnOptions(f *descriptor.Field) error {
	return nil
}