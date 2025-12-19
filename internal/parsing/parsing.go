package parsing

import (
	"fmt"
	"os"

	"github.com/danielronalds/clint/internal/pipelines"
	"gopkg.in/yaml.v3"
)

type ClintConfig struct {
	Pipelines []pipelines.Pipeline `yaml:"pipelines"`
}

func ParseClintFile(path string) (*ClintConfig, error) {
	config, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err.Error())
	}

	return parseConfig(config)
}

func parseConfig(configContents []byte) (*ClintConfig, error) {
	var config ClintConfig
	err := yaml.Unmarshal(configContents, &config)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %v", err.Error())
	}

	return &config, nil
}
