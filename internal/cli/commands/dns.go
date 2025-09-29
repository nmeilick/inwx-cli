package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/nmeilick/inwx-cli/internal/backup"
	"github.com/nmeilick/inwx-cli/internal/cli/output"
	"github.com/nmeilick/inwx-cli/internal/utils"
	"github.com/nmeilick/inwx-cli/pkg/inwx"
)

// createDNSService creates a DNS service with backup store initialized
func createDNSService(c *cli.Context, client *inwx.Client) (*inwx.DNSService, error) {
	// Initialize backup store
	backupStore, err := backup.NewStore()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize backup store: %w", err)
	}

	// Create DNS service with backup store
	dns := client.DNS(inwx.WithBackupStore(backupStore))
	return dns, nil
}

// parseCommaSeparatedValues takes a slice of strings, replaces commas with spaces,
// then splits on any whitespace, returning a flat slice of individual values.
func parseCommaSeparatedValues(values []string) []string {
	if len(values) == 0 {
		return nil
	}

	var result []string
	for _, value := range values {
		// Replace commas with spaces
		normalized := strings.ReplaceAll(value, ",", " ")
		// Split on any whitespace
		parts := strings.Fields(normalized)
		result = append(result, parts...)
	}
	return result
}

func validateDNSRecordInput(domain, recordType, name, content string, ttl int) error {
	if err := utils.ValidateDomain(domain); err != nil {
		return fmt.Errorf("invalid domain: %w", err)
	}

	if err := utils.ValidateRecordType(recordType); err != nil {
		return fmt.Errorf("invalid record type: %w", err)
	}

	if err := utils.ValidateRecordContent(recordType, content); err != nil {
		return fmt.Errorf("invalid record content: %w", err)
	}

	if err := utils.ValidateTTL(ttl); err != nil {
		return fmt.Errorf("invalid TTL: %w", err)
	}

	return nil
}

