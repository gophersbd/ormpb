package orm

import (
	"fmt"
	"path/filepath"
	"strings"
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

	err = reg.Load(&req)
	assert.Nil(t, err)

	file, err := reg.LookupFile("example.proto")
	assert.Nil(t, err)
	assert.NotNil(t, file)

	g := NewGenerator(reg)
	generatedFiles, err := g.Generate([]*descriptor.File{file})
	assert.Nil(t, err)

	name := file.GetName()
	if file.GoPkg.Path != "" {
		name = fmt.Sprintf("%s/%s", file.GoPkg.Path, filepath.Base(name))
	}
	ext := filepath.Ext(name)
	base := strings.TrimSuffix(name, ext)
	output := fmt.Sprintf("%s.pb.orm.go", base)

	assert.Equal(t, generatedFiles[0].GetName(), output)
}
