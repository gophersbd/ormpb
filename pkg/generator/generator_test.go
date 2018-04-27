package generator

import (
	"testing"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/gophersbd/ormpb/pkg/descriptor"
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
	if err := proto.UnmarshalText(src, &req); err != nil {
		t.Fatalf("proto.UnmarshalText(%s, &file) failed with %v; want success", src, err)
	}
	if err := reg.Load(&req); err != nil {
		t.Fatalf("Load CodeGeneratorRequest failed with %v; want success", err)
	}

	file, err := reg.LookupFile("example.proto")
	if err != nil {
		t.Fatalf("Load File failed with %v; want success", err)
	}
	if file == nil {
		t.Errorf("reg.files[%q] = nil; want non-nil", "example.proto")
		return
	}

	g := New(reg)

	if len(g) != 2 {
		t.Errorf("want %v generator, get %v", 2, len(g))
	}
}