func DNSCommand() *cli.Command {
	return &cli.Command{
		Name:  "dns",
		Usage: "DNS record management",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List DNS records",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:  "id",
						Usage: "Record ID(s) to list",
					},
					&cli.StringSliceFlag{
						Name:    "domain",
						Aliases: []string{"d"},
						Usage:   "Domain name(s)",
					},
					&cli.StringSliceFlag{
						Name:    "type",
						Aliases: []string{"t"},
						Usage:   "Record type(s)",
					},
					&cli.StringSliceFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "Record name(s)",
					},
					&cli.StringSliceFlag{
						Name:    "content",
						Aliases: []string{"c"},
						Usage:   "Record content(s)",
					},
					&cli.BoolFlag{
						Name:    "wildcard",
						Aliases: []string{"w"},
						Usage:   "Use shell-style wildcards for name matching",
					},
					&cli.IntFlag{
						Name:  "max",
						Usage: "Maximum number of records to display (0 = no limit)",
						Value: 0,
					},
				},
				Action: listDNSRecords,
			},
			{
				Name:  "create",
				Usage: "Create DNS record",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "domain",
						Aliases:  []string{"d"},
						Usage:    "Domain name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "type",
						Aliases:  []string{"t"},
						Usage:    "Record type",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "Record name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "content",
						Aliases:  []string{"c"},
						Usage:    "Record content",
						Required: true,
					},
					&cli.IntFlag{
						Name:  "ttl",
						Usage: "TTL value",
						Value: 3600,
					},
					&cli.IntFlag{
						Name:  "prio",
						Usage: "Priority (for MX records)",
					},
				},
				Action: createDNSRecord,
			},
			{
				Name:  "update",
				Usage: "Update DNS record",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "id",
						Usage:    "Record ID",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "type",
						Aliases: []string{"t"},
						Usage:   "Record type",
					},
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "Record name",
					},
					&cli.StringFlag{
						Name:    "content",
						Aliases: []string{"c"},
						Usage:   "Record content",
					},
					&cli.IntFlag{
						Name:  "ttl",
						Usage: "TTL value",
					},
					&cli.IntFlag{
						Name:  "prio",
						Usage: "Priority (for MX records)",
					},
				},
				Action: updateDNSRecord,
			},
			{
				Name:  "delete",
				Usage: "Delete DNS record(s)",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:  "id",
						Usage: "Record ID(s) to delete",
					},
					&cli.StringSliceFlag{
						Name:    "domain",
						Aliases: []string{"d"},
						Usage:   "Domain name(s)",
					},
					&cli.StringSliceFlag{
						Name:    "type",
						Aliases: []string{"t"},
						Usage:   "Record type(s)",
					},
					&cli.StringSliceFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "Record name(s)",
					},
					&cli.StringSliceFlag{
						Name:    "content",
						Aliases: []string{"c"},
						Usage:   "Record content(s)",
					},
					&cli.BoolFlag{
						Name:    "wildcard",
						Aliases: []string{"w"},
						Usage:   "Use shell-style wildcards for name matching",
					},
					&cli.BoolFlag{
						Name:    "dry-run",
						Aliases: []string{"R"},
						Usage:   "Show what would be deleted without actually deleting",
					},
					&cli.IntFlag{
						Name:  "max",
						Usage: "Maximum number of records to delete (0 = no limit)",
						Value: 0,
					},
				},
				Action: deleteDNSRecord,
			},
			{
				Name:  "export",
				Usage: "Export DNS records",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "domain",
						Aliases: []string{"d"},
						Usage:   "Domain name",
					},
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"f"},
						Usage:   "Export format (json, zonefile)",
						Value:   "json",
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "Output file",
					},
					&cli.StringFlag{
						Name:  "output-dir",
						Usage: "Output directory (exports one file per domain)",
					},
				},
				Action: exportDNSRecords,
			},
			{
				Name:  "import",
				Usage: "Import DNS records",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "Input file",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "domain",
						Aliases: []string{"d"},
						Usage:   "Domain name",
					},
					&cli.StringFlag{
						Name:  "format",
						Usage: "Import format (json, zonefile)",
						Value: "json",
					},
					&cli.BoolFlag{
						Name:    "dry-run",
						Aliases: []string{"R"},
						Usage:   "Dry run mode",
					},
					&cli.BoolFlag{
						Name:    "delete",
						Aliases: []string{"D"},
						Usage:   "Delete records not present in import file",
					},
				},
				Action: importDNSRecords,
			},
		},
	}
}

func listDNSRecords(c *cli.Context) error {
	log.Debug().Msg("Starting DNS list")

	// Parse flags first
	ids := parseCommaSeparatedValues(c.StringSlice("id"))
	domains := parseCommaSeparatedValues(c.StringSlice("domain"))
	types := parseCommaSeparatedValues(c.StringSlice("type"))
	names := parseCommaSeparatedValues(c.StringSlice("name"))
	contents := c.StringSlice("content") // Don't parse commas in content
	positionalHosts := c.Args().Slice()

	client, err := createClient(c)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create client")
		return err
	}

	ctx := context.Background()
	log.Debug().Msg("Attempting to login")
	if err := client.Login(ctx); err != nil {
		log.Error().Err(err).Msg("Login failed")
		return err
	}
	defer func() {
		log.Debug().Msg("Logging out")
		_ = client.Logout(ctx)
	}()

	useWildcard := c.Bool("wildcard")
	maxRecords := c.Int("max")

	// Get target records using the same logic as delete
	var targetRecords []inwx.DNSRecord

	if len(ids) > 0 {
		// If IDs are specified, get records by ID and then apply additional filters
		targetRecords, err = getRecordsByIDs(ctx, client, ids)
		if err != nil {
			return err
		}
		// Apply additional filters if IDs were specified
		targetRecords = applyAdditionalFilters(targetRecords, domains, types, names, contents, useWildcard)
	} else if len(positionalHosts) > 0 {
		// Handle positional arguments as fully-qualified hostnames
		targetRecords, err = getRecordsByHosts(ctx, client, positionalHosts)
		if err != nil {
			return err
		}
		// Apply additional filters to host-based records
		targetRecords = applyAdditionalFilters(targetRecords, domains, types, names, contents, useWildcard)
	} else {
		// Build filters and get records
		targetRecords, err = getRecordsByFilters(ctx, client, domains, types, names, contents, useWildcard)
		if err != nil {
			return err
		}
	}

	// Apply max limit if specified
	if maxRecords > 0 && len(targetRecords) > maxRecords {
		targetRecords = targetRecords[:maxRecords]
	}

	log.Debug().Int("record_count", len(targetRecords)).Msg("Records retrieved")

	return formatOutput(c, func(formatter interface{}) string {
		switch f := formatter.(type) {
		case *output.TableFormatter:
			return f.FormatDNSRecords(targetRecords)
		case *output.JSONFormatter:
			return f.FormatDNSRecords(targetRecords)
		case *output.YAMLFormatter:
			return f.FormatDNSRecords(targetRecords)
		default:
			return "Unsupported format"
		}
	})
}

