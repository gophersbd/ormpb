package generator

import (
	"strings"
	"testing"
)

func TestParseRequest(t *testing.T) {
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
	s1 := strings.NewReader(src)
	_, err := ParseRequest(s1)
	if err != nil {
		t.Error("Failed to read Proto file")
	}
}
