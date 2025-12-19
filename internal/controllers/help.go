package controllers

import (
	"fmt"

	"github.com/danielronalds/clint/internal"
)

const HELP_MENU = `clint v%v

Usage: clint [command | pipeline]

A cli tool for running CI pipelines locally

Running clint with no arguments results in the first declared pipeline being run

Commands
  list      Lists available Pipelines
  help      Show this menu
`

func Help() error {
	fmt.Printf(HELP_MENU, internal.VERSION)

	return nil
}