func createDNSRecord(c *cli.Context) error {
	// Validate input before making API calls
	domain := c.String("domain")
	recordType := c.String("type")
	name := c.String("name")
	content := c.String("content")
	ttl := c.Int("ttl")

	if err := validateDNSRecordInput(domain, recordType, name, content, ttl); err != nil {
		return fmt.Errorf("validation failed: %w", err)
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
		_ = client.Logout(ctx)
	}()

	record := inwx.DNSRecord{
		Domain:  domain,
		Type:    recordType,
		Name:    name,
		Content: content,
		TTL:     ttl,
		Prio:    c.Int("prio"),
	}

	dns, err := createDNSService(c, client)
	if err != nil {
		return err
	}
	created, err := dns.CreateRecord(ctx, record)
	if err != nil {
		return err
	}

	log.Info().Msgf("Created DNS record with ID %d", created.ID)
	return nil
}

func updateDNSRecord(c *cli.Context) error {
	recordType := c.String("type")
	content := c.String("content")
	ttl := c.Int("ttl")

	if recordType != "" {
		if err := utils.ValidateRecordType(recordType); err != nil {
			return fmt.Errorf("invalid record type: %w", err)
		}
	}

	if content != "" && recordType != "" {
		if err := utils.ValidateRecordContent(recordType, content); err != nil {
			return fmt.Errorf("invalid record content: %w", err)
		}
	}

	if ttl > 0 {
		if err := utils.ValidateTTL(ttl); err != nil {
			return fmt.Errorf("invalid TTL: %w", err)
		}
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
		_ = client.Logout(ctx)
	}()

	updates := inwx.DNSRecord{
		Type:    recordType,
		Name:    c.String("name"),
		Content: content,
		TTL:     ttl,
		Prio:    c.Int("prio"),
	}

	dns, err := createDNSService(c, client)
	if err != nil {
		return err
	}
	_, err = dns.UpdateRecord(ctx, c.Int("id"), updates)
	if err != nil {
		return err
	}

	log.Info().Msgf("Updated DNS record with ID %d", c.Int("id"))
	return nil
}

