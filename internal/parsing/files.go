package parsing

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/danielronalds/clint/internal/pipelines"
	"gopkg.in/yaml.v3"
)

type ClintConfig struct {
	Pipelines    []pipelines.Pipeline `yaml:"pipelines"`
	PipelinesDir string               `yaml:"pipelines_dir"`
}

func ParseClintFile(directory, filename string) (*ClintConfig, error) {
	clintFilepath := filepath.Join(directory, filename)

	config, err := os.ReadFile(clintFilepath)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err.Error())
	}

	parsedConfig, err := parseConfig(config)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %v", err.Error())
	}

	if strings.TrimSpace(parsedConfig.PipelinesDir) != "" {
		pipelinesDir := filepath.Join(directory, parsedConfig.PipelinesDir)

		dirPipelines, err := ParsePipelinesInDir(pipelinesDir)
		if err != nil {
			return nil, err
		}

		parsedConfig.Pipelines = append(parsedConfig.Pipelines, dirPipelines...)
	}

	return parsedConfig, nil
}

func parseConfig(configContents []byte) (*ClintConfig, error) {
	var config ClintConfig
	err := yaml.Unmarshal(configContents, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
