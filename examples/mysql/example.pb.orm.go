// Code generated by protoc-gen-orm. DO NOT EDIT.
// source: ormpb/examples/mysql/example.proto

package importpb

import (
	runtime "github.com/gophersbd/ormpb/pkg/runtime"
)

func (*Example) TableName() string {
	return "examples"
}

var (
	_ExampleTagMap = map[string]map[string]string{
		"UserId": {
			runtime.ColumnTagAutoIncrement: "true",
			runtime.ColumnTagPrimaryKey:    "true",
		},
		"Name": {
			runtime.ColumnTagName: "name",
			runtime.ColumnTagSize: "128",
		},
		"Email": {
			runtime.ColumnTagNotNull: "true",
			runtime.ColumnTagUnique:  "true",
		},
		"Point": {
			runtime.ColumnTagDefault: "17.33",
		},
	}
)

func (*Example) Tag(field, tag string) (val string, found bool) {
	val, found = _ExampleTagMap[field][tag]
	return
}

func (*ExampleAutoIncrement) TableName() string {
	return "example_auto_increment"
}

var (
	_ExampleAutoIncrementTagMap = map[string]map[string]string{
		"UserId": {
			runtime.ColumnTagAutoIncrement: "true",
		},
	}
)

func (*ExampleAutoIncrement) Tag(field, tag string) (val string, found bool) {
	val, found = _ExampleAutoIncrementTagMap[field][tag]
	return
}