func deleteDNSRecord(c *cli.Context) error {
	// Parse flags first
	ids := parseCommaSeparatedValues(c.StringSlice("id"))
	domains := parseCommaSeparatedValues(c.StringSlice("domain"))
	types := parseCommaSeparatedValues(c.StringSlice("type"))
	names := parseCommaSeparatedValues(c.StringSlice("name"))
	contents := c.StringSlice("content") // Don't parse commas in content
	positionalHosts := c.Args().Slice()

	// Safety check: require at least one filter to be specified
	if len(ids) == 0 && len(domains) == 0 && len(types) == 0 && len(names) == 0 && len(contents) == 0 && len(positionalHosts) == 0 {
		return fmt.Errorf("at least one filter must be specified (--id, --domain, --type, --name, --content, or positional hostname arguments)")
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
		_ = client.Logout(ctx)
	}()

	useWildcard := c.Bool("wildcard")
	dryRun := c.Bool("dry-run")
	maxRecords := c.Int("max")
	skipPrompt := c.Bool("yes")

	// Get target records
	var targetRecords []inwx.DNSRecord

	if len(ids) > 0 {
		// If IDs are specified, get records by ID and then apply additional filters
		targetRecords, err = getRecordsByIDs(ctx, client, ids)
		if err != nil {
			return err
		}
		// Apply additional filters if IDs were specified
		targetRecords = applyAdditionalFilters(targetRecords, domains, types, names, contents, useWildcard)
	} else if len(positionalHosts) > 0 {
		// Handle positional arguments as fully-qualified hostnames
		targetRecords, err = getRecordsByHosts(ctx, client, positionalHosts)
		if err != nil {
			return err
		}
		// Apply additional filters to host-based records
		targetRecords = applyAdditionalFilters(targetRecords, domains, types, names, contents, useWildcard)
	} else {
		// Build filters and get records
		targetRecords, err = getRecordsByFilters(ctx, client, domains, types, names, contents, useWildcard)
		if err != nil {
			return err
		}
	}

	if len(targetRecords) == 0 {
		fmt.Println("No records match the specified criteria")
		return nil
	}

	// Check safety limit
	if maxRecords > 0 && len(targetRecords) > maxRecords {
		fmt.Printf("Error: Found %d records, which exceeds the safety limit of %d.\n", len(targetRecords), maxRecords)
		fmt.Println("Matching records:")
		printRecordTable(targetRecords)
		fmt.Printf("\nPlease refine your filters or increase --max to proceed.\n")
		return fmt.Errorf("too many records to delete (limit: %d)", maxRecords)
	}

	// Show what will be deleted
	fmt.Printf("Deleting %d entries:\n", len(targetRecords))
	printRecordTable(targetRecords)

	// Dry run handling
	if dryRun {
		fmt.Println("\nDry run mode - no records were actually deleted")
		return nil
	}

	// User confirmation
	if !skipPrompt {
		confirmed, err := utils.AskSimpleConfirmation("Continue (y/N)?", false)
		if err != nil {
			return err
		}
		if !confirmed {
			fmt.Println("Operation cancelled")
			return nil
		}
	}

	// Delete records
	dns, err := createDNSService(c, client)
	if err != nil {
		return err
	}
	deleted := 0
	for _, record := range targetRecords {
		if err := dns.DeleteRecord(ctx, record.ID); err != nil {
			log.Warn().Err(err).Int("id", record.ID).Msg("Failed to delete record")
		} else {
			deleted++
		}
	}

	log.Info().Msgf("Deleted %d DNS records", deleted)
	return nil
}

