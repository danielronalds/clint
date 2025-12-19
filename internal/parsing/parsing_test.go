package parsing

import (
	"reflect"
	"testing"

	"github.com/danielronalds/clint/internal/pipelines"
)

func TestParseConfig(t *testing.T) {
	config := []byte(`
name: "Test"
steps:
  - name: "Setup"
    cmd: "mkdir test"
  - name: "Test"
    cmd: "echo Hello World"
`)

	expected := &pipelines.Pipeline{
		Name: "Test",
		Steps: []pipelines.Step{
			{
				Name: "Setup",
				Cmd:  "mkdir test",
			},
			{
				Name: "Test",
				Cmd:  "echo Hello World",
			},
		},
	}

	result, err := parseConfig(config)

	if err != nil {
		t.Fatalf("parseConfig failed: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %+v, want %+v", result, expected)
	}
}
