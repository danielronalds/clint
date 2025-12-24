package parsing

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/danielronalds/clint/internal/pipelines"
)

func TestParsePipelinesInDir(t *testing.T) {
	createTestDir()
	defer deleteTestDir()

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

	result, err := ParsePipelinesInDir(TEST_DIR)
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
const TEST_DIR = "test_dir"

func createTestDir() {
	err := os.Mkdir(TEST_DIR, 0755)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, name := range []string{"ci.yaml", "release.yaml"} {
		filePath := filepath.Join(TEST_DIR, name)

		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if _, err = file.WriteString(EXAMPLE_PIPELINE); err != nil {
			log.Fatalln(err.Error())
		}

		if err = file.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func deleteTestDir() {
	if err := os.RemoveAll(TEST_DIR); err != nil {
		log.Fatalf("unable to remove test directory, '%v': %v", TEST_DIR, err.Error())
	}
}
