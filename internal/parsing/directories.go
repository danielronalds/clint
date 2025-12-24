package parsing

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/danielronalds/clint/internal/pipelines"
	"gopkg.in/yaml.v3"
)

func ParsePipelinesInDir(directory string) ([]pipelines.Pipeline, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("unable to read directory: %w", err)
	}

	parsedPipelines := make([]pipelines.Pipeline, 0)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		pipeline, err := parsePipelineFile(directory, entry.Name())
		if err != nil {
			log.Printf("unable to parse pipeline file: %v\n", entry.Name())
			continue
		}

		parsedPipelines = append(parsedPipelines, *pipeline)
	}

	return parsedPipelines, nil
}

func parsePipelineFile(directory, filename string) (*pipelines.Pipeline, error) {
	path := filepath.Join(directory, filename)

	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	pipeline, err := parsePipeline(contents)
	if err != nil {
		return nil, err
	}

	pipeline.Name = strings.TrimSuffix(filename, ".yaml")

	return pipeline, nil
}

func parsePipeline(contents []byte) (*pipelines.Pipeline, error) {
	var pipeline pipelines.Pipeline
	err := yaml.Unmarshal(contents, &pipeline)
	if err != nil {
		return nil, err
	}

	return &pipeline, nil
}