func exportDNSRecords(c *cli.Context) error {
	log.Debug().Msg("Starting DNS export")

	client, err := createClient(c)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create client")
		return err
	}

	ctx := context.Background()
	log.Debug().Msg("Attempting to login")
	if err := client.Login(ctx); err != nil {
		log.Error().Err(err).Msg("Login failed")
		return err
	}
	defer func() {
		log.Debug().Msg("Logging out")
		client.Logout(ctx)
	}()

	domain := c.String("domain")
	outputDir := c.String("output-dir")
	outputFile := c.String("output")

	var format inwx.ExportFormat
	switch c.String("format") {
	case "json":
		format = inwx.ExportJSON
	case "zonefile":
		format = inwx.ExportZonefileFormat
	default:
		return fmt.Errorf("unsupported format: %s", c.String("format"))
	}

	// Export to directory (one file per domain)
	if outputDir != "" {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Get list of domains to export
		var domains []string
		if domain != "" {
			domains = []string{domain}
		} else {
			domainService := client.Domain()
			domainList, err := domainService.List(ctx)
			if err != nil {
				return fmt.Errorf("failed to list domains: %w", err)
			}
			for _, d := range domainList {
				domains = append(domains, d.Name)
			}
		}

		// Export each domain
		for _, d := range domains {
			// Initialize backup store
			backupStore, err := backup.NewStore()
			if err != nil {
				log.Warn().Err(err).Str("domain", d).Msg("Failed to initialize backup store for domain")
				continue
			}
			dns := client.DNS(inwx.WithDomain(d), inwx.WithBackupStore(backupStore))
			data, err := dns.ExportRecords(ctx, format)
			if err != nil {
				log.Warn().Err(err).Str("domain", d).Msg("Failed to export domain")
				continue
			}

			var ext string
			switch format {
			case inwx.ExportJSON:
				ext = "json"
			case inwx.ExportZonefileFormat:
				ext = "zone"
			}

			filename := filepath.Join(outputDir, fmt.Sprintf("%s.%s", d, ext))
			if err := os.WriteFile(filename, data, 0644); err != nil {
				log.Warn().Err(err).Str("file", filename).Msg("Failed to write file")
				continue
			}

			log.Info().Str("domain", d).Str("file", filename).Msg("Exported domain")
		}

		return nil
	}

	// Single domain export
	if domain == "" {
		return fmt.Errorf("domain is required when not using --output-dir")
	}

	log.Debug().Str("domain", domain).Msg("Exporting DNS records")

	// Initialize backup store
	backupStore, err := backup.NewStore()
	if err != nil {
		return err
	}
	dns := client.DNS(inwx.WithDomain(domain), inwx.WithBackupStore(backupStore))
	data, err := dns.ExportRecords(ctx, format)
	if err != nil {
		log.Error().Err(err).Msg("Failed to export records")
		return err
	}

	log.Debug().Int("bytes", len(data)).Msg("Export completed")

	if outputFile != "" {
		log.Debug().Str("file", outputFile).Msg("Writing to file")
		return os.WriteFile(outputFile, data, 0644)
	}

	fmt.Print(string(data))
	return nil
}

func importDNSRecords(c *cli.Context) error {
	data, err := os.ReadFile(c.String("file"))
	if err != nil {
		return err
	}

	// Determine domain and format
	domain := c.String("domain")
	var format inwx.ImportFormat
	switch c.String("format") {
	case "json":
		format = inwx.ImportJSON
	case "zonefile":
		format = inwx.ImportZonefileFormat
	default:
		return fmt.Errorf("unsupported format: %s", c.String("format"))
	}

	// For zonefile format, try to detect domain from $ORIGIN directive
	if format == inwx.ImportZonefileFormat && domain == "" {
		domain, err = detectDomainFromZonefile(data)
		if err != nil {
			return fmt.Errorf("could not detect domain from zonefile and --domain not specified: %w", err)
		}
	}

	if domain == "" {
		return fmt.Errorf("domain must be specified via --domain flag or detected from file")
	}

	client, err := createClient(c)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := client.Login(ctx); err != nil {
		return err
	}
	defer client.Logout(ctx)

	// Parse records to import
	var recordsToImport []inwx.DNSRecord
	switch format {
	case inwx.ImportJSON:
		recordsToImport, err = importJSON(data)
	case inwx.ImportZonefileFormat:
		recordsToImport, err = inwx.ImportZonefile(data, domain)
	}
	if err != nil {
		return fmt.Errorf("failed to parse import file: %w", err)
	}

	if c.Bool("dry-run") {
		log.Info().Msgf("Dry run mode - would import %d records for domain %s", len(recordsToImport), domain)
		if c.Bool("delete") {
			log.Info().Msg("Would also delete records not present in import file")
		}
		return nil
	}

	// Initialize backup store
	backupStore, err := backup.NewStore()
	if err != nil {
		return err
	}
	dns := client.DNS(inwx.WithDomain(domain), inwx.WithBackupStore(backupStore))

	// Handle --delete flag
	if c.Bool("delete") {
		skipPrompt := c.Bool("yes")
		confirmed, err := utils.AskSimpleConfirmation(
			fmt.Sprintf("This will delete existing records for %s not present in the import file. Continue?", domain),
			skipPrompt,
		)
		if err != nil {
			return err
		}
		if !confirmed {
			fmt.Println("Operation cancelled")
			return nil
		}

		err = dns.ImportRecordsWithSync(ctx, recordsToImport, format)
		if err != nil {
			return err
		}
	} else {
		err = dns.ImportRecords(ctx, data, format)
		if err != nil {
			return err
		}
	}

	log.Info().Msgf("Successfully imported %d DNS records for domain %s", len(recordsToImport), domain)
	return nil
}

