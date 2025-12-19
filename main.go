package main

import (
	"log"
	"os"

	"github.com/danielronalds/clint/internal/parsing"
	"github.com/danielronalds/clint/internal/pipelines"
)

func main() {
	path := "./clint.yaml"

	pipeline, err := parsing.ParseClintFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	allStepsPass := pipelines.Run(pipeline)

	if !allStepsPass {
		os.Exit(1)
	}
}
