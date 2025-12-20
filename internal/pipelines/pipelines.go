package pipelines

import (
	"fmt"
	"os"
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
			printFailBadge(step.Name)
			printFailedOutput(output)
			return false
		}

		printPassBadge(step.Name)
	}

	return true
}

func runStep(step Step) (string, bool) {
	if step.Cmd == "" {
		return "Empty Command", false
	}

	cmd := exec.Command("bash", "-c", step.Cmd)
	cmd.Env = append(os.Environ(), "CLICOLOR_FORCE=1")
	output, err := cmd.CombinedOutput()

	if err != nil && step.hasFailCmd() {
		cmd = exec.Command("bash", "-c", step.OnFail)
		cmd.Env = append(os.Environ(), "CLICOLOR_FORCE=1")
		output, _ := cmd.CombinedOutput()
		return string(output), false
	}

	return string(output), err == nil
}