// getRecordsByIDs retrieves records by their IDs
func getRecordsByIDs(ctx context.Context, client *inwx.Client, ids []string) ([]inwx.DNSRecord, error) {
	var records []inwx.DNSRecord
	// Initialize backup store
	backupStore, err := backup.NewStore()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize backup store: %w", err)
	}
	dns := client.DNS(inwx.WithBackupStore(backupStore))

	// Get all records first (since we need to search by ID)
	allRecords, err := dns.ListRecords(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list records: %w", err)
	}

	// Convert string IDs to integers and find matching records
	for _, idStr := range ids {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Warn().Str("id", idStr).Msg("Invalid record ID, skipping")
			continue
		}

		for _, record := range allRecords {
			if record.ID == id {
				records = append(records, record)
				break
			}
		}
	}

	return records, nil
}

// getRecordsByFilters retrieves records using API filters and local filtering
func getRecordsByFilters(ctx context.Context, client *inwx.Client, domains, types, names, contents []string, useWildcard bool) ([]inwx.DNSRecord, error) {
	var allRecords []inwx.DNSRecord

	// Determine target domains
	targetDomains := domains
	if len(targetDomains) == 0 {
		// Get all owned domains
		domainService := client.Domain()
		domainList, err := domainService.List(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list domains: %w", err)
		}
		for _, d := range domainList {
			targetDomains = append(targetDomains, d.Name)
		}
	}

	// Initialize backup store
	backupStore, err := backup.NewStore()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize backup store: %w", err)
	}

	// Get records for each domain
	dns := client.DNS(inwx.WithBackupStore(backupStore))
	for _, domain := range targetDomains {
		var filters []inwx.RecordFilter
		filters = append(filters, inwx.WithDomainFilter(domain))

		// Apply type filter if specified
		if len(types) > 0 {
			filters = append(filters, inwx.WithRecordType(types...))
		}

		records, err := dns.ListRecords(ctx, filters...)
		if err != nil {
			log.Warn().Err(err).Str("domain", domain).Msg("Failed to get records for domain, skipping")
			continue
		}

		allRecords = append(allRecords, records...)
	}

	// Apply local filters
	return applyLocalFilters(allRecords, names, contents, useWildcard), nil
}

// applyAdditionalFilters applies domain, type, name, and content filters to records
func applyAdditionalFilters(records []inwx.DNSRecord, domains, types, names, contents []string, useWildcard bool) []inwx.DNSRecord {
	var filtered []inwx.DNSRecord

	for _, record := range records {
		// Domain filter
		if len(domains) > 0 && !containsIgnoreCase(domains, record.Domain) {
			continue
		}

		// Type filter
		if len(types) > 0 && !containsIgnoreCase(types, record.Type) {
			continue
		}

		// Name and content filters
		if !matchesLocalFilters(record, names, contents, useWildcard) {
			continue
		}

		filtered = append(filtered, record)
	}

	return filtered
}

// applyLocalFilters applies name and content filters locally
func applyLocalFilters(records []inwx.DNSRecord, names, contents []string, useWildcard bool) []inwx.DNSRecord {
	var filtered []inwx.DNSRecord

	for _, record := range records {
		if matchesLocalFilters(record, names, contents, useWildcard) {
			filtered = append(filtered, record)
		}
	}

	return filtered
}

