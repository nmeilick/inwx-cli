package commands

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/nmeilick/inwx-cli/internal/cli/output"
)

func DomainCommand() *cli.Command {
	return &cli.Command{
		Name:  "domain",
		Usage: "Domain management",
		Subcommands: []*cli.Command{
			{
				Name:   "list",
				Usage:  "List domains",
				Action: listDomains,
			},
		},
	}
}

func listDomains(c *cli.Context) error {
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

	domain := client.Domain()
	domains, err := domain.List(ctx)
	if err != nil {
		return err
	}

	return formatOutput(c, func(formatter interface{}) string {
		switch f := formatter.(type) {
		case *output.TableFormatter:
			return f.FormatDomains(domains)
		case *output.JSONFormatter:
			return f.FormatDomains(domains)
		case *output.YAMLFormatter:
			return f.FormatDomains(domains)
		default:
			return "Unsupported format"
		}
	})
}
