package main

import (
	"flag"
	"os"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/gophersbd/ormpb/pkg/descriptor"
	_ "github.com/gophersbd/ormpb/pkg/dialect/mysql"
	_ "github.com/gophersbd/ormpb/pkg/dialect/postgres"
	"github.com/gophersbd/ormpb/pkg/generator"
	"github.com/spf13/pflag"
)

func main() {
	Start()
}

func init() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	// Convinces goflags that we have called Parse() to avoid noisy logs.
	_ = flag.CommandLine.Parse([]string{})
}

// Start starts running the ormpb generator
func Start() {
	req, err := generator.ParseRequest(os.Stdin)
	if err != nil {
		glog.Fatal(err)
	}

	reg := descriptor.NewRegistry()
	reg.CommandLineParameters(req.GetParameter())

	generators := generator.New(reg)

	if err = reg.Load(req); err != nil {
		writeError(err)
		return
	}

	var targets []*descriptor.File
	for _, target := range req.FileToGenerate {
		var f *descriptor.File
		if f, err = reg.LookupFile(target); err != nil {
			glog.Fatal(err)
		}
		targets = append(targets, f)
	}

	for _, g := range generators {
		out, err := g.Generate(targets)
		glog.V(1).Info("Processed code generator request")
		if err != nil {
			writeError(err)
			return
		}
		writeFiles(out)
	}

}

func writeFiles(out []*plugin.CodeGeneratorResponse_File) {
	writeResp(&plugin.CodeGeneratorResponse{File: out})
}

func writeError(err error) {
	writeResp(&plugin.CodeGeneratorResponse{Error: proto.String(err.Error())})
}

func writeResp(resp *plugin.CodeGeneratorResponse) {
	buf, err := proto.Marshal(resp)
	if err != nil {
		glog.Fatal(err)
	}
	if _, err := os.Stdout.Write(buf); err != nil {
		glog.Fatal(err)
	}
}