// matchesLocalFilters checks if a record matches name and content filters
func matchesLocalFilters(record inwx.DNSRecord, names, contents []string, useWildcard bool) bool {
	// Name filter
	if len(names) > 0 {
		nameMatches := false
		recordName := record.Name
		if recordName == "" {
			recordName = "@"
		}

		for _, name := range names {
			if useWildcard {
				if matched, _ := filepath.Match(strings.ToLower(name), strings.ToLower(recordName)); matched {
					nameMatches = true
					break
				}
			} else {
				if strings.EqualFold(name, recordName) {
					nameMatches = true
					break
				}
			}
		}
		if !nameMatches {
			return false
		}
	}

	// Content filter
	if len(contents) > 0 {
		contentMatches := false
		for _, content := range contents {
			if useWildcard {
				if matched, _ := filepath.Match(strings.ToLower(content), strings.ToLower(record.Content)); matched {
					contentMatches = true
					break
				}
			} else {
				if strings.Contains(strings.ToLower(record.Content), strings.ToLower(content)) {
					contentMatches = true
					break
				}
			}
		}
		if !contentMatches {
			return false
		}
	}

	return true
}

// containsIgnoreCase checks if a slice contains a string (case-insensitive)
func containsIgnoreCase(slice []string, item string) bool {
	itemLower := strings.ToLower(item)
	for _, s := range slice {
		if strings.ToLower(s) == itemLower {
			return true
		}
	}
	return false
}

// getRecordsByHosts retrieves records for specific hostnames
func getRecordsByHosts(ctx context.Context, client *inwx.Client, hosts []string) ([]inwx.DNSRecord, error) {
	var allRecords []inwx.DNSRecord
	// Initialize backup store
	backupStore, err := backup.NewStore()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize backup store: %w", err)
	}
	dns := client.DNS(inwx.WithBackupStore(backupStore))

	for _, host := range hosts {
		domain, name, err := utils.InferDomainAndName(client, ctx, host)
		if err != nil {
			return nil, fmt.Errorf("failed to infer domain for host: %s", host)
		}

		// Build filters for this specific host
		var filters []inwx.RecordFilter
		filters = append(filters, inwx.WithDomainFilter(domain))
		if name != "" {
			filters = append(filters, inwx.WithRecordName(name))
		}

		// Get records for this host
		records, err := dns.ListRecords(ctx, filters...)
		if err != nil {
			log.Warn().Err(err).Str("host", host).Msg("Failed to get records for host, skipping")
			continue
		}

		if len(records) == 0 {
			return nil, fmt.Errorf("no DNS records found for host: %s", host)
		}

		// Filter to exact matches only (in case the API filter was too broad)
		for _, record := range records {
			var fullHostname string
			if record.Name == "" || record.Name == "@" {
				fullHostname = record.Domain
			} else {
				fullHostname = record.Name + "." + record.Domain
			}

			if strings.EqualFold(fullHostname, host) {
				allRecords = append(allRecords, record)
			}
		}
	}

	return allRecords, nil
}

// printRecordTable prints a concise table of records
func printRecordTable(records []inwx.DNSRecord) {
	if len(records) == 0 {
		return
	}

	fmt.Printf("%-8s %-20s %-15s %-8s %s\n", "ID", "Domain", "Name", "Type", "Content")
	fmt.Println(strings.Repeat("-", 80))

	for _, record := range records {
		name := record.Name
		if name == "" {
			name = "@"
		}

		// Truncate long content for display
		content := record.Content
		if len(content) > 40 {
			content = content[:37] + "..."
		}

		fmt.Printf("%-8d %-20s %-15s %-8s %s\n",
			record.ID,
			record.Domain,
			name,
			record.Type,
			content)
	}
}

func detectDomainFromZonefile(data []byte) (string, error) {
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "$ORIGIN") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				domain := strings.TrimSuffix(parts[1], ".")
				return domain, nil
			}
		}
	}
	return "", fmt.Errorf("no $ORIGIN directive found in zonefile")
}

func importJSON(data []byte) ([]inwx.DNSRecord, error) {
	var records []inwx.DNSRecord
	err := json.Unmarshal(data, &records)
	return records, err
}
