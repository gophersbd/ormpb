package dialect

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/gophersbd/ormpb/pkg/descriptor"
)

// Dialect interface contains behaviors that differ across SQL database
type Dialect interface {
	// GetUpSQL returns migration sql to create Table
	GetUpSQL(*descriptor.Message) (string, error)
	// GetDownSQL returns migration sql to delete Table
	GetDownSQL(*descriptor.Message) (string, error)
}

var dialectsRegistry struct {
	v map[string]Dialect
	sync.Mutex
}

func init() {
	dialectsRegistry.v = make(map[string]Dialect)
}

// RegisterDialect register new dialect
func RegisterDialect(name string, dialect Dialect) {
	dialectsRegistry.Lock()
	dialectsRegistry.v[name] = dialect
	dialectsRegistry.Unlock()
}

// NewDialect return registered Dialect
func NewDialect(name string) (Dialect, error) {
	dialectsRegistry.Lock()
	value, ok := dialectsRegistry.v[name]
	dialectsRegistry.Unlock()
	if !ok {
		return nil, fmt.Errorf("dialect not fount for %s", name)
	}
	dialect := reflect.New(reflect.TypeOf(value).Elem()).Interface().(Dialect)
	return dialect, nil
}
