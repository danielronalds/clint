package pipelines

import (
	"fmt"

	a "github.com/logrusorgru/aurora/v4"
)

func printPassBadge(stepName string) {
	fmt.Printf("%v %v\n", a.Black(" PASS ").Bold().BgBrightGreen(), stepName)
}

func printFailBadge(stepName string) {
	fmt.Printf("%v %v\n", a.Black(" FAIL ").Bold().BgBrightRed(), stepName)
}

func printFailedOutput(output string) {
	fmt.Printf("\n%v\n", output)
}
