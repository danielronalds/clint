package pipelines

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/danielronalds/clint/internal"
	a "github.com/logrusorgru/aurora/v4"
)

func printPassBadge(stepName string) {
	fmt.Printf("%v %v\n", a.Black(" PASS ").Bold().BgBrightGreen(), stepName)
}

func printFailBadge(stepName string) {
	fmt.Printf("%v %v\n", a.Black(" FAIL ").Bold().BgBrightRed(), stepName)
}

func printFailedOutput(output string) {
	linesInOutput := strings.Split(output, "\n")

	if len(linesInOutput) < internal.MAX_LINES_WITHOUT_PAGER {
		fmt.Printf("\n%s\n", output)
		return
	}

	pager, isSet := os.LookupEnv("PAGER")

	if !isSet {
		pager = internal.FALLBACK_PAGER
	}

	cmd := exec.Command(pager)
	if !isSet {
		cmd = exec.Command(pager, "-R")
	}
	cmd.Stdin = strings.NewReader(output)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("\n%s\n", output)
	}
}
