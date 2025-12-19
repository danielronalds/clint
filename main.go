package main

import (
	"fmt"
	"log"

	"github.com/danielronalds/clint/internal/parsing"
)

func main() {
	path := "./clint.yaml"

	pipeline, err := parsing.ParseClintFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(pipeline.Name)
}
