package generator

import (
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/generator/migration"
	"github.com/gophersbd/ormpb/pkg/generator/orm"
)

// Generator is an abstraction of code generators.
type Generator interface {
	// Generate generates output files from input .proto files.
	Generate(targets []*descriptor.File) ([]*plugin.CodeGeneratorResponse_File, error)
}

// New return Generator interface
func New(reg *descriptor.Registry) []Generator {
	ormGenerator := orm.NewGenerator(reg)
	migrationGenerator := migration.NewGenerator(reg)

	return []Generator{ormGenerator, migrationGenerator}
}
