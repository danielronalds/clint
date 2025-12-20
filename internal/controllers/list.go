package controllers

import (
	"fmt"

	"github.com/danielronalds/clint/internal/directories"
	"github.com/danielronalds/clint/internal/parsing"
)

func List() error {
	configPath, err := directories.FindConfigPath()
	if err != nil {
		return err
	}

	config, err := parsing.ParseClintFile(configPath)
	if err != nil {
		return err
	}

	fmt.Println("Available Pipelines")

	for _, pipeline := range config.Pipelines {
		fmt.Printf("- %v\n", pipeline.Name)
	}

	return nil
}
