package generator

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/gophersbd/ormpb/pkg/descriptor"
)

func TestLoadFile(t *testing.T) {
	reg := descriptor.NewRegistry()
	src := `
		file_to_generate: 'example.proto'
		proto_file <
			name: 'example.proto'
			message_type <
				name: 'Example'
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
	generatedFiles, err := g.Generate([]*descriptor.File{file})
	if err != nil {
		t.Fatalf("Generate File failed with %v; want success", err)
	}

	name := file.GetName()
	if file.GoPkg.Path != "" {
		name = fmt.Sprintf("%s/%s", file.GoPkg.Path, filepath.Base(name))
	}
	ext := filepath.Ext(name)
	base := strings.TrimSuffix(name, ext)
	output := fmt.Sprintf("%s.pb.orm.go", base)

	if got, want := generatedFiles[0].GetName(), output; want != got {
		t.Errorf("generated file name = %v; want %v", got, want)
	}
}
