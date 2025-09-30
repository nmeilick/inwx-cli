#!/usr/bin/env bash

#############################################################################
# Comprehensive Test Script for inwx-cli
# Tests all functionality against the INWX test environment
#############################################################################

set -e  # Exit on error
set -o pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
MAGENTA='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color
BOLD='\033[1m'

# Test configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BINARY="${SCRIPT_DIR}/bin/inwx"
TEST_DOMAIN="${TEST_DOMAIN:-}" # Set this to your test domain
PASSED=0
FAILED=0
SKIPPED=0

# Ensure binary exists
if [[ ! -f "$BINARY" ]]; then
    echo -e "${RED}Error: Binary not found at $BINARY${NC}"
    echo "Please run 'make build' or 'go build -o bin/inwx ./cmd/inwx' first"
    exit 1
fi

# Function to print section headers
print_section() {
    echo ""
    echo -e "${BOLD}${CYAN}═══════════════════════════════════════════════════════════${NC}"
    echo -e "${BOLD}${CYAN} $1${NC}"
    echo -e "${BOLD}${CYAN}═══════════════════════════════════════════════════════════${NC}"
    echo ""
}

# Function to print test headers
print_test() {
    echo -e "${BLUE}▶ Testing: ${BOLD}$1${NC}"
}

# Function to print success
print_success() {
    echo -e "${GREEN}✓ PASS${NC}: $1"
    ((PASSED++))
}

# Function to print failure
print_failure() {
    echo -e "${RED}✗ FAIL${NC}: $1"
    echo -e "${RED}  Error: $2${NC}"
    ((FAILED++))
}

# Function to print skip
print_skip() {
    echo -e "${YELLOW}⊘ SKIP${NC}: $1"
    ((SKIPPED++))
}

# Function to print info
print_info() {
    echo -e "${MAGENTA}ℹ${NC} $1"
}

# Function to print warning
print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

# Function to run command and capture output
run_cmd() {
    local output
    local exit_code

    if output=$("$@" 2>&1); then
        exit_code=0
    else
        exit_code=$?
    fi

    echo "$output"
    return $exit_code
}

# Function to test command success
test_cmd_success() {
    local test_name="$1"
    shift
    local output

    print_test "$test_name"

    if output=$(run_cmd "$@"); then
        print_success "$test_name"
        return 0
    else
        print_failure "$test_name" "Command failed with exit code $?"
        echo "$output"
        return 1
    fi
}

# Function to test command failure (should fail)
test_cmd_failure() {
    local test_name="$1"
    local expected_error="$2"
    shift 2
    local output

    print_test "$test_name"

    if output=$(run_cmd "$@" 2>&1); then
        print_failure "$test_name" "Command should have failed but succeeded"
        return 1
    else
        if [[ -z "$expected_error" ]] || echo "$output" | grep -q "$expected_error"; then
            print_success "$test_name"
            return 0
        else
            print_failure "$test_name" "Did not find expected error: $expected_error"
            echo "$output"
            return 1
        fi
    fi
}

# Function to check if command output contains string
test_output_contains() {
    local test_name="$1"
    local expected="$2"
    shift 2
    local output

    print_test "$test_name"

    if output=$(run_cmd "$@" 2>&1); then
        if echo "$output" | grep -q "$expected"; then
            print_success "$test_name"
            return 0
        else
            print_failure "$test_name" "Output does not contain: $expected"
            echo "$output"
            return 1
        fi
    else
        print_failure "$test_name" "Command failed"
        echo "$output"
        return 1
    fi
}

