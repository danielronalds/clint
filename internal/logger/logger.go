package logger

import (
	"fmt"
	"time"

	a "github.com/logrusorgru/aurora/v4"
)

func Warning(msg string, args ...any) {
	warningBadge := a.Black(" WARN ").Bold().BgBrightYellow().String()

	time := a.Bold(fmt.Sprintf("[%s]", time.Now().Format(time.TimeOnly))).String()

	fmt.Printf(warningBadge+" "+time+" "+msg, args...)
}

func Error(msg string, args ...any) {
	errorBadge := a.Black(" ERROR ").Bold().BgBrightRed().String()

	time := a.Bold(fmt.Sprintf("[%s]", time.Now().Format(time.TimeOnly))).String()

	fmt.Printf(errorBadge+" "+time+" "+msg, args...)
}
