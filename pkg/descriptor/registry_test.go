package descriptor

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/stretchr/testify/assert"
)

func loadFile(t *testing.T, reg *Registry, src string) *descriptor.FileDescriptorProto {
	var file descriptor.FileDescriptorProto
	err := proto.UnmarshalText(src, &file)
	assert.Nil(t, err)
	reg.loadFile(&file)
	return &file
}

func load(t *testing.T, reg *Registry, src string) error {
	var req plugin.CodeGeneratorRequest
	err := proto.UnmarshalText(src, &req)
	assert.Nil(t, err)
	return reg.Load(&req)
}

func TestLoadFile(t *testing.T) {
	reg := NewRegistry()
	fd := loadFile(t, reg, `
		name: 'example.proto'
		package: 'example'
		options <
			go_package : "ormpb.examples;examples"
		>
		message_type <
			name: 'Example'
			field <
				name: 'label'
				type: TYPE_STRING
			>
		>
	`)

	file := reg.files["example.proto"]
	assert.NotNil(t, file)
	assert.Equal(t, file.GoPkg, GoPackage{Path: ".", Name: "examples"})

	f, err := reg.LookupFile("example.proto")
	assert.Nil(t, err)

	assert.Equal(t, f.GetName(), "example.proto")

	msg, err := reg.LookupMsg("", ".example.Example")
	assert.Nil(t, err)
	assert.Equal(t, msg.DescriptorProto, fd.MessageType[0])
	assert.Equal(t, msg.File, file)
	assert.Nil(t, msg.Outers)
	assert.Equal(t, len(msg.Fields), 1)
	assert.Equal(t, msg.Fields[0].FieldDescriptorProto, fd.MessageType[0].Field[0])
	assert.Equal(t, msg.Fields[0].Message, msg)
	assert.Equal(t, len(file.Messages), 1)
	assert.Equal(t, file.Messages[0], msg)
}

func TestLoadFileNestedPackage(t *testing.T) {
	reg := NewRegistry()
	loadFile(t, reg, `
		name: 'example.proto'
		package: 'ormpb.example'
	`)

	file := reg.files["example.proto"]
	assert.NotNil(t, file)
	assert.Equal(t, file.GoPkg, GoPackage{Path: ".", Name: "ormpb_example"})
}

func TestLoadWithInconsistentTargetPackage(t *testing.T) {
	for _, spec := range []struct {
		req        string
		consistent bool
	}{
		// root package, no explicit go package
		{
			req: `
				file_to_generate: 'a.proto'
				file_to_generate: 'b.proto'
				proto_file <
					name: 'a.proto'
					message_type <
						name: 'Example'
						field <
							name: 'label'
							type: TYPE_STRING
						>
					>
				>
				proto_file <
					name: 'b.proto'
					message_type <
						name: 'Example'
						field <
							name: 'label'
							type: TYPE_STRING
						>
					>
				>
			`,
			consistent: false,
		},
	} {
		reg := NewRegistry()
		err := load(t, reg, spec.req)
		assert.NotNil(t, err)
	}
}

func TestExtension(t *testing.T) {
	reg := NewRegistry()
	loadFile(t, reg, `
		name: 'example.proto'
		package: 'example'
		message_type <
			name: 'Example'
			options <
				[ormpb.protobuf.table] < name : "examples", type: "postgres" >
			>
			field <
				name: 'label'
				type: TYPE_STRING
				options <
					[ormpb.protobuf.column] < name : "labels" type: "blob">
				>
			>
		>
	`)

	msg, err := reg.LookupMsg("", ".example.Example")
	assert.Nil(t, err)
	assert.NotNil(t, msg.TableOptions)
	assert.Equal(t, msg.TableOptions.GetName(), "examples")
	assert.Equal(t, len(msg.Fields), 1)

	column := msg.Fields[0].Column
	assert.NotNil(t, column.Options)
	assert.Equal(t, column.Options.GetName(), "labels")
	assert.Equal(t, column.Options.GetType(), "blob")
}

func TestPathPrefix(t *testing.T) {
	reg := NewRegistry()
	loadFile(t, reg, `
		name: 'example.proto'
		package: 'example'
		message_type <
			name: 'Example'
		>
	`)

	_, err := reg.LookupMsg("", ".example.Example")
	assert.Nil(t, err)
	_, err = reg.LookupMsg("", "example.Example")
	assert.Nil(t, err)
	_, err = reg.LookupMsg(".example", "Example")
	assert.Nil(t, err)
	_, err = reg.LookupMsg("example", "Example")
	assert.Nil(t, err)
}

func TestCommandLineParameters(t *testing.T) {
	reg := NewRegistry()
	reg.CommandLineParameters("migrations=test")
	loadFile(t, reg, `
		name: 'example.proto'
		package: 'example'
		message_type <
			name: 'Example'
		>
	`)
	file, err := reg.LookupFile("example.proto")
	assert.Nil(t, err)
	assert.Equal(t, file.MigrationDir, "test")
}
