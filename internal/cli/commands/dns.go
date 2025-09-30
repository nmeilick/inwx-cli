package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
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
				Name:      "create",
				Usage:     "Create DNS record",
				ArgsUsage: "[hostname]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "domain",
						Aliases: []string{"d"},
						Usage:   "Domain name",
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
						Value: 3600,
					},
					&cli.IntFlag{
						Name:  "prio",
						Usage: "Priority (for MX records)",
					},
					&cli.BoolFlag{
						Name:    "dry-run",
						Aliases: []string{"R"},
						Usage:   "Show what would be created without actually creating",
					},
					&cli.BoolFlag{
						Name:  "wait",
						Usage: "Wait for DNS propagation to authoritative nameservers (120s timeout)",
					},
					&cli.BoolFlag{
						Name:    "interactive",
						Aliases: []string{"i"},
						Usage:   "Interactive mode - prompts for all record details",
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
					&cli.BoolFlag{
						Name:    "dry-run",
						Aliases: []string{"R"},
						Usage:   "Show what would be updated without actually updating",
					},
					&cli.BoolFlag{
						Name:  "wait",
						Usage: "Wait for DNS propagation to authoritative nameservers (120s timeout)",
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
				Name:      "export",
				Usage:     "Export DNS records",
				ArgsUsage: "[domain...]",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "domain",
						Aliases: []string{"d"},
						Usage:   "Domain name(s)",
					},
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"f"},
						Usage:   "Export format (json, zonefile)",
						Value:   "zonefile",
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "Output file (optional for single domain, defaults to stdout)",
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
			{
				Name:      "edit",
				Usage:     "Edit DNS records in your $EDITOR",
				ArgsUsage: "[domain]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "format",
						Usage: "Edit format (json, zonefile)",
						Value: "zonefile",
					},
				},
				Action: editDNSRecords,
			},
			{
				Name:  "validate",
				Usage: "Validate DNS records for configuration issues",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "domain",
						Aliases: []string{"d"},
						Usage:   "Domain(s) to validate (validates all domains if not specified)",
					},
					&cli.StringFlag{
						Name:  "severity",
						Usage: "Minimum severity to report (error, warning, info)",
						Value: "warning",
					},
				},
				Action: validateDNSRecords,
			},
			{
				Name:      "verify",
				Usage:     "Verify DNS records are live and propagated",
				ArgsUsage: "[domain|hostname...]",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "domain",
						Aliases: []string{"d"},
						Usage:   "Domain(s) to verify",
					},
					&cli.StringSliceFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "Record name(s) to verify (@ for root)",
					},
					&cli.StringSliceFlag{
						Name:    "type",
						Aliases: []string{"t"},
						Usage:   "Record type(s) to verify",
					},
					&cli.DurationFlag{
						Name:  "wait",
						Usage: "Wait for propagation (e.g., 5m, 30s)",
						Value: 0,
					},
				},
				Action: verifyDNSRecords,
			},
		},
	}
}

