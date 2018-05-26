package migration

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/dialect"
	"github.com/gophersbd/ormpb/pkg/validation"
)

// Generator holds Registry
type Generator struct {
	reg *descriptor.Registry
}

// NewGenerator return Generator interface
func NewGenerator(reg *descriptor.Registry) *Generator {
	return &Generator{reg: reg}
}

// Generate receives target files and returns generated migration files
func (g *Generator) Generate(targets []*descriptor.File) ([]*plugin.CodeGeneratorResponse_File, error) {
	mTime := time.Now().Format("20060102")

	var files []*plugin.CodeGeneratorResponse_File
	for _, file := range targets {
		if file.MigrationDir == "" {
			continue
		}

		comment, err := applyTemplateComment(file)
		if err != nil {
			return nil, err
		}

		fileName := filepath.Base(file.GetName())

		for _, m := range file.Messages {
			if err := validation.ValidateTableOptions(m); err != nil {
				return nil, err
			}

			if err := validation.ValidateColumnOptions(m.Fields); err != nil {
				return nil, err
			}

			dbType := m.TableOptions.GetType()
			d, err := dialect.NewDialect(dbType)
			if err != nil {
				return nil, err
			}

			generatedUpSQL, err := d.GetUpSQL(m)
			if err != nil {
				return nil, err
			}
			upMigration := []string{
				comment,
				generatedUpSQL,
			}

			generatedDownSQL, err := d.GetDownSQL(m)
			if err != nil {
				return nil, err
			}
			downMigration := []string{
				comment,
				generatedDownSQL,
			}

			name := m.TableOptions.GetName()
			files = append(files, &plugin.CodeGeneratorResponse_File{
				Name:    proto.String(fmt.Sprintf("%s/%s", file.MigrationDir, fmt.Sprintf("%s_%s_up.sql", mTime, name))),
				Content: proto.String(strings.Join(upMigration, "\n\n")),
			})
			glog.V(1).Infof("Will emit %s", fileName)

			files = append(files, &plugin.CodeGeneratorResponse_File{
				Name:    proto.String(fmt.Sprintf("%s/%s", file.MigrationDir, fmt.Sprintf("%s_%s_down.sql", mTime, name))),
				Content: proto.String(strings.Join(downMigration, "\n\n")),
			})
			glog.V(1).Infof("Will emit %s", fileName)
		}

	}
	return files, nil
}
