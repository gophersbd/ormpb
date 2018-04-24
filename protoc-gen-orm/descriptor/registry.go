package descriptor

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	goGen "github.com/golang/protobuf/protoc-gen-go/generator"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/gophersbd/ormpb/protobuf"
)

// Registry is a registry of information extracted from plugin.CodeGeneratorRequest.
type Registry struct {
	// msgs is a mapping from fully-qualified message name to descriptor
	msgs map[string]*Message

	// files is a mapping from file path to descriptor
	files map[string]*File

	// prefix is a prefix to be inserted to golang package paths generated from proto package names.
	prefix string

	// pkgMap is a user-specified mapping from file path to proto package.
	pkgMap map[string]string

	// importPath is used as the package if no input files declare go_package. If it contains slashes, everything up to the rightmost slash is ignored.
	importPath string
}

// NewRegistry returns a new Registry.
func NewRegistry() *Registry {
	return &Registry{
		msgs:  make(map[string]*Message),
		files: make(map[string]*File),
	}
}

// Load loads definitions of services, methods, messages, enumerations and fields from "req".
func (r *Registry) Load(req *plugin.CodeGeneratorRequest) error {
	for _, file := range req.GetProtoFile() {
		r.loadFile(file)
	}

	var targetPkg string
	for _, name := range req.FileToGenerate {
		target := r.files[name]
		if target == nil {
			return fmt.Errorf("no such file: %s", name)
		}
		name := r.packageIdentityName(target.FileDescriptorProto)
		if targetPkg == "" {
			targetPkg = name
		} else {
			if targetPkg != name {
				return fmt.Errorf("inconsistent package names: %s %s", targetPkg, name)
			}
		}
	}
	return nil
}

// loadFile loads messages, enumerations and fields from "file".
// It does not loads services and methods in "file".  You need to call
// loadServices after loadFiles is called for all files to load services and methods.
func (r *Registry) loadFile(file *descriptor.FileDescriptorProto) {
	pkg := GoPackage{
		Path: r.goPackagePath(file),
		Name: r.defaultGoPackageName(file),
	}

	f := &File{
		FileDescriptorProto: file,
		GoPkg:               pkg,
	}

	r.files[file.GetName()] = f
	r.registerMsg(f, nil, file.GetMessageType())
}

func (r *Registry) registerMsg(file *File, outerPath []string, msgs []*descriptor.DescriptorProto) {
	for _, md := range msgs {
		m := &Message{
			DescriptorProto: md,
		}

		if md.Options != nil {
			if proto.HasExtension(md.Options, ormpb_protobuf.E_Table) {
				to, _ := proto.GetExtension(md.Options, ormpb_protobuf.E_Table)
				m.TableOption = to.(*ormpb_protobuf.TableOptions)
			}
		}

		for _, fd := range md.GetField() {

			f := &Field{
				FieldDescriptorProto: fd,
			}

			if fd.Options != nil {
				if proto.HasExtension(fd.Options, ormpb_protobuf.E_Column) {
					to, _ := proto.GetExtension(fd.Options, ormpb_protobuf.E_Column)
					f.ColumnOption = to.(*ormpb_protobuf.ColumnOptions)
				}
			}

			f.Name = goGen.CamelCase(fd.GetName())

			m.Fields = append(m.Fields, f)
		}
		file.Messages = append(file.Messages, m)

		var outers []string
		outers = append(outers, outerPath...)
		outers = append(outers, m.GetName())
		r.registerMsg(file, outers, m.GetNestedType())
	}
}

// goPackagePath returns the go package path which go files generated from "f" should have.
// It respects the mapping registered by AddPkgMap if exists. Or use go_package as import path
// if it includes a slash,  Otherwide, it generates a path from the file name of "f".
func (r *Registry) goPackagePath(f *descriptor.FileDescriptorProto) string {
	name := f.GetName()
	if pkg, ok := r.pkgMap[name]; ok {
		return path.Join(r.prefix, pkg)
	}

	gopkg := f.Options.GetGoPackage()
	idx := strings.LastIndex(gopkg, "/")
	if idx >= 0 {
		if sc := strings.LastIndex(gopkg, ";"); sc > 0 {
			gopkg = gopkg[:sc+1-1]
		}
		return gopkg
	}

	return path.Join(r.prefix, path.Dir(name))
}

// defaultGoPackageName returns the default go package name to be used for go files generated from "f".
// You might need to use an unique alias for the package when you import it.  Use ReserveGoPackageAlias to get a unique alias.
func (r *Registry) defaultGoPackageName(f *descriptor.FileDescriptorProto) string {
	name := r.packageIdentityName(f)
	return sanitizePackageName(name)
}

// sanitizePackageName replaces unallowed character in package name
// with allowed character.
func sanitizePackageName(pkgName string) string {
	pkgName = strings.Replace(pkgName, ".", "_", -1)
	pkgName = strings.Replace(pkgName, "-", "_", -1)
	return pkgName
}

// packageIdentityName returns the identity of packages.
// protoc-gen-grpc-gateway rejects CodeGenerationRequests which contains more than one packages
// as protoc-gen-go does.
func (r *Registry) packageIdentityName(f *descriptor.FileDescriptorProto) string {
	if f.Options != nil && f.Options.GoPackage != nil {
		gopkg := f.Options.GetGoPackage()
		idx := strings.LastIndex(gopkg, "/")
		if idx < 0 {
			gopkg = gopkg[idx+1:]
		}

		gopkg = gopkg[idx+1:]
		// package name is overrided with the string after the
		// ';' character
		sc := strings.IndexByte(gopkg, ';')
		if sc < 0 {
			return sanitizePackageName(gopkg)

		}
		return sanitizePackageName(gopkg[sc+1:])
	}
	if p := r.importPath; len(p) != 0 {
		if i := strings.LastIndex(p, "/"); i >= 0 {
			p = p[i+1:]
		}
		return p
	}

	if f.Package == nil {
		base := filepath.Base(f.GetName())
		ext := filepath.Ext(base)
		return strings.TrimSuffix(base, ext)
	}
	return f.GetPackage()
}

// LookupFile looks up a file by name.
func (r *Registry) LookupFile(name string) (*File, error) {
	f, ok := r.files[name]
	if !ok {
		return nil, fmt.Errorf("no such file given: %s", name)
	}
	return f, nil
}