# Check for credentials
check_credentials() {
    print_section "Checking Prerequisites"

    if [[ -z "${INWX_USERNAME:-}" ]]; then
        print_warning "INWX_USERNAME environment variable not set"
        print_info "Please set credentials:"
        echo "  export INWX_USERNAME='your_username'"
        echo "  export INWX_PASSWORD='your_password'"
        echo "  export TEST_DOMAIN='your-test-domain.com'"
        return 1
    fi

    if [[ -z "${INWX_PASSWORD:-}" ]]; then
        print_warning "INWX_PASSWORD environment variable not set"
        return 1
    fi

    print_success "Credentials configured"

    if [[ -z "$TEST_DOMAIN" ]]; then
        print_warning "TEST_DOMAIN not set - some tests will be skipped"
        print_info "Set with: export TEST_DOMAIN='your-test-domain.com'"
        return 0
    fi

    print_success "Test domain configured: $TEST_DOMAIN"
    return 0
}

#############################################################################
# Test: Basic Commands
#############################################################################
test_basic_commands() {
    print_section "Basic Commands"

    test_cmd_success "Version command" "$BINARY" --version
    test_cmd_success "Help command" "$BINARY" --help
    test_cmd_success "DNS help" "$BINARY" dns --help
}

#############################################################################
# Test: Account Operations
#############################################################################
test_account_operations() {
    print_section "Account Operations"

    test_output_contains "Account info (test env)" "Username" \
        "$BINARY" --test account info

    test_output_contains "Account info (JSON format)" "username" \
        "$BINARY" --test account info -o json
}

#############################################################################
# Test: Domain Operations
#############################################################################
test_domain_operations() {
    print_section "Domain Operations"

    test_cmd_success "List domains (test env)" \
        "$BINARY" --test domain list

    test_output_contains "List domains (JSON format)" "name" \
        "$BINARY" --test domain list -o json

    test_output_contains "List domains (YAML format)" "name:" \
        "$BINARY" --test domain list -o yaml

    test_cmd_success "List domains (CSV format)" \
        "$BINARY" --test domain list -o csv
}

#############################################################################
# Test: DNS Record Listing
#############################################################################
test_dns_listing() {
    print_section "DNS Record Listing"

    if [[ -z "$TEST_DOMAIN" ]]; then
        print_skip "DNS listing tests require TEST_DOMAIN"
        return 0
    fi

    test_cmd_success "List all records for domain" \
        "$BINARY" --test dns list -d "$TEST_DOMAIN"

    test_cmd_success "List A records" \
        "$BINARY" --test dns list -d "$TEST_DOMAIN" -t A

    test_cmd_success "List multiple types" \
        "$BINARY" --test dns list -d "$TEST_DOMAIN" -t A,AAAA,CNAME

    test_cmd_success "List records (JSON)" \
        "$BINARY" --test dns list -d "$TEST_DOMAIN" -o json

    test_cmd_success "List records (YAML)" \
        "$BINARY" --test dns list -d "$TEST_DOMAIN" -o yaml

    test_cmd_success "List records (CSV)" \
        "$BINARY" --test dns list -d "$TEST_DOMAIN" -o csv
}

#############################################################################
# Test: DNS Record Creation
#############################################################################
test_dns_create() {
    print_section "DNS Record Creation"

    if [[ -z "$TEST_DOMAIN" ]]; then
        print_skip "DNS creation tests require TEST_DOMAIN"
        return 0
    fi

    local timestamp=$(date +%s)
    local test_name="test-${timestamp}"

    # Test dry-run mode
    test_cmd_success "Create A record (dry-run)" \
        "$BINARY" --test dns create -d "$TEST_DOMAIN" -t A -n "${test_name}-dry" -c "192.0.2.1" --dry-run

    # Create A record
    test_cmd_success "Create A record" \
        "$BINARY" --test dns create -d "$TEST_DOMAIN" -t A -n "$test_name" -c "192.0.2.1"

    # Create AAAA record
    test_cmd_success "Create AAAA record" \
        "$BINARY" --test dns create -d "$TEST_DOMAIN" -t AAAA -n "${test_name}-v6" -c "2001:db8::1"

    # Create TXT record
    test_cmd_success "Create TXT record" \
        "$BINARY" --test dns create -d "$TEST_DOMAIN" -t TXT -n "${test_name}-txt" -c "v=test1"

    # Create MX record
    test_cmd_success "Create MX record" \
        "$BINARY" --test dns create -d "$TEST_DOMAIN" -t MX -n "$test_name" -c "mail.example.com" --prio 10

    # Create CNAME record
    test_cmd_success "Create CNAME record" \
        "$BINARY" --test dns create -d "$TEST_DOMAIN" -t CNAME -n "${test_name}-alias" -c "$test_name.$TEST_DOMAIN"

    # Test invalid record creation
    test_cmd_failure "Create invalid A record (bad IP)" "invalid" \
        "$BINARY" --test dns create -d "$TEST_DOMAIN" -t A -n "invalid" -c "999.999.999.999" || true

    print_info "Created test records with prefix: $test_name"
    echo "$test_name" > /tmp/inwx-test-prefix.txt
}