// classifyAndDeduplicateTargets separates positional arguments into domains (for full listing)
// and hosts (for specific records), with deduplication logic:
// - Removes duplicate domains
// - Removes hosts whose domain is already in the domains list
// - Validates all domains against owned domains
func classifyAndDeduplicateTargets(ctx context.Context, client *inwx.Client, args []string) (domains []string, hosts []string, err error) {
	if len(args) == 0 {
		return nil, nil, nil
	}

	// Get list of owned domains
	domainService := client.Domain()
	ownedDomains, err := domainService.List(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list owned domains: %w", err)
	}

	// Create a map for quick lookup
	ownedDomainsMap := make(map[string]bool)
	for _, d := range ownedDomains {
		ownedDomainsMap[strings.ToLower(d.Name)] = true
	}

	domainsMap := make(map[string]bool)
	var hostsList []string

	for _, arg := range args {
		argLower := strings.ToLower(arg)

		// Check if this is an owned domain
		if ownedDomainsMap[argLower] {
			// It's a domain - add to domains map (automatic deduplication)
			domainsMap[argLower] = true
		} else {
			// It's either a host or invalid domain
			// Try to infer domain and name
			_, _, inferErr := utils.InferDomainAndName(client, ctx, arg)
			if inferErr != nil {
				// Cannot infer domain - this is an invalid/unowned domain
				return nil, nil, fmt.Errorf("invalid or unowned domain/host: %s", arg)
			}
			// It's a valid host - add to hosts list (we'll deduplicate later)
			hostsList = append(hostsList, arg)
		}
	}

	// Convert domains map to slice
	var domainsList []string
	for domain := range domainsMap {
		domainsList = append(domainsList, domain)
	}

	// Deduplicate hosts: remove hosts whose domain is in the domains list
	// and remove duplicate hosts
	hostsMap := make(map[string]bool)
	for _, host := range hostsList {
		domain, _, err := utils.InferDomainAndName(client, ctx, host)
		if err != nil {
			continue // Skip if we can't infer (shouldn't happen since we checked above)
		}
		domainLower := strings.ToLower(domain)

		// Skip this host if its domain is in the domains list
		if domainsMap[domainLower] {
			log.Debug().
				Str("host", host).
				Str("domain", domain).
				Msg("Skipping host because its domain is already being listed")
			continue
		}

		// Add to hosts map for deduplication
		hostLower := strings.ToLower(host)
		hostsMap[hostLower] = true
	}

	// Convert hosts map to slice
	var deduplicatedHosts []string
	for host := range hostsMap {
		deduplicatedHosts = append(deduplicatedHosts, host)
	}

	return domainsList, deduplicatedHosts, nil
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
		if err := client.Logout(ctx); err != nil {
			log.Error().Err(err).Msg("Failed to logout")
		}
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
		// Handle positional arguments as both domains and hosts
		// Classify and deduplicate the arguments
		targetDomains, targetHosts, err := classifyAndDeduplicateTargets(ctx, client, positionalHosts)
		if err != nil {
			return err
		}

		log.Debug().
			Strs("domains", targetDomains).
			Strs("hosts", targetHosts).
			Msg("Classified and deduplicated positional arguments")

		// Initialize backup store
		backupStore, err := backup.NewStore()
		if err != nil {
			return fmt.Errorf("failed to initialize backup store: %w", err)
		}
		dns := client.DNS(inwx.WithBackupStore(backupStore))

		// Use a map to deduplicate records by ID
		recordsMap := make(map[int]inwx.DNSRecord)

		// Get all records for each domain
		for _, domain := range targetDomains {
			filters := []inwx.RecordFilter{inwx.WithDomainFilter(domain)}
			records, err := dns.ListRecords(ctx, filters...)
			if err != nil {
				log.Warn().Err(err).Str("domain", domain).Msg("Failed to get records for domain, skipping")
				continue
			}
			for _, record := range records {
				recordsMap[record.ID] = record
			}
		}

		// Get specific records for each host
		for _, host := range targetHosts {
			domain, name, err := utils.InferDomainAndName(client, ctx, host)
			if err != nil {
				log.Warn().Err(err).Str("host", host).Msg("Failed to infer domain for host, skipping")
				continue
			}

			filters := []inwx.RecordFilter{inwx.WithDomainFilter(domain)}
			if name != "" && name != "@" {
				filters = append(filters, inwx.WithRecordName(name))
			}

			records, err := dns.ListRecords(ctx, filters...)
			if err != nil {
				log.Warn().Err(err).Str("host", host).Msg("Failed to get records for host, skipping")
				continue
			}

			// Filter to exact matches only
			for _, record := range records {
				var fullHostname string
				if record.Name == "" || record.Name == "@" {
					fullHostname = record.Domain
				} else {
					fullHostname = record.Name + "." + record.Domain
				}

				if strings.EqualFold(fullHostname, host) {
					recordsMap[record.ID] = record
				}
			}
		}

		// Convert map to slice
		for _, record := range recordsMap {
			targetRecords = append(targetRecords, record)
		}

		// Apply additional filters to the combined results
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

// createDNSRecordInteractive handles interactive DNS record creation
func createDNSRecordInteractive(c *cli.Context) error {
	fmt.Println("\n✨ Interactive DNS Record Creation")

	// Create client to get owned domains
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

	// Get owned domains for suggestions
	domainService := client.Domain()
	ownedDomains, err := domainService.List(ctx)
	if err != nil {
		return fmt.Errorf("failed to list owned domains: %w", err)
	}

	domainNames := make([]string, len(ownedDomains))
	for i, d := range ownedDomains {
		domainNames[i] = d.Name
	}

	// Prompt for hostname
	var hostname string
	err = survey.AskOne(&survey.Input{
		Message: "Hostname (e.g., www.example.com or example.com for root):",
		Help:    "Enter the full hostname including domain, or just the domain for a root record",
	}, &hostname, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}

	// Infer domain and name
	domain, name, err := utils.InferDomainAndName(client, ctx, hostname)
	if err != nil {
		return fmt.Errorf("invalid hostname: %w", err)
	}

	// Prompt for record type
	var recordType string
	err = survey.AskOne(&survey.Select{
		Message: "Record type:",
		Options: []string{"A", "AAAA", "CNAME", "MX", "TXT", "NS", "SRV", "CAA"},
		Default: "A",
	}, &recordType)
	if err != nil {
		return err
	}

	// Prompt for content with type-specific help
	var content string
	var contentHelp string
	switch recordType {
	case "A":
		contentHelp = "IPv4 address (e.g., 192.0.2.1)"
	case "AAAA":
		contentHelp = "IPv6 address (e.g., 2001:db8::1)"
	case "CNAME":
		contentHelp = "Target hostname (e.g., example.com)"
	case "MX":
		contentHelp = "Mail server hostname (priority will be asked separately)"
	case "TXT":
		contentHelp = "Text content (e.g., for SPF, DKIM, verification)"
	case "NS":
		contentHelp = "Nameserver hostname"
	case "SRV":
		contentHelp = "Target hostname (priority, weight, port will be asked separately)"
	case "CAA":
		contentHelp = "CAA record value (e.g., 0 issue \"letsencrypt.org\")"
	default:
		contentHelp = "Record content"
	}

	contentPrompt := &survey.Input{
		Message: fmt.Sprintf("Content (%s):", recordType),
		Help:    contentHelp,
	}

	// Add validator for content
	err = survey.AskOne(contentPrompt, &content, survey.WithValidator(func(val interface{}) error {
		str, ok := val.(string)
		if !ok || str == "" {
			return fmt.Errorf("content is required")
		}
		// Validate content based on type
		if err := utils.ValidateRecordContent(recordType, str); err != nil {
			return err
		}
		return nil
	}))
	if err != nil {
		return err
	}

	// Prompt for TTL
	var ttlStr string
	err = survey.AskOne(&survey.Input{
		Message: "TTL (seconds):",
		Default: "3600",
		Help:    "Time To Live - how long DNS resolvers should cache this record (common: 300, 3600, 86400)",
	}, &ttlStr, survey.WithValidator(func(val interface{}) error {
		str, ok := val.(string)
		if !ok {
			return fmt.Errorf("invalid input")
		}
		ttl, err := strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf("TTL must be a number")
		}
		return utils.ValidateTTL(ttl)
	}))
	if err != nil {
		return err
	}
	ttl, _ := strconv.Atoi(ttlStr)

	// Prompt for priority (only for MX and SRV)
	var prio int
	if recordType == "MX" || recordType == "SRV" {
		var prioStr string
		prioHelp := "Lower values have higher priority (e.g., 10)"
		if recordType == "SRV" {
			prioHelp = "Priority value for SRV record"
		}

		err = survey.AskOne(&survey.Input{
			Message: "Priority:",
			Default: "10",
			Help:    prioHelp,
		}, &prioStr, survey.WithValidator(func(val interface{}) error {
			str, ok := val.(string)
			if !ok {
				return fmt.Errorf("invalid input")
			}
			_, err := strconv.Atoi(str)
			if err != nil {
				return fmt.Errorf("priority must be a number")
			}
			return nil
		}))
		if err != nil {
			return err
		}
		prio, _ = strconv.Atoi(prioStr)
	}

	// Show preview
	fmt.Println("\n" + strings.Repeat("─", 50))
	fmt.Println("📋 Record Preview:")
	fmt.Println(strings.Repeat("─", 50))
	fmt.Printf("  Domain:   %s\n", domain)
	fmt.Printf("  Name:     %s\n", name)
	fmt.Printf("  Type:     %s\n", recordType)
	fmt.Printf("  Content:  %s\n", content)
	fmt.Printf("  TTL:      %d seconds\n", ttl)
	if prio > 0 {
		fmt.Printf("  Priority: %d\n", prio)
	}

	fullName := name
	if fullName == "" || fullName == "@" {
		fullName = domain
	} else {
		fullName = name + "." + domain
	}
	fmt.Printf("  Full:     %s\n", fullName)
	fmt.Println(strings.Repeat("─", 50))

	// Confirm creation
	var confirm bool
	err = survey.AskOne(&survey.Confirm{
		Message: "Create this DNS record?",
		Default: true,
	}, &confirm)
	if err != nil {
		return err
	}

	if !confirm {
		fmt.Println("❌ Record creation cancelled")
		return nil
	}

	// Create the record
	record := inwx.DNSRecord{
		Domain:  domain,
		Type:    recordType,
		Name:    name,
		Content: content,
		TTL:     ttl,
		Prio:    prio,
	}

	dns, err := createDNSService(c, client)
	if err != nil {
		return err
	}

	created, err := dns.CreateRecord(ctx, record)
	if err != nil {
		return err
	}

	// Display success message
	displayRecordCreated(created)

	// Ask about propagation wait
	var waitForPropagation bool
	err = survey.AskOne(&survey.Confirm{
		Message: "Wait for DNS propagation to authoritative nameservers? (120s timeout)",
		Default: false,
	}, &waitForPropagation)
	if err != nil {
		return err
	}

	if waitForPropagation {
		fmt.Println("\nWaiting for DNS propagation to authoritative nameservers...")
		err := waitForDNSPropagation(ctx, dns, domain, name, recordType, 120*time.Second)
		if err != nil {
			return fmt.Errorf("record created but propagation verification failed: %w", err)
		}
		if isatty(os.Stdout) {
			fmt.Println("\033[32m✓\033[0m Record propagated to all authoritative nameservers")
		} else {
			fmt.Println("✓ Record propagated to all authoritative nameservers")
		}
	}

	return nil
}

