# inwx-cli

[![Build Status](https://img.shields.io/github/actions/workflow/status/nmeilick/inwx-cli/release.yml?branch=main)](https://github.com/nmeilick/inwx-cli/actions/workflows/release.yml)
[![License](https://img.shields.io/github/license/nmeilick/inwx-cli)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/nmeilick/inwx-cli)](https://goreportcard.com/report/github.com/nmeilick/inwx-cli)

**inwx-cli** is a powerful command-line interface for managing DNS records and domains through the INWX DomRobot API. It
provides an intuitive way to create, update, delete, and backup DNS records with built-in safety features and multiple
output formats.

## Features

*   **Complete DNS Management:** Create, update, delete, and list DNS records with full type support (A, AAAA, CNAME, MX, TXT, NS, etc.).
*   **Interactive Mode:** Guided DNS record creation with prompts, validation, and preview.
*   **DNS Validation:** Analyze DNS configurations for common issues (orphaned CNAMEs, missing targets, RFC violations).
*   **DNS Verification:** Verify DNS propagation across multiple resolvers with real-time status updates.
*   **Backup & Recovery:** Automatic backup of all DNS operations with rollback capability.
*   **Batch Operations:** Update multiple records simultaneously.
*   **Multiple Output Formats:** Table, JSON, YAML, and CSV output formats.
*   **Bulk Operations:** Filter and operate on multiple records using wildcards and patterns.
*   **Import/Export:** Backup domains as JSON or zonefile format.
*   **Safety Features:** Dry-run mode, confirmation prompts, and operation limits.
*   **Flexible Configuration:** TOML configuration files with XDG directory support and environment variables.
*   **Cross-Platform:** Binaries available for Linux, Windows, and macOS.

## Installation

### Pre-compiled Binaries

Download the latest pre-compiled binaries for your platform from the [Releases page](https://github.com/nmeilick/inwx-cli/releases).

Extract the archive and place the `inwx` binary in a directory included in your system's `PATH`.

### Configuration

Create a configuration file at `~/.config/inwx/inwx.toml`:

```toml
[api]
username = "your_username"
password = "your_password"
endpoint = "https://api.domrobot.com/jsonrpc/"
test_mode = false

[output]
format = "table"
colors = true

[logging]
level = "warn"
colors = true
```

Instead of, or in addition to, the config file, you can also use environment variables or command-line flags.

**Available environment variables:**
- `INWX_USERNAME` - API username
- `INWX_PASSWORD` - API password
- `INWX_ENDPOINT` - API endpoint URL
- `INWX_CONFIG` - Path to config file
- `INWX_TIMEOUT` - Request timeout in seconds (default: 30)

## Usage

### DNS Record Management

#### Creating Records

**Interactive Mode:**

The easiest way to create DNS records is using interactive mode, which guides you through the process with prompts and validation:

```bash
# Start interactive mode (auto-starts when no arguments provided)
inwx dns create

# Or explicitly enable interactive mode
inwx dns create --interactive
```

Interactive mode will:
- Prompt for hostname (e.g., www.example.com or example.com for root record)
- Let you select the record type from a list (A, AAAA, CNAME, MX, TXT, NS, SRV, CAA)
- Request content with type-specific help text and validation
- Ask for TTL (default: 3600 seconds)
- Request priority for MX and SRV records
- Show a preview of the record before creation
- Ask for confirmation before creating
- Optionally wait for DNS propagation verification

**Command-Line Mode:**

You can also create records directly using flags:

```bash
# Create an A record
inwx dns create -d example.com -t A -n www -c 192.168.1.100

# Create an MX record with priority
inwx dns create -d example.com -t MX -n @ -c mail.example.com --prio 10

# Create a CNAME record
inwx dns create -d example.com -t CNAME -n blog -c www.example.com

# Create a TXT record for SPF
inwx dns create -d example.com -t TXT -n @ -c "v=spf1 include:_spf.google.com ~all"

# Preview changes before creating (dry-run mode)
inwx dns create -d example.com -t A -n test -c 192.168.1.100 --dry-run
```

#### Listing Records

```bash
# List records across all domains
inwx dns list

# List all records for a domain
inwx dns list -d example.com

# List specific record types
inwx dns list -d example.com -t A,AAAA

# List records by name pattern
inwx dns list -d example.com -n "mail*" --wildcard

# Output as JSON
inwx dns list -d example.com -o json
```

#### Updating Records

```bash
# Update a record by ID
inwx dns update --id 12345 -c 192.168.1.200

# Update TTL and content
inwx dns update --id 12345 -c 192.168.1.200 --ttl 7200

# Preview changes before updating (dry-run mode)
inwx dns update --id 12345 -c 192.168.1.200 --dry-run

# Batch update multiple records at once (more efficient)
inwx dns update --ids 12345,12346,12347 --ttl 3600
```

#### Deleting Records

```bash
# Delete by record ID
inwx dns delete --id 12345

# Delete by FQDN (domain and name inferred automatically)
inwx dns delete www.example.com

# Delete multiple records by pattern
inwx dns delete -d example.com -n "test*" --wildcard

# Delete with confirmation (dry-run first)
inwx dns delete -d example.com -t TXT --dry-run
inwx dns delete -d example.com -t TXT --yes
```

### Backup and Export

#### Export Domain Records

```bash
# Export domain as JSON
inwx dns export -d example.com -f json -o example.com.json

# Export as zonefile
inwx dns export -d example.com -f zonefile -o example.com.zone

# Export all domains to directory
inwx dns export --output-dir ./backups -f json
```

#### Import Records *(Work in Progress)*

```bash
# Import from JSON file
inwx dns import -f example.com.json -d example.com --format json

# Import with sync (delete records not in file)
inwx dns import -f example.com.json -d example.com --delete --dry-run
```

### Backup Management

```bash
# List recent backup entries
inwx backup list

# List backups for specific domain
inwx backup list -d example.com

# Revert a DNS operation
inwx backup revert <backup-id>

# Clean up old backups
inwx backup purge --older-than 30d
```

### Domain Management

```bash
# List all domains
inwx domain list

# Show account information
inwx account info
```

### DNS Validation

Validate your DNS configuration for common issues and best practices:

```bash
# Validate all domains
inwx dns validate

# Validate specific domain
inwx dns validate -d example.com

# Show only errors and warnings (hide info messages)
inwx dns validate -d example.com --severity warning

# Validation checks include:
# - Orphaned CNAME records pointing to non-existent targets
# - Missing MX/NS/SRV record targets
# - CNAME conflicts (RFC 1034 violations)
# - Common best practices (root record, www subdomain)
# - Unusual TTL values
```

### DNS Verification

Verify that your DNS changes have propagated to public DNS servers:

```bash
# Verify all records for a domain
inwx dns verify -d example.com

# Verify specific hostnames
inwx dns verify www.example.com mail.example.com

# Wait for propagation (up to timeout)
inwx dns verify -d example.com --wait --timeout 300

# Verify specific record types
inwx dns verify -d example.com -t A,AAAA

# The command checks:
# - Record existence on authoritative nameservers
# - Propagation to multiple public DNS resolvers (Google, Cloudflare, Quad9)
# - Consistency across resolvers
# - Time to propagation
```

### Advanced Usage

#### Bulk Operations with Filters

```bash
# Delete all test records
inwx dns delete -n "test*" --wildcard --max 50

# List records with specific content
inwx dns list -c "192.168.1.*" --wildcard

# List records with specific TTL
inwx dns list -d example.com --ttl 300

# List records with specific priority
inwx dns list -d example.com --prio 10

# Update TTL for multiple A records in one batch operation
inwx dns list -d example.com -t A -o json | jq -r '.[].id | @csv' | tr -d '"' | xargs -I {} inwx dns update --ids {} --ttl 3600
```

#### Different Output Formats

```bash
# Table format (default)
inwx dns list -d example.com

# JSON format
inwx dns list -d example.com -o json

# YAML format
inwx dns list -d example.com -o yaml

# CSV format for spreadsheets
inwx dns list -d example.com -o csv > records.csv
```

#### Using Test Environment

```bash
# Use INWX test environment
inwx --test dns list

# Or set in config file
[api]
test_mode = true
```

## Building from Source

**Prerequisites:**

*   Go 1.21 or later
*   Make (optional, for using the Makefile)

**Steps:**

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/nmeilick/inwx-cli.git
    cd inwx-cli
    ```

2.  **Build using Make:**
    ```bash
    make build
    # Binary will be created as 'build/inwx'
    ```

3.  **Build using Go:**
    ```bash
    go build -o inwx ./cmd/inwx
    ```

4.  **Install dependencies:**
    ```bash
    make deps
    ```

5.  **Build distribution binaries:**
    ```bash
    make dist
    ```

## Configuration File Locations

inwx-cli looks for configuration files in the following order:

1. `--config` flag path
2. `$XDG_CONFIG_HOME/inwx/inwx.toml` (usually `~/.config/inwx/inwx.toml`)
3. `./inwx.toml` (current directory)

## Safety Features

*   **Automatic Backups:** All DNS operations are automatically backed up and can be reverted.
*   **Dry-Run Mode:** Use `--dry-run` on create, update, and delete operations to preview changes before applying them.
*   **DNS Validation:** Proactively detect configuration issues before they cause problems.
*   **DNS Verification:** Confirm propagation of changes to public DNS servers.
*   **Confirmation Prompts:** Interactive confirmation for destructive operations.
*   **Operation Limits:** Built-in limits prevent accidental bulk deletions.
*   **Detailed Logging:** Configurable logging levels

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