#############################################################################
# Test: DNS Record Updates
#############################################################################
test_dns_update() {
    print_section "DNS Record Updates"

    if [[ -z "$TEST_DOMAIN" ]]; then
        print_skip "DNS update tests require TEST_DOMAIN"
        return 0
    fi

    if [[ ! -f /tmp/inwx-test-prefix.txt ]]; then
        print_skip "DNS update tests require created records"
        return 0
    fi

    local test_name=$(cat /tmp/inwx-test-prefix.txt)

    # Get record ID
    local record_json
    record_json=$("$BINARY" --test dns list -d "$TEST_DOMAIN" -n "$test_name" -t A -o json 2>/dev/null)
    local record_id=$(echo "$record_json" | grep -o '"id":[0-9]*' | head -1 | cut -d: -f2)

    if [[ -z "$record_id" ]]; then
        print_skip "Could not find test record for update"
        return 0
    fi

    print_info "Found test record ID: $record_id"

    # Test dry-run update
    test_cmd_success "Update record (dry-run)" \
        "$BINARY" --test dns update --id "$record_id" -c "192.0.2.2" --dry-run

    # Update record content
    test_cmd_success "Update record content" \
        "$BINARY" --test dns update --id "$record_id" -c "192.0.2.2"

    # Update record TTL
    test_cmd_success "Update record TTL" \
        "$BINARY" --test dns update --id "$record_id" --ttl 7200

    # Test batch update (if we have multiple records)
    local all_ids
    all_ids=$("$BINARY" --test dns list -d "$TEST_DOMAIN" -n "$test_name*" --wildcard -o json 2>/dev/null | grep -o '"id":[0-9]*' | cut -d: -f2 | head -3 | tr '\n' ',' | sed 's/,$//')

    if [[ -n "$all_ids" ]] && [[ "$all_ids" == *,* ]]; then
        test_cmd_success "Batch update TTL" \
            "$BINARY" --test dns update --ids "$all_ids" --ttl 1800
    else
        print_skip "Batch update (need multiple records)"
    fi
}

#############################################################################
# Test: DNS Record Deletion
#############################################################################
test_dns_delete() {
    print_section "DNS Record Deletion"

    if [[ -z "$TEST_DOMAIN" ]]; then
        print_skip "DNS deletion tests require TEST_DOMAIN"
        return 0
    fi

    if [[ ! -f /tmp/inwx-test-prefix.txt ]]; then
        print_skip "DNS deletion tests require created records"
        return 0
    fi

    local test_name=$(cat /tmp/inwx-test-prefix.txt)

    # Test dry-run deletion
    test_cmd_success "Delete by name (dry-run)" \
        "$BINARY" --test dns delete -d "$TEST_DOMAIN" -n "${test_name}*" --wildcard --dry-run

    # Delete all test records
    test_cmd_success "Delete test records" \
        "$BINARY" --test dns delete -d "$TEST_DOMAIN" -n "${test_name}*" --wildcard --yes

    print_info "Cleaned up test records"
    rm -f /tmp/inwx-test-prefix.txt
}

