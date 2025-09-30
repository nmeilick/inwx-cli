package cli

import (
	"github.com/urfave/cli/v2"

	"github.com/nmeilick/inwx-cli/internal/cli/commands"
)

func NewApp() *cli.App {
	app := &cli.App{
		Name:  "inwx",
		Usage: "INWX DomRobot API CLI tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
				EnvVars: []string{"INWX_CONFIG"},
			},
			&cli.StringFlag{
				Name:    "endpoint",
				Usage:   "API endpoint URL (overrides test/production environment)",
				EnvVars: []string{"INWX_ENDPOINT"},
			},
			&cli.StringFlag{
				Name:    "username",
				Aliases: []string{"u"},
				Usage:   "INWX username",
				EnvVars: []string{"INWX_USERNAME"},
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "INWX password",
				EnvVars: []string{"INWX_PASSWORD"},
			},
			&cli.BoolFlag{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Use test environment",
				EnvVars: []string{"INWX_TEST"},
			},
			&cli.IntFlag{
				Name:    "timeout",
				Usage:   "API request timeout in seconds",
				Value:   30,
				EnvVars: []string{"INWX_TIMEOUT"},
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output format (table, json, yaml, csv)",
				Value:   "table",
				EnvVars: []string{"INWX_OUTPUT"},
			},
			&cli.BoolFlag{
				Name:    "no-colors",
				Usage:   "Disable colored output",
				EnvVars: []string{"INWX_NO_COLORS"},
			},
			&cli.StringFlag{
				Name:    "log-level",
				Usage:   "Log level (trace, debug, info, warn, error, fatal)",
				Value:   "warn",
				EnvVars: []string{"INWX_LOG_LEVEL"},
			},
			&cli.BoolFlag{
				Name:    "yes",
				Aliases: []string{"y"},
				Usage:   "Skip all interactive confirmation prompts",
				EnvVars: []string{"INWX_YES"},
			},
		},
		Commands: []*cli.Command{
			commands.DNSCommand(),
			commands.DomainCommand(),
			commands.AccountCommand(),
			commands.BackupCommand(),
		},
		Before: func(c *cli.Context) error {
			return setupLogging(c)
		},
	}

	return app
}
