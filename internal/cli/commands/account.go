package commands

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/nmeilick/inwx-cli/internal/cli/output"
)

func AccountCommand() *cli.Command {
	return &cli.Command{
		Name:  "account",
		Usage: "Account management",
		Subcommands: []*cli.Command{
			{
				Name:   "info",
				Usage:  "Show account information",
				Action: accountInfo,
			},
		},
	}
}

func accountInfo(c *cli.Context) error {
	client, err := createClient(c)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := client.Login(ctx); err != nil {
		return err
	}
	defer func() {
		if err := client.Logout(ctx); err != nil {
			log.Error().Err(err).Msg("Failed to logout")
		}
	}()

	account := client.Account()
	info, err := account.Info(ctx)
	if err != nil {
		return err
	}

	return formatOutput(c, func(formatter interface{}) string {
		switch f := formatter.(type) {
		case *output.TableFormatter:
			return f.FormatAccountInfo(info)
		case *output.JSONFormatter:
			return f.FormatAccountInfo(info)
		case *output.YAMLFormatter:
			return f.FormatAccountInfo(info)
		default:
			return "Unsupported format"
		}
	})
}