#############################################################################
# Test: DNS Export/Import
#############################################################################
test_dns_export_import() {
    print_section "DNS Export/Import"

    if [[ -z "$TEST_DOMAIN" ]]; then
        print_skip "Export/import tests require TEST_DOMAIN"
        return 0
    fi

    local export_dir="/tmp/inwx-export-test"
    mkdir -p "$export_dir"

    # Export as JSON
    test_cmd_success "Export domain (JSON)" \
        "$BINARY" --test dns export -d "$TEST_DOMAIN" -f json -o "$export_dir/${TEST_DOMAIN}.json"

    # Export as zonefile
    test_cmd_success "Export domain (zonefile)" \
        "$BINARY" --test dns export -d "$TEST_DOMAIN" -f zonefile -o "$export_dir/${TEST_DOMAIN}.zone"

    # Verify export files exist and have content
    if [[ -f "$export_dir/${TEST_DOMAIN}.json" ]]; then
        local file_size=$(stat -f%z "$export_dir/${TEST_DOMAIN}.json" 2>/dev/null || stat -c%s "$export_dir/${TEST_DOMAIN}.json" 2>/dev/null)
        if [[ $file_size -gt 10 ]]; then
            print_success "Export file created with content"
        else
            print_failure "Export file" "File too small"
        fi
    else
        print_failure "Export file" "File not created"
    fi

    # Test import dry-run
    if [[ -f "$export_dir/${TEST_DOMAIN}.json" ]]; then
        test_cmd_success "Import (dry-run)" \
            "$BINARY" --test dns import -f "$export_dir/${TEST_DOMAIN}.json" -d "$TEST_DOMAIN" --format json --dry-run
    else
        print_skip "Import test (no export file)"
    fi

    # Cleanup
    rm -rf "$export_dir"
}

#############################################################################
# Test: DNS Validation
#############################################################################
test_dns_validation() {
    print_section "DNS Validation"

    if [[ -z "$TEST_DOMAIN" ]]; then
        print_skip "Validation tests require TEST_DOMAIN"
        return 0
    fi

    # Validate domain
    test_cmd_success "Validate domain" \
        "$BINARY" --test dns validate -d "$TEST_DOMAIN"

    # Validate with severity filter
    test_cmd_success "Validate (errors only)" \
        "$BINARY" --test dns validate -d "$TEST_DOMAIN" --severity error

    test_cmd_success "Validate (warnings and up)" \
        "$BINARY" --test dns validate -d "$TEST_DOMAIN" --severity warning
}

#############################################################################
# Test: DNS Verification
#############################################################################
test_dns_verification() {
    print_section "DNS Verification"

    if [[ -z "$TEST_DOMAIN" ]]; then
        print_skip "Verification tests require TEST_DOMAIN"
        return 0
    fi

    print_info "Note: Verification queries live DNS, may fail if records not propagated"

    # Verify domain
    test_cmd_success "Verify domain records" \
        "$BINARY" --test dns verify -d "$TEST_DOMAIN" || true

    # Verify specific type
    test_cmd_success "Verify A records" \
        "$BINARY" --test dns verify -d "$TEST_DOMAIN" -t A || true

    # Verify specific hostname (if exists)
    print_test "Verify specific hostname"
    if "$BINARY" --test dns verify "${TEST_DOMAIN}" 2>/dev/null; then
        print_success "Verify specific hostname"
    else
        print_info "Verification failed or no matching records (this is expected for test domains)"
        print_skip "Verification (test environment may not have live DNS)"
    fi
}

#############################################################################
# Test: Backup Operations
#############################################################################
test_backup_operations() {
    print_section "Backup Operations"

    # List backups
    test_cmd_success "List backups" \
        "$BINARY" backup list

    test_cmd_success "List backups (JSON)" \
        "$BINARY" backup list -o json

    # Note: We don't test revert as it would modify production data
    # and we don't test purge as it would delete backup history
    print_info "Revert and purge operations not tested (would affect backup data)"
}

