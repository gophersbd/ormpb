package descriptor

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/protobuf"
)

// GoPackage represents a golang package
type GoPackage struct {
	// Path is the package path to the package.
	Path string
	// Name is the package name of the package
	Name string
	// Alias is an alias of the package unique within the current invokation of orm generator.
	Alias string
}

// Standard returns whether the import is a golang standard package.
func (p GoPackage) Standard() bool {
	return !strings.Contains(p.Path, ".")
}

// File wraps descriptor.FileDescriptorProto for richer features.
type File struct {
	*descriptor.FileDescriptorProto
	// GoPkg is the go package of the go file generated from this file..
	GoPkg GoPackage
	// Messages is the list of messages defined in this file.
	Messages []*Message
}

// Message describes a protocol buffer message types
type Message struct {
	// File is the file where the message is defined
	File *File
	// Outers is a list of outer messages if this message is a nested type.
	Outers []string
	*descriptor.DescriptorProto
	TableOptions *protobuf.TableOptions
	Fields       []*Field

	// Index is proto path index of this message in File.
	Index int
}

// Field wraps descriptor.FieldDescriptorProto for richer features.
type Field struct {
	// Message is the message type which this field belongs to.
	Message *Message
	*descriptor.FieldDescriptorProto
	Name   string
	Column *Column
}

type Column struct {
	Options   *protobuf.ColumnOptions
	Tags      map[string]interface{}
	Name      string
	Signature string
}

// FQMN return register name
func (m *Message) FQMN() string {
	components := []string{""}
	if m.File.Package != nil {
		components = append(components, m.File.GetPackage())
	}
	components = append(components, m.Outers...)
	components = append(components, m.GetName())
	return strings.Join(components, ".")
}
