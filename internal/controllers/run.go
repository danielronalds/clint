package controllers

import (
	"os"

	"github.com/danielronalds/clint/internal/parsing"
	"github.com/danielronalds/clint/internal/pipelines"
)

func Run() error {
	path := "./clint.yaml"

	pipeline, err := parsing.ParseClintFile(path)
	if err != nil {
		return err
	}

	allStepsPass := pipelines.Run(pipeline)

	if !allStepsPass {
		os.Exit(1)
	}

	return nil
}