#############################################################################
# Test: Configuration and Options
#############################################################################
test_configuration() {
    print_section "Configuration and Options"

    # Test with different log levels
    test_cmd_success "Command with debug logging" \
        "$BINARY" --test --log-level debug account info

    test_cmd_success "Command with custom timeout" \
        "$BINARY" --test --timeout 60 account info

    # Test with custom endpoint (should still use test mode)
    test_cmd_success "Command with custom endpoint" \
        "$BINARY" --endpoint "https://api.ote.domrobot.com/jsonrpc/" account info
}

#############################################################################
# Test: Error Handling
#############################################################################
test_error_handling() {
    print_section "Error Handling"

    # Invalid domain
    test_cmd_failure "List records for non-existent domain" "not found" \
        "$BINARY" --test dns list -d "this-domain-definitely-does-not-exist-12345.com" || true

    # Invalid record ID
    test_cmd_failure "Update non-existent record" "" \
        "$BINARY" --test dns update --id 999999999 -c "192.0.2.1" || true

    # Invalid credentials (if we can test this safely)
    print_test "Invalid credentials"
    if INWX_USERNAME="invalid" INWX_PASSWORD="invalid" "$BINARY" --test account info 2>/dev/null; then
        print_failure "Invalid credentials" "Should have failed"
    else
        print_success "Invalid credentials properly rejected"
    fi
}

#############################################################################
# Test: Output Formats
#############################################################################
test_output_formats() {
    print_section "Output Formats"

    # Test all output formats for domain list
    test_output_contains "Table output (default)" "Name" \
        "$BINARY" --test domain list

    test_output_contains "JSON output" '"name"' \
        "$BINARY" --test domain list -o json

    test_output_contains "YAML output" 'name:' \
        "$BINARY" --test domain list -o yaml

    test_cmd_success "CSV output" \
        "$BINARY" --test domain list -o csv
}

#############################################################################
# Print Summary
#############################################################################
print_summary() {
    local total=$((PASSED + FAILED + SKIPPED))

    echo ""
    echo -e "${BOLD}${CYAN}═══════════════════════════════════════════════════════════${NC}"
    echo -e "${BOLD}${CYAN} TEST SUMMARY${NC}"
    echo -e "${BOLD}${CYAN}═══════════════════════════════════════════════════════════${NC}"
    echo ""
    echo -e "${BOLD}Total Tests:${NC}    $total"
    echo -e "${GREEN}${BOLD}Passed:${NC}         $PASSED"
    echo -e "${RED}${BOLD}Failed:${NC}         $FAILED"
    echo -e "${YELLOW}${BOLD}Skipped:${NC}        $SKIPPED"
    echo ""

    if [[ $FAILED -eq 0 ]]; then
        echo -e "${GREEN}${BOLD}✓ ALL TESTS PASSED${NC}"
        echo ""
        return 0
    else
        echo -e "${RED}${BOLD}✗ SOME TESTS FAILED${NC}"
        echo ""
        return 1
    fi
}

#############################################################################
# Main Test Execution
#############################################################################
main() {
    echo -e "${BOLD}${MAGENTA}"
    echo "╔════════════════════════════════════════════════════════╗"
    echo "║                                                        ║"
    echo "║        INWX-CLI Comprehensive Test Suite              ║"
    echo "║                                                        ║"
    echo "╚════════════════════════════════════════════════════════╝"
    echo -e "${NC}"

    print_info "Testing binary: $BINARY"
    print_info "Test mode: INWX Test Environment"

    # Check prerequisites
    if ! check_credentials; then
        echo ""
        print_warning "Cannot run full test suite without credentials"
        print_info "Set INWX_USERNAME and INWX_PASSWORD environment variables"
        exit 1
    fi

    # Run all test suites
    test_basic_commands
    test_account_operations
    test_domain_operations
    test_dns_listing
    test_dns_create
    test_dns_update
    test_dns_delete
    test_dns_export_import
    test_dns_validation
    test_dns_verification
    test_backup_operations
    test_configuration
    test_error_handling
    test_output_formats

    # Print summary and exit
    if print_summary; then
        exit 0
    else
        exit 1
    fi
}

# Run main function
main "$@"
