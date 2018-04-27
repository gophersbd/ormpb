package descriptor

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func loadFile(t *testing.T, reg *Registry, src string) *descriptor.FileDescriptorProto {
	var file descriptor.FileDescriptorProto
	if err := proto.UnmarshalText(src, &file); err != nil {
		t.Fatalf("proto.UnmarshalText(%s, &file) failed with %v; want success", src, err)
	}
	reg.loadFile(&file)
	return &file
}

func load(t *testing.T, reg *Registry, src string) error {
	var req plugin.CodeGeneratorRequest
	if err := proto.UnmarshalText(src, &req); err != nil {
		t.Fatalf("proto.UnmarshalText(%s, &file) failed with %v; want success", src, err)
	}
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
	if file == nil {
		t.Errorf("reg.files[%q] = nil; want non-nil", "example.proto")
		return
	}
	wantPkg := GoPackage{Path: ".", Name: "examples"}
	if got, want := file.GoPkg, wantPkg; got != want {
		t.Errorf("file.GoPkg = %#v; want %#v", got, want)
	}

	f, err := reg.LookupFile("example.proto")
	if err != nil {
		t.Errorf("reg.LookupFile(%q) failed with %v; want success", ".example.Example", err)
		return
	}
	if f.GetName() != "example.proto" {
		t.Errorf("file.Name = %v; want example.proto", f.GetName())
	}

	msg, err := reg.LookupMsg("", ".example.Example")
	if err != nil {
		t.Errorf("reg.LookupMsg(%q, %q)) failed with %v; want success", "", ".example.Example", err)
		return
	}
	if got, want := msg.DescriptorProto, fd.MessageType[0]; got != want {
		t.Errorf("reg.lookupMsg(%q, %q).DescriptorProto = %#v; want %#v", "", ".example.Example", got, want)
	}

	if got, want := msg.File, file; got != want {
		t.Errorf("msg.File = %v; want %v", got, want)
	}

	if got := msg.Outers; got != nil {
		t.Errorf("msg.Outers = %v; want %v", got, nil)
	}

	if got, want := len(msg.Fields), 1; got != want {
		t.Errorf("len(msg.Fields) = %d; want %d", got, want)
	} else if got, want := msg.Fields[0].FieldDescriptorProto, fd.MessageType[0].Field[0]; got != want {
		t.Errorf("msg.Fields[0].FieldDescriptorProto = %v; want %v", got, want)
	} else if got, want := msg.Fields[0].Message, msg; got != want {
		t.Errorf("msg.Fields[0].Message = %v; want %v", got, want)
	}

	if got, want := len(file.Messages), 1; got != want {
		t.Errorf("file.Meeesages = %#v; want %#v", file.Messages, []*Message{msg})
	}
	if got, want := file.Messages[0], msg; got != want {
		t.Errorf("file.Meeesages[0] = %v; want %v", got, want)
	}
}

func TestLoadFileNestedPackage(t *testing.T) {
	reg := NewRegistry()
	loadFile(t, reg, `
		name: 'example.proto'
		package: 'ormpb.example'
	`)

	file := reg.files["example.proto"]
	if file == nil {
		t.Errorf("reg.files[%q] = nil; want non-nil", "example.proto")
		return
	}
	wantPkg := GoPackage{Path: ".", Name: "ormpb_example"}
	if got, want := file.GoPkg, wantPkg; got != want {
		t.Errorf("file.GoPkg = %#v; want %#v", got, want)
	}
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
		if got, want := err == nil, spec.consistent; got != want {
			if want {
				t.Errorf("reg.Load(%s) failed with %v; want success", spec.req, err)
				continue
			}
			t.Errorf("reg.Load(%s) succeeded; want an package inconsistency error", spec.req)
		}
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
				[ormpb.protobuf.table] < name : "examples" >
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
	if err != nil {
		t.Errorf("reg.LookupMsg(%q, %q)) failed with %v; want success", "", ".example.Example", err)
		return
	}

	if msg.TableOptions == nil {
		t.Error("failed to parse TableOptions")
	}

	tableOptions := msg.TableOptions
	if got, want := tableOptions.GetName(), "examples"; got != want {
		t.Errorf("failed to parse TableOptions; want %v, got %v", want, got)
	}

	if len(msg.Fields) != 1 {
		t.Error("failed to parse Fileds")
	}

	field := msg.Fields[0]
	if field.ColumnOptions == nil {
		t.Error("failed to parse ColumnOptions")
	}

	columnOption := field.ColumnOptions
	if got, want := columnOption.GetName(), "labels"; got != want {
		t.Errorf("failed to parse ColumnOptions; want %v, got %v", want, got)
	}
	if got, want := columnOption.GetType(), "blob"; got != want {
		t.Errorf("failed to parse ColumnOptions; want %v, got %v", want, got)
	}
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
	if err != nil {
		t.Errorf("reg.LookupMsg(%q, %q)) failed with %v; want success", "", ".example.Example", err)
		return
	}
	_, err = reg.LookupMsg("", "example.Example")
	if err != nil {
		t.Errorf("reg.LookupMsg(%q, %q)) failed with %v; want success", "", ".example.Example", err)
		return
	}
	_, err = reg.LookupMsg(".example", "Example")
	if err != nil {
		t.Errorf("reg.LookupMsg(%q, %q)) failed with %v; want success", "", ".example.Example", err)
		return
	}
	_, err = reg.LookupMsg("example", "Example")
	if err != nil {
		t.Errorf("reg.LookupMsg(%q, %q)) failed with %v; want success", "", ".example.Example", err)
		return
	}
}
