package pipelines

import (
	"fmt"
	"os/exec"
	"strings"
)

type Pipeline struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd"`
}

func Run(pipeline *Pipeline) bool {
	fmt.Printf("Running '%v' pipline\n", pipeline.Name)

	for _, step := range pipeline.Steps {
		succeeds := runStep(step)

		if !succeeds {
			fmt.Printf("- [x] %v\n", step.Name)
			return false
		}

		fmt.Printf("- [✔] %v\n", step.Name)
	}

	return true
}

func runStep(step Step) bool {
	splitCmd := strings.Split(step.Cmd, " ")

	program := splitCmd[0]

	args := splitCmd[1:]

	_, err := exec.Command(program, args...).Output()

	return err == nil
}
