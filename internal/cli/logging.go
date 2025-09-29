package cli

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func setupLogging(c *cli.Context) error {
	config, err := LoadConfig(c)
	if err != nil {
		return err
	}

	// Set log level
	level, err := zerolog.ParseLevel(config.Logging.Level)
	if err != nil {
		level = zerolog.WarnLevel
	}
	zerolog.SetGlobalLevel(level)

	// Setup console writer for stderr
	var output io.Writer = os.Stderr
	if config.Logging.Colors && isatty.IsTerminal(os.Stderr.Fd()) {
		output = colorable.NewColorableStderr()
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        output,
		TimeFormat: "15:04:05",
		NoColor:    !config.Logging.Colors || !isatty.IsTerminal(os.Stderr.Fd()),
	})

	return nil
}
