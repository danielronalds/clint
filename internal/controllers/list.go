package controllers

import (
	"fmt"

	"github.com/danielronalds/clint/internal/config"
	"github.com/danielronalds/clint/internal/parsing"
	"github.com/danielronalds/clint/internal/pipelines"
	a "github.com/logrusorgru/aurora/v4"
)

const COLUMN_GAP = 3

func List() error {
	configPath, err := config.FindConfigPath()
	if err != nil {
		return err
	}

	config, err := parsing.ParseClintFile(configPath, config.CONFIG_NAME)
	if err != nil {
		return err
	}

	columnWidth := longestWidthName(config.Pipelines, len("Name")) + COLUMN_GAP

	fmt.Printf("%-*s%s\n", columnWidth, a.Italic("Name").Bold(), a.Italic("Description").Bold())

	for _, pipeline := range config.Pipelines {
		fmt.Printf("%-*s%s\n", columnWidth, pipeline.Name, a.Italic(pipeline.Description))
	}

	return nil
}

func longestWidthName(pipelines []pipelines.Pipeline, startMax int) int {
	if len(pipelines) == 0 {
		return 0
	}

	max := startMax

	for _, pipeline := range pipelines {
		if len(pipeline.Name) > max {
			max = len(pipeline.Name)
		}
	}

	return max
}
