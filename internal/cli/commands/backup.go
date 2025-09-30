package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/nmeilick/inwx-cli/internal/backup"
	"github.com/nmeilick/inwx-cli/internal/cli/output"
	"github.com/nmeilick/inwx-cli/pkg/inwx"
)

func BackupCommand() *cli.Command {
	return &cli.Command{
		Name:  "backup",
		Usage: "Backup management",
		Subcommands: []*cli.Command{
			{
				Name:   "list",
				Usage:  "List backup entries",
				Action: listBackups,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "domain",
						Aliases: []string{"d"},
						Usage:   "Filter by domain",
					},
					&cli.StringFlag{
						Name:    "operation",
						Aliases: []string{"op"},
						Usage:   "Filter by operation type (create, update, delete)",
					},
					&cli.StringFlag{
						Name:  "since",
						Usage: "Show entries since duration (e.g., 24h, 7d)",
					},
				},
			},
			{
				Name:      "revert",
				Usage:     "Revert backup entries",
				ArgsUsage: "<backup-id> [backup-id...]",
				Action:    revertBackup,
			},
			{
				Name:   "purge",
				Usage:  "Purge old backup entries",
				Action: purgeBackups,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "older-than",
						Usage:    "Remove entries older than duration (e.g., 30d, 6m)",
						Required: true,
					},
				},
			},
			{
				Name:   "verify",
				Usage:  "Verify backup integrity",
				Action: verifyBackups,
			},
		},
	}
}

func listBackups(c *cli.Context) error {
	store, err := backup.NewStore()
	if err != nil {
		return err
	}

	entries, err := store.List()
	if err != nil {
		return err
	}

	// Apply filters
	domain := c.String("domain")
	operation := c.String("operation")
	since := c.String("since")

	var filtered []*inwx.BackupEntry
	for _, entry := range entries {
		if domain != "" && entry.Record.Domain != domain {
			continue
		}
		if operation != "" && string(entry.Operation) != operation {
			continue
		}
		if since != "" {
			duration, err := time.ParseDuration(since)
			if err != nil {
				return fmt.Errorf("invalid duration format: %s", since)
			}
			if time.Since(entry.Timestamp) > duration {
				continue
			}
		}
		filtered = append(filtered, entry)
	}

	return formatOutput(c, func(formatter interface{}) string {
		switch f := formatter.(type) {
		case *output.TableFormatter:
			return f.FormatBackupEntries(filtered)
		case *output.JSONFormatter:
			return f.FormatBackupEntries(filtered)
		case *output.YAMLFormatter:
			return f.FormatBackupEntries(filtered)
		case *output.CSVFormatter:
			return f.FormatBackupEntries(filtered)
		default:
			return "Unsupported format"
		}
	})
}

func revertBackup(c *cli.Context) error {
	args := c.Args().Slice()
	if len(args) == 0 {
		return fmt.Errorf("at least one backup ID must be specified")
	}

	// Deduplicate IDs while preserving order
	seen := make(map[string]bool)
	var uniqueIDs []string
	for _, id := range args {
		if !seen[id] {
			seen[id] = true
			uniqueIDs = append(uniqueIDs, id)
		}
	}

	store, err := backup.NewStore()
	if err != nil {
		return err
	}

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

	// Create DNS service with backup store to track the revert operation
	backupStore, err := backup.NewStore()
	if err != nil {
		return fmt.Errorf("failed to initialize backup store: %w", err)
	}
	dns := client.DNS(inwx.WithBackupStore(backupStore))

	var errors []error
	successCount := 0

	for _, entryID := range uniqueIDs {
		entry, err := store.Get(entryID)
		if err != nil {
			errors = append(errors, fmt.Errorf("backup ID %s: %w", entryID, err))
			continue
		}

		switch entry.Operation {
		case inwx.OperationDelete:
			// Recreate the deleted record
			_, err = dns.CreateRecord(ctx, entry.Record)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to recreate record for backup %s: %w", entryID, err))
				continue
			}
			fmt.Printf("Successfully recreated deleted record (backup ID: %s, original record ID: %d)\n", entryID, entry.Record.ID)

		case inwx.OperationUpdate:
			// Revert to the previous state
			_, err = dns.UpdateRecord(ctx, entry.Record.ID, entry.Record)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to revert record for backup %s: %w", entryID, err))
				continue
			}
			fmt.Printf("Successfully reverted record (backup ID: %s, record ID: %d)\n", entryID, entry.Record.ID)

		case inwx.OperationCreate:
			// Delete the created record
			err = dns.DeleteRecord(ctx, entry.Record.ID)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to delete record for backup %s: %w", entryID, err))
				continue
			}
			fmt.Printf("Successfully deleted created record (backup ID: %s, record ID: %d)\n", entryID, entry.Record.ID)

		default:
			errors = append(errors, fmt.Errorf("backup ID %s: unknown operation type: %s", entryID, entry.Operation))
			continue
		}

		successCount++
	}

	if len(errors) > 0 {
		fmt.Printf("\nCompleted with %d successes and %d errors:\n", successCount, len(errors))
		for _, err := range errors {
			fmt.Printf("Error: %v\n", err)
		}
		if successCount == 0 {
			return fmt.Errorf("all revert operations failed")
		}
	} else {
		fmt.Printf("\nSuccessfully reverted %d backup entries\n", successCount)
	}

	return nil
}

func purgeBackups(c *cli.Context) error {
	olderThan := c.String("older-than")
	duration, err := time.ParseDuration(olderThan)
	if err != nil {
		return fmt.Errorf("invalid duration format: %s", olderThan)
	}

	store, err := backup.NewStore()
	if err != nil {
		return err
	}

	if atomicStore, ok := store.(*backup.AtomicStore); ok {
		err = atomicStore.PurgeOlderThan(duration)
	} else {
		return fmt.Errorf("purge operation requires atomic store")
	}

	if err != nil {
		return err
	}

	fmt.Printf("Successfully purged backup entries older than %s\n", olderThan)
	return nil
}

func verifyBackups(c *cli.Context) error {
	store, err := backup.NewStore()
	if err != nil {
		return err
	}

	if atomicStore, ok := store.(*backup.AtomicStore); ok {
		err = atomicStore.Verify()
	} else {
		return fmt.Errorf("verify operation requires atomic store")
	}

	if err != nil {
		return err
	}

	fmt.Println("All backup files are valid")
	return nil
}
