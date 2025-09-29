package output

import (
	"os"

	"github.com/mattn/go-isatty"
)

func IsColorSupported() bool {
	return isatty.IsTerminal(os.Stdout.Fd())
}

func IsStderrColorSupported() bool {
	return isatty.IsTerminal(os.Stderr.Fd())
}

func ShouldUseColors(forceColors bool, noColors bool) bool {
	if noColors {
		return false
	}
	if forceColors {
		return true
	}
	return IsColorSupported()
}