func createDNSRecord(c *cli.Context) error {
	// Auto-start interactive mode if no arguments or flags provided
	hasArgs := c.NArg() > 0
	hasFlags := c.String("domain") != "" || c.String("name") != "" ||
		c.String("type") != "" || c.String("content") != ""

	if !hasArgs && !hasFlags && !c.Bool("interactive") {
		// No arguments or flags - start interactive mode automatically
		return createDNSRecordInteractive(c)
	}

	// Handle explicit interactive mode
	if c.Bool("interactive") {
		return createDNSRecordInteractive(c)
	}

	// Parse hostname from positional argument or flags
	domain := c.String("domain")
	name := c.String("name")
	recordType := c.String("type")
	content := c.String("content")
	ttl := c.Int("ttl")
	prio := c.Int("prio")
	dryRun := c.Bool("dry-run")

	// Validate required fields for non-interactive mode
	if recordType == "" {
		return fmt.Errorf("--type is required (or use --interactive mode)")
	}
	if content == "" {
		return fmt.Errorf("--content is required (or use --interactive mode)")
	}

	// Create client first (needed for hostname inference)
	client, err := createClient(c)
	if err != nil {
		return err
	}

	ctx := context.Background()

	// Check if hostname provided as positional argument
	if c.NArg() > 0 {
		hostname := c.Args().First()

		// Login to infer domain/name
		if err := client.Login(ctx); err != nil {
			return err
		}
		defer func() {
			if err := client.Logout(ctx); err != nil {
				log.Error().Err(err).Msg("Failed to logout")
			}
		}()

		inferredDomain, inferredName, err := utils.InferDomainAndName(client, ctx, hostname)
		if err != nil {
			return fmt.Errorf("invalid hostname %s: %w", hostname, err)
		}

		// Use inferred values (override flags if they were set)
		domain = inferredDomain
		name = inferredName

		log.Debug().
			Str("hostname", hostname).
			Str("domain", domain).
			Str("name", name).
			Msg("Parsed hostname from positional argument")
	} else if domain == "" || name == "" {
		return fmt.Errorf("hostname must be provided as positional argument or via --domain/-d and --name/-n flags")
	}

	if err := validateDNSRecordInput(domain, recordType, name, content, ttl); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Dry-run mode: show what would be created
	if dryRun {
		log.Info().Msg("Dry run mode - would create DNS record:")
		log.Info().Msgf("  Domain:  %s", domain)
		log.Info().Msgf("  Type:    %s", recordType)
		log.Info().Msgf("  Name:    %s", name)
		log.Info().Msgf("  Content: %s", content)
		log.Info().Msgf("  TTL:     %d", ttl)
		if prio > 0 {
			log.Info().Msgf("  Priority: %d", prio)
		}
		return nil
	}

	// Login if not already logged in (from hostname inference)
	if c.NArg() == 0 {
		if err := client.Login(ctx); err != nil {
			return err
		}
		defer func() {
			if err := client.Logout(ctx); err != nil {
				log.Error().Err(err).Msg("Failed to logout")
			}
		}()
	}

	record := inwx.DNSRecord{
		Domain:  domain,
		Type:    recordType,
		Name:    name,
		Content: content,
		TTL:     ttl,
		Prio:    prio,
	}

	dns, err := createDNSService(c, client)
	if err != nil {
		return err
	}
	created, err := dns.CreateRecord(ctx, record)
	if err != nil {
		return err
	}

	// Display success message with record details
	displayRecordCreated(created)

	// Wait for propagation if requested
	if c.Bool("wait") {
		fmt.Println("\nWaiting for DNS propagation to authoritative nameservers...")
		err := waitForDNSPropagation(ctx, dns, domain, name, recordType, 120*time.Second)
		if err != nil {
			return fmt.Errorf("record created but propagation verification failed: %w", err)
		}
		if isatty(os.Stdout) {
			fmt.Println("\033[32m✓\033[0m Record propagated to all authoritative nameservers")
		} else {
			fmt.Println("✓ Record propagated to all authoritative nameservers")
		}
	}

	return nil
}

