package parsing

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/danielronalds/clint/internal/pipelines"
)

type path = string

func TestParsePipelinesInDir(t *testing.T) {
	testDir := createsTestDir()
	defer os.RemoveAll(testDir)

	expected := []pipelines.Pipeline{
		{
			Name: "ci",
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
			Name: "release",
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
	}

	result, err := ParsePipelinesInDir(testDir)
	if err != nil {
		log.Fatalf("ParsePipelinesInDir failed: %v", err.Error())
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %+v, want %+v", result, expected)
	}
}

const EXAMPLE_PIPELINE = `
steps:
  - name: "Setup"
    cmd: "mkdir test"
  - name: "Test"
    cmd: "echo Hello World"
`

func createsTestDir() path {
	testDir := "test_dir"

	err := os.Mkdir(testDir, 0755)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, name := range []string{"ci.yaml", "release.yaml"} {
		filePath := filepath.Join(testDir, name)

		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalln(err.Error())
		}

		file.WriteString(EXAMPLE_PIPELINE)

		file.Close()
	}

	return testDir
}
