package controllers

import (
	"errors"
	"os"

	"github.com/danielronalds/clint/internal"
	"github.com/danielronalds/clint/internal/parsing"
	"github.com/danielronalds/clint/internal/pipelines"
)

func RunDefault() error {
	config, err := parsing.ParseClintFile(internal.DEFAULT_CONFIG_PATH)
	if err != nil {
		return err
	}

	if len(config.Pipelines) < 1 {
		return errors.New("no pilelines defined in config")
	}

	allStepsPass := pipelines.Run(&config.Pipelines[0])

	if !allStepsPass {
		os.Exit(1)
	}

	return nil
}
