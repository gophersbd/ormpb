package generator

import (
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
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
	if err := proto.UnmarshalText(src, req); err != nil {
		t.Fatalf("proto.UnmarshalText(%s, &file) failed with %v; want success", src, err)
	}

	data, _ := proto.Marshal(req)
	s1 := strings.NewReader(string(data))

	_, err := ParseRequest(s1)

	if err != nil {
		t.Errorf("Failed to read Proto file with error %v", err)
	}
}
