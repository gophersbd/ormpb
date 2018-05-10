package generator

import (
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	src := `
		file_to_generate: 'a.proto'
		proto_file <
			name: 'a.proto'
			message_type <
				name: 'Example'
			>
		>
	`
	req := new(plugin.CodeGeneratorRequest)
	err := proto.UnmarshalText(src, req)
	assert.Nil(t, err)

	data, _ := proto.Marshal(req)
	s1 := strings.NewReader(string(data))

	_, err = ParseRequest(s1)
	assert.Nil(t, err)
}
