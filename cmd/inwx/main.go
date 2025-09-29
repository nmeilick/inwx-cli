package main

import (
	"fmt"
	"os"
	_ "time/tzdata"

	"github.com/nmeilick/inwx-cli/internal/cli"
)

var (
	Version   = "dev"
	Commit    = "unknown"
	BuildDate = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Version = Version

	if err := app.Run(os.Args); err != nil {
		// Print the error to stderr to ensure the user sees the cause of the failure
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
