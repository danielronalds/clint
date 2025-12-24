package parsing

import (
	"reflect"
	"testing"

	"github.com/danielronalds/clint/internal/pipelines"
)

func TestParseConfig(t *testing.T) {
	config := []byte(`
pipelines:
- name: "Test"
  steps:
    - name: "Setup"
      cmd: "mkdir test"
    - name: "Test"
      cmd: "echo Hello World"
- name: "Pipeline 2"
  steps:
    - name: "Setup"
      cmd: "mkdir test"
      on_fail: "echo 'unable to make dir'"
    - name: "Test"
      cmd: "echo Hello World"
`)

	expected := &ClintConfig{
		Pipelines: []pipelines.Pipeline{
			{
				Name: "Test",
				Steps: []pipelines.Step{
					{
						Name:   "Setup",
						Cmd:    "mkdir test",
						OnFail: "",
					},
					{
						Name:   "Test",
						Cmd:    "echo Hello World",
						OnFail: "",
					},
				},
			},
			{
				Name: "Pipeline 2",
				Steps: []pipelines.Step{
					{
						Name:   "Setup",
						Cmd:    "mkdir test",
						OnFail: "echo 'unable to make dir'",
					},
					{
						Name:   "Test",
						Cmd:    "echo Hello World",
						OnFail: "",
					},
				},
			},
		},
	}

	result, err := parseConfig(config)

	if err != nil {
		t.Fatalf("parseConfig failed: %v", err.Error())
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %+v, want %+v", result, expected)
	}
}
