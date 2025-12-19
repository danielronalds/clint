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

	return nil
}
