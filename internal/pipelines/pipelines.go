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
	Name   string `yaml:"name"`
	Cmd    string `yaml:"cmd"`
	OnFail string `yaml:"on_fail"`
}

func (s Step) hasFailCmd() bool {
	return strings.TrimSpace(s.OnFail) != ""
}

func Run(pipeline *Pipeline) bool {
	fmt.Printf("Running '%v' pipline\n\n", a.Bold(pipeline.Name))

	for _, step := range pipeline.Steps {
		output, succeeds := runStep(step)

		if !succeeds {
			fmt.Printf("%v %v\n", a.Black(" FAIL ").Bold().BgBrightRed(), step.Name)
			fmt.Printf("\n%v", output)
			return false
		}

		fmt.Printf("%v %v\n", a.Black(" PASS ").Bold().BgBrightGreen(), step.Name)
	}

	return true
}

func runStep(step Step) (string, bool) {
	if step.Cmd == "" {
		return "Empty Command", false
	}

	output, err := exec.Command("bash", "-c", step.Cmd).CombinedOutput()

	if err != nil && step.hasFailCmd() {
		output, _ := exec.Command("bash", "-c", step.OnFail).CombinedOutput()
		return string(output), false
	}

	return string(output), err == nil
}