func updateDNSRecord(c *cli.Context) error {
	recordID := c.Int("id")
	recordType := c.String("type")
	name := c.String("name")
	content := c.String("content")
	ttl := c.Int("ttl")
	prio := c.Int("prio")
	dryRun := c.Bool("dry-run")

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

	// Dry-run mode: show what would be updated
	if dryRun {
		log.Info().Msgf("Dry run mode - would update DNS record ID %d:", recordID)
		if recordType != "" {
			log.Info().Msgf("  Type:    %s", recordType)
		}
		if name != "" {
			log.Info().Msgf("  Name:    %s", name)
		}
		if content != "" {
			log.Info().Msgf("  Content: %s", content)
		}
		if ttl > 0 {
			log.Info().Msgf("  TTL:     %d", ttl)
		}
		if prio > 0 {
			log.Info().Msgf("  Priority: %d", prio)
		}
		return nil
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

	updates := inwx.DNSRecord{
		Type:    recordType,
		Name:    name,
		Content: content,
		TTL:     ttl,
		Prio:    prio,
	}

	dns, err := createDNSService(c, client)
	if err != nil {
		return err
	}

	// Get the original record to know its domain/name for verification
	originalRecord, err := dns.GetRecord(ctx, recordID)
	if err != nil {
		return fmt.Errorf("failed to get record: %w", err)
	}

	updated, err := dns.UpdateRecord(ctx, recordID, updates)
	if err != nil {
		return err
	}

	// Display success message with record details
	displayRecordUpdated(updated)

	// Wait for propagation if requested
	if c.Bool("wait") {
		// Use updated values if they were changed, otherwise use original values
		verifyDomain := updated.Domain
		verifyName := updated.Name
		verifyType := updated.Type

		if verifyName == "" {
			verifyName = originalRecord.Name
		}

		fmt.Println("\nWaiting for DNS propagation to authoritative nameservers...")
		err := waitForDNSPropagation(ctx, dns, verifyDomain, verifyName, verifyType, 120*time.Second)
		if err != nil {
			return fmt.Errorf("record updated but propagation verification failed: %w", err)
		}
		if isatty(os.Stdout) {
			fmt.Println("\033[32m✓\033[0m Record propagated to all authoritative nameservers")
		} else {
			fmt.Println("✓ Record propagated to all authoritative nameservers")
		}
	}

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
		if err := client.Logout(ctx); err != nil {
			log.Error().Err(err).Msg("Failed to logout")
		}
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
		if err := client.Logout(ctx); err != nil {
			log.Error().Err(err).Msg("Failed to logout")
		}
	}()

	// Collect domains from flags and positional arguments
	domainFlags := parseCommaSeparatedValues(c.StringSlice("domain"))
	domainArgs := c.Args().Slice()

	// Combine and deduplicate domains
	domainsMap := make(map[string]bool)
	for _, d := range domainFlags {
		domainsMap[strings.ToLower(d)] = true
	}
	for _, d := range domainArgs {
		domainsMap[strings.ToLower(d)] = true
	}

	var domains []string
	for d := range domainsMap {
		domains = append(domains, d)
	}

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

	// If no domains specified, export all owned domains
	if len(domains) == 0 {
		domainService := client.Domain()
		domainList, err := domainService.List(ctx)
		if err != nil {
			return fmt.Errorf("failed to list domains: %w", err)
		}
		for _, d := range domainList {
			domains = append(domains, d.Name)
		}
	}

	// Export to directory (one file per domain)
	if outputDir != "" {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
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

	// Single or multiple domain export (not to directory)
	if len(domains) > 1 && outputFile == "" {
		return fmt.Errorf("when exporting multiple domains, use --output-dir or specify only one domain")
	}

	// Single domain export
	domain := domains[0]
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

	// Output to file if specified, otherwise to stdout
	if outputFile != "" {
		log.Debug().Str("file", outputFile).Msg("Writing to file")
		return os.WriteFile(outputFile, data, 0644)
	}

	// Output to stdout
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
	defer func() {
		if err := client.Logout(ctx); err != nil {
			log.Error().Err(err).Msg("Failed to logout")
		}
	}()

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

	// Use API's direct lookup by ID (more efficient than listing all records)
	for _, idStr := range ids {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Warn().Str("id", idStr).Msg("Invalid record ID, skipping")
			continue
		}

		record, err := dns.GetRecord(ctx, id)
		if err != nil {
			log.Warn().Int("id", id).Err(err).Msg("Failed to get record by ID, skipping")
			continue
		}

		records = append(records, *record)
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

func editDNSRecords(c *cli.Context) error {
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

	// Determine which domain to edit
	var domain string
	if c.NArg() > 0 {
		domain = c.Args().First()
	} else {
		// Get list of owned domains
		domainService := client.Domain()
		ownedDomains, err := domainService.List(ctx)
		if err != nil {
			return fmt.Errorf("failed to list domains: %w", err)
		}

		if len(ownedDomains) == 0 {
			return fmt.Errorf("no domains found")
		}

		if len(ownedDomains) == 1 {
			// Auto-select if only one domain
			domain = ownedDomains[0].Name
			fmt.Printf("Editing domain: %s\n", domain)
		} else {
			// Present interactive selection
			domainNames := make([]string, len(ownedDomains))
			for i, d := range ownedDomains {
				domainNames[i] = d.Name
			}

			err = survey.AskOne(&survey.Select{
				Message: "Select domain to edit:",
				Options: domainNames,
			}, &domain)
			if err != nil {
				return err
			}
		}
	}

	// Determine format
	var format inwx.ExportFormat
	var importFormat inwx.ImportFormat
	var ext string
	switch c.String("format") {
	case "json":
		format = inwx.ExportJSON
		importFormat = inwx.ImportJSON
		ext = "json"
	case "zonefile":
		format = inwx.ExportZonefileFormat
		importFormat = inwx.ImportZonefileFormat
		ext = "zone"
	default:
		return fmt.Errorf("unsupported format: %s", c.String("format"))
	}

	// Export current records to temporary file
	backupStore, err := backup.NewStore()
	if err != nil {
		return err
	}
	dns := client.DNS(inwx.WithDomain(domain), inwx.WithBackupStore(backupStore))
	data, err := dns.ExportRecords(ctx, format)
	if err != nil {
		return fmt.Errorf("failed to export records: %w", err)
	}

	// Create temporary file
	tmpFile, err := os.CreateTemp("", fmt.Sprintf("inwx-dns-%s-*.%s", domain, ext))
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	if _, err := tmpFile.Write(data); err != nil {
		tmpFile.Close()
		return fmt.Errorf("failed to write to temporary file: %w", err)
	}
	tmpFile.Close()

	// Get original file modification time
	origStat, err := os.Stat(tmpPath)
	if err != nil {
		return fmt.Errorf("failed to stat temporary file: %w", err)
	}
	origModTime := origStat.ModTime()

	// Determine editor
	editor := os.Getenv("EDITOR")
	if editor == "" {
		// OS-specific fallbacks
		if _, err := os.Stat("/usr/bin/nano"); err == nil {
			editor = "nano"
		} else if _, err := os.Stat("/usr/bin/vi"); err == nil {
			editor = "vi"
		} else if _, err := os.Stat("/usr/bin/vim"); err == nil {
			editor = "vim"
		} else {
			editor = "vi" // Last resort
		}
	}

	// Open editor
	fmt.Printf("Opening %s in editor...\n", tmpPath)
	cmd := exec.Command(editor, tmpPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("editor failed: %w", err)
	}

	// Check if file was modified
	newStat, err := os.Stat(tmpPath)
	if err != nil {
		return fmt.Errorf("failed to stat temporary file after editing: %w", err)
	}

	if newStat.ModTime().Equal(origModTime) {
		fmt.Println("No changes made, exiting")
		return nil
	}

	// Read modified file
	modifiedData, err := os.ReadFile(tmpPath)
	if err != nil {
		return fmt.Errorf("failed to read modified file: %w", err)
	}

	// Parse records from modified file
	var newRecords []inwx.DNSRecord
	switch importFormat {
	case inwx.ImportJSON:
		newRecords, err = importJSON(modifiedData)
	case inwx.ImportZonefileFormat:
		newRecords, err = inwx.ImportZonefile(modifiedData, domain)
	}
	if err != nil {
		return fmt.Errorf("failed to parse modified file: %w", err)
	}

	// Get current records
	currentRecords, err := dns.ListRecords(ctx)
	if err != nil {
		return fmt.Errorf("failed to list current records: %w", err)
	}

	// Filter out system-managed records (SOA, NS at root)
	currentRecords = filterEditableRecords(currentRecords)
	newRecords = filterEditableRecords(newRecords)

	// Compare and show differences
	toAdd, toRemove := compareRecords(currentRecords, newRecords)

	if len(toAdd) == 0 && len(toRemove) == 0 {
		fmt.Println("No changes detected")
		return nil
	}

	// Display changes
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("Changes to be applied:")
	fmt.Println(strings.Repeat("=", 60))

	if len(toRemove) > 0 {
		fmt.Printf("\n🗑️  Records to be REMOVED (%d):\n", len(toRemove))
		for _, rec := range toRemove {
			displayName := rec.Name
			if displayName == "" {
				displayName = "@"
			}
			if rec.Type == "MX" || rec.Type == "SRV" {
				fmt.Printf("  - %s %s %d %s (TTL: %d)\n", rec.Type, displayName, rec.Prio, rec.Content, rec.TTL)
			} else {
				fmt.Printf("  - %s %s %s (TTL: %d)\n", rec.Type, displayName, rec.Content, rec.TTL)
			}
		}
	}

	if len(toAdd) > 0 {
		fmt.Printf("\n➕ Records to be ADDED (%d):\n", len(toAdd))
		for _, rec := range toAdd {
			displayName := rec.Name
			if displayName == "" {
				displayName = "@"
			}
			if rec.Type == "MX" || rec.Type == "SRV" {
				fmt.Printf("  + %s %s %d %s (TTL: %d)\n", rec.Type, displayName, rec.Prio, rec.Content, rec.TTL)
			} else {
				fmt.Printf("  + %s %s %s (TTL: %d)\n", rec.Type, displayName, rec.Content, rec.TTL)
			}
		}
	}

	fmt.Println(strings.Repeat("=", 60))

	// Ask for confirmation
	var confirm bool
	err = survey.AskOne(&survey.Confirm{
		Message: "Apply these changes?",
		Default: false,
	}, &confirm)
	if err != nil {
		return err
	}

	if !confirm {
		fmt.Println("❌ Changes cancelled")
		return nil
	}

	// Apply changes
	// Remove records
	for _, rec := range toRemove {
		if err := dns.DeleteRecord(ctx, rec.ID); err != nil {
			log.Error().Err(err).Int("id", rec.ID).Msg("Failed to delete record")
			return fmt.Errorf("failed to delete record ID %d: %w", rec.ID, err)
		}
	}

	// Add records
	for _, rec := range toAdd {
		if _, err := dns.CreateRecord(ctx, rec); err != nil {
			log.Error().Err(err).Str("record", fmt.Sprintf("%s %s", rec.Type, rec.Name)).Msg("Failed to create record")
			return fmt.Errorf("failed to create record %s %s: %w", rec.Type, rec.Name, err)
		}
	}

	fmt.Printf("\n✅ Successfully applied changes: removed %d, added %d records\n", len(toRemove), len(toAdd))
	return nil
}

// filterEditableRecords filters out system-managed records that shouldn't be edited
func filterEditableRecords(records []inwx.DNSRecord) []inwx.DNSRecord {
	var filtered []inwx.DNSRecord
	for _, rec := range records {
		// Skip SOA records - these are system-managed
		if rec.Type == "SOA" {
			continue
		}
		// Skip NS records at root - these are usually system-managed
		if rec.Type == "NS" && (rec.Name == "" || rec.Name == "@") {
			continue
		}
		filtered = append(filtered, rec)
	}
	return filtered
}

// normalizeRecordName normalizes record names for comparison (@ and "" both mean root)
func normalizeRecordName(name string) string {
	if name == "@" || name == "" {
		return "@"
	}
	return name
}

// normalizeRecordContent normalizes record content for comparison
func normalizeRecordContent(recordType, content string) string {
	// Normalize content (trim spaces, lowercase for case-insensitive comparison)
	content = strings.TrimSpace(content)

	// For some record types, order within the content doesn't matter
	// For now, just do basic normalization
	return content
}

// recordKey generates a unique key for a DNS record for comparison
func recordKey(rec inwx.DNSRecord) string {
	name := normalizeRecordName(rec.Name)
	content := normalizeRecordContent(rec.Type, rec.Content)

	// For MX and SRV records, include priority in the key since it's part of the record identity
	if rec.Type == "MX" || rec.Type == "SRV" {
		return fmt.Sprintf("%s|%s|%d|%s|%d", rec.Type, name, rec.Prio, content, rec.TTL)
	}

	// Include TTL in the comparison since changing TTL is a meaningful change
	return fmt.Sprintf("%s|%s|%s|%d", rec.Type, name, content, rec.TTL)
}

// compareRecords compares current and new records to determine what to add and remove
func compareRecords(current, new []inwx.DNSRecord) (toAdd, toRemove []inwx.DNSRecord) {
	// Create map of new records
	newMap := make(map[string]inwx.DNSRecord)
	for _, rec := range new {
		key := recordKey(rec)
		newMap[key] = rec
	}

	// Create map of current records
	currentMap := make(map[string]inwx.DNSRecord)
	for _, rec := range current {
		key := recordKey(rec)
		currentMap[key] = rec
	}

	// Find records to remove (in current but not in new)
	for _, rec := range current {
		key := recordKey(rec)
		if _, exists := newMap[key]; !exists {
			toRemove = append(toRemove, rec)
		}
	}

	// Find records to add (in new but not in current)
	for _, rec := range new {
		key := recordKey(rec)
		if _, exists := currentMap[key]; !exists {
			toAdd = append(toAdd, rec)
		}
	}

	return toAdd, toRemove
}

func validateDNSRecords(c *cli.Context) error {
	domains := parseCommaSeparatedValues(c.StringSlice("domain"))
	minSeverity := strings.ToLower(c.String("severity"))

	// Validate severity level
	validSeverities := map[string]int{
		"error":   3,
		"warning": 2,
		"info":    1,
	}

	minLevel, ok := validSeverities[minSeverity]
	if !ok {
		return fmt.Errorf("invalid severity level: %s (must be error, warning, or info)", minSeverity)
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

	// If no domains specified, get all domains
	if len(domains) == 0 {
		domainService := client.Domain()
		domainList, err := domainService.List(ctx)
		if err != nil {
			return fmt.Errorf("failed to list domains: %w", err)
		}
		for _, d := range domainList {
			domains = append(domains, d.Name)
		}
	}

	if len(domains) == 0 {
		return fmt.Errorf("no domains found to validate")
	}

	dns := client.DNS()
	totalErrors := 0
	totalWarnings := 0
	totalInfo := 0

	for _, domain := range domains {
		log.Info().Msgf("Validating %s...", domain)
		fmt.Printf("\nValidating %s...\n", domain)

		result, err := dns.ValidateDomain(ctx, domain)
		if err != nil {
			log.Error().Err(err).Str("domain", domain).Msg("Validation failed")
			fmt.Printf("  ✗ Failed: %v\n", err)
			continue
		}

		if result.Summary.Total == 0 {
			fmt.Printf("  ✓ No issues found\n")
			continue
		}

		// Display issues
		for _, issue := range result.Issues {
			issueLevel := validSeverities[issue.Severity]
			if issueLevel < minLevel {
				continue // Skip issues below minimum severity
			}

			var icon string
			switch issue.Severity {
			case "error":
				icon = "✗"
				totalErrors++
			case "warning":
				icon = "⚠"
				totalWarnings++
			case "info":
				icon = "ℹ"
				totalInfo++
			}

			fmt.Printf("\n%s %s: %s", icon, strings.ToUpper(issue.Severity), issue.Message)
			if issue.RecordID > 0 {
				fmt.Printf(" (ID: %d)", issue.RecordID)
			}
			fmt.Println()

			if issue.Suggestion != "" {
				fmt.Printf("  → %s\n", issue.Suggestion)
			}
		}

		// Show summary for this domain
		if result.Summary.Errors > 0 || result.Summary.Warnings > 0 || result.Summary.Info > 0 {
			fmt.Printf("\nSummary for %s:", domain)
			if result.Summary.Errors > 0 {
				fmt.Printf(" %d errors", result.Summary.Errors)
			}
			if result.Summary.Warnings > 0 {
				fmt.Printf(" %d warnings", result.Summary.Warnings)
			}
			if result.Summary.Info > 0 {
				fmt.Printf(" %d info", result.Summary.Info)
			}
			fmt.Println()
		}
	}

	// Overall summary
	if len(domains) > 1 {
		fmt.Println("\n" + strings.Repeat("=", 60))
		fmt.Println("Overall Summary:")
		fmt.Printf("  Domains validated: %d\n", len(domains))
		if totalErrors > 0 {
			fmt.Printf("  Total errors:      %d\n", totalErrors)
		}
		if totalWarnings > 0 {
			fmt.Printf("  Total warnings:    %d\n", totalWarnings)
		}
		if totalInfo > 0 {
			fmt.Printf("  Total info:        %d\n", totalInfo)
		}
	}

	if totalErrors > 0 {
		return fmt.Errorf("validation found %d error(s)", totalErrors)
	}

	return nil
}

func verifyDNSRecords(c *cli.Context) error {
	// Parse flags and positional arguments
	domains := parseCommaSeparatedValues(c.StringSlice("domain"))
	names := parseCommaSeparatedValues(c.StringSlice("name"))
	types := parseCommaSeparatedValues(c.StringSlice("type"))
	positionalArgs := c.Args().Slice()
	waitDuration := c.Duration("wait")

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

	dns := client.DNS()

	// Determine what to verify
	var targetDomains []string
	var targetHosts []string

	if len(positionalArgs) > 0 {
		// Use positional arguments - classify them as domains or hosts
		targetDomains, targetHosts, err = classifyAndDeduplicateTargets(ctx, client, positionalArgs)
		if err != nil {
			return err
		}
	} else if len(domains) > 0 {
		// Use flag-specified domains
		targetDomains = domains
	} else {
		return fmt.Errorf("no domains or hosts specified - provide domain(s) as arguments or use --domain flag")
	}

	if waitDuration > 0 {
		fmt.Printf("Waiting for propagation (timeout: %v)...\n\n", waitDuration)
		// TODO: Implement wait with new verification
		return fmt.Errorf("--wait not yet implemented with improved verification")
	}

	// Track overall results
	allMatch := true
	totalRecordGroups := 0
	totalMatches := 0
	totalPartial := 0
	totalMismatches := 0
	totalMissing := 0

	// Verify all target domains
	for _, domain := range targetDomains {
		var recordType string
		if len(types) > 0 {
			recordType = types[0] // Use first type if specified
		}

		var name string
		if len(names) > 0 {
			name = names[0] // Use first name if specified
		}

		fmt.Printf("\n%s\n", strings.Repeat("=", 80))
		fmt.Printf("Verifying domain: %s", domain)
		if name != "" {
			fmt.Printf(" (hostname: %s)", name)
		}
		if recordType != "" {
			fmt.Printf(" (type: %s)", recordType)
		}
		fmt.Println()
		fmt.Printf("%s\n\n", strings.Repeat("=", 80))

		result, err := dns.VerifyDomainRecords(ctx, domain, name, recordType)
		if err != nil {
			log.Error().Err(err).Str("domain", domain).Msg("Failed to verify domain")
			allMatch = false
			continue
		}

		// Display results grouped by record
		for _, rec := range result.Records {
			displayRecordVerification(rec)
		}

		// Accumulate summary
		totalRecordGroups += result.Summary.Total
		totalMatches += result.Summary.Match
		totalPartial += result.Summary.Partial
		totalMismatches += result.Summary.Mismatch
		totalMissing += result.Summary.Missing

		if result.Summary.Match != result.Summary.Total {
			allMatch = false
		}
	}

	// Verify all target hosts (specific hostname.domain combinations)
	for _, host := range targetHosts {
		// Parse hostname and domain from full host using proper domain inference
		domain, hostname, err := utils.InferDomainAndName(client, ctx, host)
		if err != nil {
			log.Warn().Str("host", host).Err(err).Msg("Failed to infer domain and name, skipping")
			continue
		}

		var recordType string
		if len(types) > 0 {
			recordType = types[0] // Use first type if specified
		}

		fmt.Printf("\n%s\n", strings.Repeat("=", 80))
		fmt.Printf("Verifying host: %s", host)
		if recordType != "" {
			fmt.Printf(" (type: %s)", recordType)
		}
		fmt.Println()
		fmt.Printf("%s\n\n", strings.Repeat("=", 80))

		result, err := dns.VerifyDomainRecords(ctx, domain, hostname, recordType)
		if err != nil {
			log.Error().Err(err).Str("host", host).Msg("Failed to verify host")
			allMatch = false
			continue
		}

		// Display results grouped by record
		for _, rec := range result.Records {
			displayRecordVerification(rec)
		}

		// Accumulate summary
		totalRecordGroups += result.Summary.Total
		totalMatches += result.Summary.Match
		totalPartial += result.Summary.Partial
		totalMismatches += result.Summary.Mismatch
		totalMissing += result.Summary.Missing

		if result.Summary.Match != result.Summary.Total {
			allMatch = false
		}
	}

	// Display overall summary
	fmt.Printf("\n%s\n", strings.Repeat("=", 80))
	fmt.Printf("Overall Summary: ")

	summaryParts := []string{}
	if totalMatches > 0 {
		summaryParts = append(summaryParts, fmt.Sprintf("%d match", totalMatches))
	}
	if totalPartial > 0 {
		summaryParts = append(summaryParts, fmt.Sprintf("%d partial", totalPartial))
	}
	if totalMismatches > 0 {
		summaryParts = append(summaryParts, fmt.Sprintf("%d mismatch", totalMismatches))
	}
	if totalMissing > 0 {
		summaryParts = append(summaryParts, fmt.Sprintf("%d missing", totalMissing))
	}

	fmt.Printf("%s (total: %d record groups)\n", strings.Join(summaryParts, ", "), totalRecordGroups)

	// Overall status icon
	var overallIcon string
	if allMatch && totalMatches == totalRecordGroups {
		overallIcon = "✓"
		fmt.Printf("\n%s All records verified successfully!\n", overallIcon)
	} else if totalMatches > 0 {
		overallIcon = "⚠"
		fmt.Printf("\n%s Partial verification - some records don't match\n", overallIcon)
	} else {
		overallIcon = "✗"
		fmt.Printf("\n%s Verification failed\n", overallIcon)
	}

	if !allMatch {
		return fmt.Errorf("verification incomplete")
	}

	return nil
}

// isatty checks if the given file descriptor is a terminal
func isatty(f *os.File) bool {
	// Check if file descriptor is a terminal using stat
	fileInfo, err := f.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

// displayRecordCreated shows a nicely formatted success message with record details
func displayRecordCreated(record *inwx.DNSRecord) {
	isTTY := isatty(os.Stdout)

	// Colors
	green := ""
	cyan := ""
	yellow := ""
	bold := ""
	reset := ""

	if isTTY {
		green = "\033[32m"
		cyan = "\033[36m"
		yellow = "\033[33m"
		bold = "\033[1m"
		reset = "\033[0m"
	}

	// Success header
	fmt.Printf("\n%s%s✓ DNS Record Created Successfully%s\n\n", green, bold, reset)

	// Record details
	fmt.Printf("  %sID:%s         %d\n", cyan, reset, record.ID)
	fmt.Printf("  %sType:%s       %s\n", cyan, reset, record.Type)
	fmt.Printf("  %sDomain:%s     %s\n", cyan, reset, record.Domain)
	fmt.Printf("  %sName:%s       %s\n", cyan, reset, record.Name)
	fmt.Printf("  %sContent:%s    %s\n", cyan, reset, record.Content)
	fmt.Printf("  %sTTL:%s        %s%d%s seconds\n", cyan, reset, yellow, record.TTL, reset)

	if record.Prio > 0 {
		fmt.Printf("  %sPriority:%s   %d\n", cyan, reset, record.Prio)
	}

	// Full hostname
	fullName := record.Name
	if fullName == "" || fullName == "@" {
		fullName = record.Domain
	} else {
		fullName = record.Name + "." + record.Domain
	}
	fmt.Printf("\n  %sFull name:%s  %s%s%s\n", cyan, reset, bold, fullName, reset)
}

// displayRecordUpdated shows a nicely formatted success message for updated record
func displayRecordUpdated(record *inwx.DNSRecord) {
	isTTY := isatty(os.Stdout)

	// Colors
	green := ""
	cyan := ""
	yellow := ""
	bold := ""
	reset := ""

	if isTTY {
		green = "\033[32m"
		cyan = "\033[36m"
		yellow = "\033[33m"
		bold = "\033[1m"
		reset = "\033[0m"
	}

	// Success header
	fmt.Printf("\n%s%s✓ DNS Record Updated Successfully%s\n\n", green, bold, reset)

	// Record details
	fmt.Printf("  %sID:%s         %d\n", cyan, reset, record.ID)
	fmt.Printf("  %sType:%s       %s\n", cyan, reset, record.Type)
	fmt.Printf("  %sDomain:%s     %s\n", cyan, reset, record.Domain)
	fmt.Printf("  %sName:%s       %s\n", cyan, reset, record.Name)
	fmt.Printf("  %sContent:%s    %s\n", cyan, reset, record.Content)
	fmt.Printf("  %sTTL:%s        %s%d%s seconds\n", cyan, reset, yellow, record.TTL, reset)

	if record.Prio > 0 {
		fmt.Printf("  %sPriority:%s   %d\n", cyan, reset, record.Prio)
	}

	// Full hostname
	fullName := record.Name
	if fullName == "" || fullName == "@" {
		fullName = record.Domain
	} else {
		fullName = record.Name + "." + record.Domain
	}
	fmt.Printf("\n  %sFull name:%s  %s%s%s\n", cyan, reset, bold, fullName, reset)
}

// waitForDNSPropagation waits for DNS record to propagate to authoritative nameservers
func waitForDNSPropagation(ctx context.Context, dns *inwx.DNSService, domain, name, recordType string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	checkInterval := 3 * time.Second
	attempt := 0

	for {
		attempt++

		// Check if we've exceeded the timeout
		if time.Now().After(deadline) {
			return fmt.Errorf("timeout after %v", timeout)
		}

		// Verify the record
		result, err := dns.VerifyDomainRecords(ctx, domain, name, recordType)
		if err != nil {
			fmt.Printf("  [Attempt %d] Checking... (error: %v)\n", attempt, err)
			time.Sleep(checkInterval)
			continue
		}

		// Check if any records were found
		if len(result.Records) == 0 {
			fmt.Printf("  [Attempt %d] Checking... (no records found in INWX yet)\n", attempt)
			time.Sleep(checkInterval)
			continue
		}

		// Check only authoritative nameservers (ignore public DNS)
		allAuthMatch := true
		authCount := 0
		authMatchCount := 0

		for _, rec := range result.Records {
			for _, ns := range rec.Nameservers {
				if ns.Type == "authoritative" {
					authCount++
					if ns.Status == "match" {
						authMatchCount++
					} else {
						allAuthMatch = false
					}
				}
			}
		}

		// If we have no authoritative nameservers, something is wrong
		if authCount == 0 {
			fmt.Printf("  [Attempt %d] Checking... (no authoritative nameservers found)\n", attempt)
			time.Sleep(checkInterval)
			continue
		}

		// Success: all authoritative nameservers have the record
		if allAuthMatch && authMatchCount == authCount {
			fmt.Printf("  [Attempt %d] All %d authoritative nameserver(s) updated\n", attempt, authCount)
			return nil
		}

		// Still propagating
		fmt.Printf("  [Attempt %d] Propagating... (%d/%d nameserver(s) updated)\n", attempt, authMatchCount, authCount)

		// Wait before next check
		select {
		case <-time.After(checkInterval):
			continue
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func displayRecordVerification(rec inwx.RecordVerification) {
	// Status icon
	var statusIcon string
	switch rec.Status {
	case "match":
		statusIcon = "✓"
	case "partial":
		statusIcon = "⚠"
	case "mismatch":
		statusIcon = "✗"
	case "missing":
		statusIcon = "○"
	}

	// Header
	fmt.Printf("%s %s %s\n", statusIcon, rec.Hostname, rec.Type)

	// Expected values (truncate if too long)
	fmt.Printf("  Expected: ")
	if len(rec.Expected) <= 3 {
		fmt.Printf("%s\n", strings.Join(rec.Expected, ", "))
	} else {
		fmt.Printf("%s, ... (%d total)\n", strings.Join(rec.Expected[:3], ", "), len(rec.Expected))
	}

	// Nameserver summary
	authMatch := 0
	authTotal := 0
	publicMatch := 0
	publicTotal := 0

	for _, ns := range rec.Nameservers {
		if ns.Type == "authoritative" {
			authTotal++
			if ns.Status == "match" {
				authMatch++
			}
		} else {
			publicTotal++
			if ns.Status == "match" {
				publicMatch++
			}
		}
	}

	fmt.Printf("  Authoritative: %d/%d match", authMatch, authTotal)
	if authMatch < authTotal {
		fmt.Printf(" (")
		for i, ns := range rec.Nameservers {
			if ns.Type == "authoritative" && ns.Status != "match" {
				if i > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%s: %s", ns.Server, ns.Status)
			}
		}
		fmt.Printf(")")
	}
	fmt.Println()

	fmt.Printf("  Public DNS:    %d/%d match", publicMatch, publicTotal)
	if publicMatch < publicTotal {
		fmt.Printf(" (")
		first := true
		for _, ns := range rec.Nameservers {
			if ns.Type == "public" && ns.Status != "match" {
				if !first {
					fmt.Printf(", ")
				}
				first = false
				if ns.Error != "" {
					fmt.Printf("%s: %s", ns.Server, ns.Error)
				} else {
					fmt.Printf("%s: %s", ns.Server, ns.Status)
				}
			}
		}
		fmt.Printf(")")
	}
	fmt.Println()

	// Show details if there are mismatches
	if rec.Status == "mismatch" || rec.Status == "partial" {
		for _, ns := range rec.Nameservers {
			if ns.Status == "mismatch" && len(ns.Response) > 0 {
				fmt.Printf("    %s returned: %s\n", ns.Server, strings.Join(ns.Response, ", "))
				break // Just show one example
			}
		}
	}

	fmt.Println()
}
