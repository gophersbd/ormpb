package migration

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/dialect"
	"github.com/gophersbd/ormpb/pkg/generator/common"
	"github.com/gophersbd/ormpb/pkg/validation"
)

type generator struct {
	reg *descriptor.Registry
}

// NewGenerator return Generator interface
func NewGenerator(reg *descriptor.Registry) common.Generator {
	return &generator{reg: reg}
}

func (g *generator) Generate(targets []*descriptor.File) ([]*plugin.CodeGeneratorResponse_File, error) {
	var files []*plugin.CodeGeneratorResponse_File
	for _, file := range targets {

		for _, m := range file.Messages {
			if err := validation.ValidateTableOptions(m); err != nil {
				return nil, err
			}

			if err := validation.ValidateColumnOptions(m.Fields); err != nil {
				return nil, err
			}

			msgName := m.TableOptions.GetName()
			dbType := m.TableOptions.GetType()
			d, err := dialect.NewDialect(dbType)
			if err != nil {
				return nil, err
			}
			for _, f := range m.Fields {
				fc := d.DataTypeOf(f)
				f.ColumnConstraint = fc
				cn := f.ColumnOptions.GetName()
				if cn == "" {
					cn = f.GetName()
				}
				f.ColumnName = cn
			}

			mTime := time.Now().Unix()

			{
				code, err := g.generateUp(m)
				if err != nil {
					return nil, err
				}

				fileName := file.GetName()
				if file.GoPkg.Path != "" {
					fileName = fmt.Sprintf("%s/%s", file.GoPkg.Path, fmt.Sprintf("%d_%s_up.sql", mTime, msgName))
				}

				files = append(files, &plugin.CodeGeneratorResponse_File{
					Name:    proto.String(fileName),
					Content: proto.String(code),
				})
				glog.V(1).Infof("Will emit %s", fileName)
			}

			{
				code, err := g.generateDown(m)
				if err != nil {
					return nil, err
				}

				fileName := file.GetName()
				if file.GoPkg.Path != "" {
					fileName = fmt.Sprintf("%s/%s", file.GoPkg.Path, fmt.Sprintf("%d_%s_down.sql", mTime, msgName))
				}

				files = append(files, &plugin.CodeGeneratorResponse_File{
					Name:    proto.String(fileName),
					Content: proto.String(code),
				})
				glog.V(1).Infof("Will emit %s", fileName)
			}
		}
	}
	return files, nil
}

func (g *generator) generateUp(msg *descriptor.Message) (string, error) {
	return applyTemplateUp(param{Message: msg})
}

func (g *generator) generateDown(msg *descriptor.Message) (string, error) {
	return applyTemplateDown(param{Message: msg})
}
