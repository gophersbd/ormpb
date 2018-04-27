package generator

import (
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/generator/common"
	"github.com/gophersbd/ormpb/pkg/generator/migration"
	"github.com/gophersbd/ormpb/pkg/generator/orm"
)

// New return Generator interface
func New(reg *descriptor.Registry) []common.Generator {
	ormGenerator := orm.NewGenerator(reg)
	migrationGenerator := migration.NewGenerator(reg)

	return []common.Generator{ormGenerator, migrationGenerator}
}
