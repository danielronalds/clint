package parsing

import (
	"fmt"
	"os"

	"github.com/danielronalds/clint/internal/pipelines"
	"gopkg.in/yaml.v3"
)

func ParseClintFile(path string) (*pipelines.Pipeline, error) {
	config, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err.Error())
	}

	return parseConfig(config)
}

func parseConfig(config []byte) (*pipelines.Pipeline, error) {
	var pipeline pipelines.Pipeline
	err := yaml.Unmarshal(config, &pipeline)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %v", err.Error())
	}

	return &pipeline, nil
}
