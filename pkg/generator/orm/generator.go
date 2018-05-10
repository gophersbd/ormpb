package orm

import (
	"fmt"
	"go/format"
	"path"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/gophersbd/ormpb/pkg/validation"
)

// Generator holds Registry
type Generator struct {
	reg         *descriptor.Registry
	baseImports []descriptor.GoPackage
}

// NewGenerator return Generator interface
func NewGenerator(reg *descriptor.Registry) *Generator {
	var imports []descriptor.GoPackage
	for _, pkgpath := range []string{
		"github.com/gophersbd/ormpb/pkg/runtime",
	} {
		pkg := descriptor.GoPackage{
			Path: pkgpath,
			Name: path.Base(pkgpath),
		}

		imports = append(imports, pkg)
	}
	return &Generator{reg: reg, baseImports: imports}
}

// Generate receives target files and returns generated code files
func (g *Generator) Generate(targets []*descriptor.File) ([]*plugin.CodeGeneratorResponse_File, error) {
	var files []*plugin.CodeGeneratorResponse_File
	for _, file := range targets {
		for _, m := range file.Messages {
			if err := validation.ValidateTableOptions(m); err != nil {
				return nil, err
			}

			if err := validation.ValidateColumnOptions(m.Fields); err != nil {
				return nil, err
			}
		}

		glog.V(1).Infof("Processing %s", file.GetName())
		code, err := g.generate(file)
		if err != nil {
			return nil, err
		}
		formatted, err := format.Source([]byte(code))
		if err != nil {
			glog.Errorf("%v: %s", err, code)
			return nil, err
		}
		name := file.GetName()
		if file.GoPkg.Path != "" {
			name = fmt.Sprintf("%s/%s", file.GoPkg.Path, filepath.Base(name))
		}
		ext := filepath.Ext(name)
		base := strings.TrimSuffix(name, ext)
		output := fmt.Sprintf("%s.pb.orm.go", base)
		files = append(files, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(output),
			Content: proto.String(string(formatted)),
		})
		glog.V(1).Infof("Will emit %s", output)
	}
	return files, nil
}

func (g *Generator) generate(file *descriptor.File) (string, error) {
	pkgSeen := make(map[string]bool)
	var imports []descriptor.GoPackage
	for _, pkg := range g.baseImports {
		pkgSeen[pkg.Path] = true
		imports = append(imports, pkg)
	}
	return applyTemplate(param{File: file, Imports: imports})
}
