package pipelines

import (
	"fmt"
	"os/exec"
	"strings"

	a "github.com/logrusorgru/aurora/v4"
)

type Pipeline struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd"`
}

// TODO: There needs to be an option to see the failing output...
func Run(pipeline *Pipeline) bool {
	fmt.Printf("Running '%v' pipline\n\n", a.Bold(pipeline.Name))

	for _, step := range pipeline.Steps {
		output, succeeds := runStep(step)

		if !succeeds {
			fmt.Printf("%v %v\n\n%v", a.Black(" FAIL ").Bold().BgBrightRed(), step.Name,output)
			return false
		}

		fmt.Printf("%v %v\n", a.Black(" PASS ").Bold().BgBrightGreen(), step.Name)
	}

	return true
}

func runStep(step Step) (string, bool) {
	splitCmd := strings.Split(step.Cmd, " ")

	program := splitCmd[0]

	args := splitCmd[1:]

	output, err := exec.Command(program, args...).CombinedOutput()

	return string(output), err == nil
}
