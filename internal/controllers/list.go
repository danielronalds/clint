package controllers

import (
	"fmt"

	"github.com/danielronalds/clint/internal"
	"github.com/danielronalds/clint/internal/parsing"
)

func List() error {
	config, err := parsing.ParseClintFile(internal.DEFAULT_CONFIG_PATH)
	if err != nil {
		return err
	}

	fmt.Println("Available Pipelines")

	for _, pipeline := range config.Pipelines {
		fmt.Printf("- %v\n", pipeline.Name)
	}

	return nil
}
