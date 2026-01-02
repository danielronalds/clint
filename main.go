package main

import (
	"os"

	"github.com/danielronalds/clint/internal/controllers"
	"github.com/danielronalds/clint/internal/logger"
)

func main() {
	args := os.Args[1:]

	if err := run(args); err != nil {
		logger.Error(err.Error() + "\n")
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return controllers.RunDefault()
	}

	switch args[0] {
	case "list":
		return controllers.List()
	case "help":
		return controllers.Help()
	case "run":
		return controllers.RunPipeline(args[1:])
	default:
		return controllers.RunPipeline(args)
	}

}
