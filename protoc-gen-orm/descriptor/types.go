package descriptor

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/gophersbd/ormpb/protobuf"
)

// GoPackage represents a golang package
type GoPackage struct {
	// Path is the package path to the package.
	Path string
	// Name is the package name of the package
	Name string
	// Alias is an alias of the package unique within the current invokation of grpc-gateway generator.
	Alias string
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
	*descriptor.DescriptorProto
	TableOption *protobuf.TableOptions
	Fields      []*Field
}

// Field wraps descriptor.FieldDescriptorProto for richer features.
type Field struct {
	*descriptor.FieldDescriptorProto
	Name         string
	ColumnOption *protobuf.ColumnOptions
}
