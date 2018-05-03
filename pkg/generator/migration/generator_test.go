package migration

import (
	"testing"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {
	reg := descriptor.NewRegistry()
	src := `
		file_to_generate: 'example.proto'
		proto_file <
			name: 'example.proto'
			message_type <
				name: 'Example'
				options <
					[ormpb.protobuf.table] < name : "examples", type: "postgres" >
				>
				field <
					name: 'label'
					type: TYPE_STRING
				>
			>
		>
	`

	var req plugin.CodeGeneratorRequest
	err := proto.UnmarshalText(src, &req)
	assert.Nil(t, err)

	reg.CommandLineParameters("migrations=.")
	err = reg.Load(&req)
	assert.Nil(t, err)

	file, err := reg.LookupFile("example.proto")
	assert.Nil(t, err)
	assert.NotNil(t, file)

	g := NewGenerator(reg)
	_, err = g.Generate([]*descriptor.File{file})
	assert.Nil(t, err)
}
