package main

import (
	"log"
	"os"

	"github.com/danielronalds/clint/internal/controllers"
)

func main() {
	args := os.Args[1:]

	if err := run(args); err != nil {
		log.Fatalln(err.Error())
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return controllers.RunDefault()
	}

	switch args[0] {
	case "list":
		return controllers.List()
	default:
		return controllers.RunPipeline(args)
	}

}
