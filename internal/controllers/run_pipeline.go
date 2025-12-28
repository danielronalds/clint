package controllers

import (
	"errors"
	"os"

	"github.com/danielronalds/clint/internal/config"
	"github.com/danielronalds/clint/internal/parsing"
	"github.com/danielronalds/clint/internal/pipelines"
)

func RunPipeline(args []string) error {
	switch len(args) {
	case 0:
		return errors.New("no pipeline provided")
	case 1:
		break
	default:
		return errors.New("too many arguments provided")
	}

	pipelineName := args[0]

	configPath, err := config.FindConfigPath()
	if err != nil {
		return err
	}

	config, err := parsing.ParseClintFile(configPath, config.CONFIG_NAME)
	if err != nil {
		return err
	}

	pipeline := findPipeline(pipelineName, *config)
	if pipeline == nil {
		return errors.New("unable to find pipeline with the given name")
	}

	allStepsPass := pipelines.Run(pipeline)

	if !allStepsPass {
		os.Exit(1)
	}

	return nil
}

func findPipeline(name string, config parsing.ClintConfig) *pipelines.Pipeline {
	for _, pipeline := range config.Pipelines {
		if pipeline.Name == name {
			return &pipeline
		}
	}

	return nil
}
