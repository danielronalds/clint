package parsing

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Pipeline struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd"`
}

func ParseClintFile(path string) (*Pipeline, error) {
	config, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err.Error())
	}

	return parseConfig(config)
}

func parseConfig(config []byte) (*Pipeline, error) {
	var pipeline Pipeline
	err := yaml.Unmarshal(config, &pipeline)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %v", err.Error())
	}

	return &pipeline, nil
}
