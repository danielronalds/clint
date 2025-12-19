package controllers

import (
	"errors"
	"os"

	"github.com/danielronalds/clint/internal"
	"github.com/danielronalds/clint/internal/parsing"
	"github.com/danielronalds/clint/internal/pipelines"
)

func RunPipeline(args []string) error {
	if len(args) != 1 {
		return errors.New("too many arguments")
	}

	pipelineName := args[0]

	config, err := parsing.ParseClintFile(internal.DEFAULT_CONFIG_PATH)
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
