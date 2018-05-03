package generator

import (
	"testing"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
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

	err = reg.Load(&req)
	assert.Nil(t, err)

	file, err := reg.LookupFile("example.proto")
	assert.Nil(t, err)
	assert.NotNil(t, file)

	g := New(reg)
	assert.Equal(t, len(g), 2)
}
