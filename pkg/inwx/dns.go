package inwx

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

// BackupStore interface for dependency injection
type BackupStore interface {
	AtomicChange(operation OperationType, record DNSRecord, context map[string]interface{}, callback func() error) (*BackupEntry, error)
	Save(operation OperationType, record DNSRecord, context map[string]interface{}) (*BackupEntry, error)
	Remove(entryID string) error
}

type OperationType string

const (
	OperationCreate OperationType = "create"
	OperationUpdate OperationType = "update"
	OperationDelete OperationType = "delete"
)

type BackupEntry struct {
	ID        string                 `json:"id"`
	Timestamp time.Time              `json:"timestamp"`
	Operation OperationType          `json:"operation"`
	Record    DNSRecord              `json:"record"`
	Context   map[string]interface{} `json:"context"`
}

type DNSService struct {
	client      *Client
	domain      string
	defaultTTL  int
	backupStore BackupStore
}

type DNSOption func(*DNSService)

func WithDomain(domain string) DNSOption {
	return func(s *DNSService) {
		s.domain = domain
	}
}

func WithDefaultTTL(ttl int) DNSOption {
	return func(s *DNSService) {
		s.defaultTTL = ttl
	}
}

func WithBackupStore(store BackupStore) DNSOption {
	return func(s *DNSService) {
		s.backupStore = store
	}
}

type DNSRecord struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
	Prio    int    `json:"prio,omitempty"`
	Domain  string `json:"domain,omitempty"`
}

type RecordFilter func(*RecordQuery)

type RecordQuery struct {
	Domain  string
	Types   []string
	Name    string
	Content string
}

func WithRecordType(types ...string) RecordFilter {
	return func(q *RecordQuery) {
		q.Types = types
	}
}

func WithRecordName(pattern string) RecordFilter {
	return func(q *RecordQuery) {
		q.Name = pattern
	}
}

func WithRecordContent(pattern string) RecordFilter {
	return func(q *RecordQuery) {
		q.Content = pattern
	}
}

func WithRecordID(id int) RecordFilter {
	return func(q *RecordQuery) {
		// This would be used for single record lookups
		// Implementation depends on API capabilities
	}
}

func WithRecordTTL(ttl int) RecordFilter {
	return func(q *RecordQuery) {
		// This would be used for TTL-based filtering
		// Implementation depends on API capabilities
	}
}

func WithDomainFilter(domain string) RecordFilter {
	return func(q *RecordQuery) {
		q.Domain = domain
	}
}

type ExportFormat int

const (
	ExportJSON ExportFormat = iota
	ExportZonefileFormat
)

type ImportFormat int

const (
	ImportJSON ImportFormat = iota
	ImportZonefileFormat
)

func (s *DNSService) ListRecords(ctx context.Context, filters ...RecordFilter) ([]DNSRecord, error) {
	query := &RecordQuery{
		Domain: s.domain,
	}

	for _, filter := range filters {
		filter(query)
	}

	// If no domain is specified, we need to list all domains first and then get records for each
	if query.Domain == "" {
		return s.listAllRecords(ctx, query)
	}

	// If multiple types are specified, make separate calls for each type
	if len(query.Types) > 1 {
		return s.listRecordsMultipleTypes(ctx, query)
	}

	params := make(map[string]interface{})
	params["domain"] = query.Domain

	if len(query.Types) == 1 {
		params["type"] = query.Types[0]
	}
	if query.Name != "" {
		params["name"] = query.Name
	}
	if query.Content != "" {
		params["content"] = query.Content
	}

	log.Debug().
		Interface("params", params).
		Msg("Calling nameserver.info")

	response, err := s.client.transport.Call(ctx, "nameserver.info", params)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	records := s.parseRecordsFromResponse(response, query.Domain)
	return records, nil
}

func (s *DNSService) listRecordsMultipleTypes(ctx context.Context, query *RecordQuery) ([]DNSRecord, error) {
	var allRecords []DNSRecord

	for _, recordType := range query.Types {
		params := make(map[string]interface{})
		params["domain"] = query.Domain
		params["type"] = recordType

		if query.Name != "" {
			params["name"] = query.Name
		}
		if query.Content != "" {
			params["content"] = query.Content
		}

		log.Debug().
			Interface("params", params).
			Str("type", recordType).
			Msg("Calling nameserver.info for specific type")

		response, err := s.client.transport.Call(ctx, "nameserver.info", params)
		if err != nil {
			log.Warn().
				Err(err).
				Str("type", recordType).
				Msg("Failed to get records for type, skipping")
			continue
		}

		records := s.parseRecordsFromResponse(response, query.Domain)
		allRecords = append(allRecords, records...)
	}

	return allRecords, nil
}

func (s *DNSService) listAllRecords(ctx context.Context, query *RecordQuery) ([]DNSRecord, error) {
	// First, get list of all domains
	domainResponse, err := s.client.transport.Call(ctx, "domain.list", map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to list domains: %w", err)
	}

	var allRecords []DNSRecord

	// Extract domains from response
	var domains []string
	if resData, ok := domainResponse["resData"].(map[string]interface{}); ok {
		log.Debug().
			Interface("resData", resData).
			Msg("Processing domain list resData")

		if domainList, ok := resData["domain"].([]interface{}); ok {
			log.Debug().
				Int("domain_list_length", len(domainList)).
				Msg("Found domain list")

			for i, d := range domainList {
				if domain, ok := d.(map[string]interface{}); ok {
					log.Debug().
						Int("index", i).
						Interface("domain_object", domain).
						Msg("Processing domain object")

					if domainName, ok := domain["domain"].(string); ok {
						log.Debug().
							Str("domain_name", domainName).
							Msg("Found domain name")
						domains = append(domains, domainName)
					} else {
						log.Debug().
							Interface("domain_field", domain["domain"]).
							Msg("Domain field is not a string or missing")
					}
				} else {
					log.Debug().
						Int("index", i).
						Interface("domain_item", d).
						Msg("Domain item is not a map")
				}
			}
		} else {
			log.Debug().
				Interface("domain_field", resData["domain"]).
				Msg("Domain field is not an array or missing")
		}
	} else {
		log.Debug().
			Interface("response", domainResponse).
			Msg("resData is not a map or missing")
	}

	log.Debug().
		Int("domain_count", len(domains)).
		Msg("Found domains, fetching records for each")

	// Get records for each domain
	for _, domain := range domains {
		// If multiple types are specified, make separate calls for each type
		if len(query.Types) > 1 {
			for _, recordType := range query.Types {
				params := map[string]interface{}{
					"domain": domain,
					"type":   recordType,
				}

				if query.Name != "" {
					params["name"] = query.Name
				}
				if query.Content != "" {
					params["content"] = query.Content
				}

				log.Debug().
					Str("domain", domain).
					Str("type", recordType).
					Interface("params", params).
					Msg("Calling nameserver.info for domain and type")

				response, err := s.client.transport.Call(ctx, "nameserver.info", params)
				if err != nil {
					log.Warn().
						Err(err).
						Str("domain", domain).
						Str("type", recordType).
						Msg("Failed to get records for domain and type, skipping")
					continue
				}

				records := s.parseRecordsFromResponse(response, domain)
				allRecords = append(allRecords, records...)
			}
			continue
		}

		params := map[string]interface{}{
			"domain": domain,
		}

		if len(query.Types) == 1 {
			params["type"] = query.Types[0]
		}
		if query.Name != "" {
			params["name"] = query.Name
		}
		if query.Content != "" {
			params["content"] = query.Content
		}

		log.Debug().
			Str("domain", domain).
			Interface("params", params).
			Msg("Calling nameserver.info for domain")

		response, err := s.client.transport.Call(ctx, "nameserver.info", params)
		if err != nil {
			log.Warn().
				Err(err).
				Str("domain", domain).
				Msg("Failed to get records for domain, skipping")
			continue
		}

		records := s.parseRecordsFromResponse(response, domain)
		allRecords = append(allRecords, records...)
	}

	return allRecords, nil
}

func (s *DNSService) parseRecordsFromResponse(response map[string]interface{}, domain string) []DNSRecord {
	var records []DNSRecord

	log.Debug().
		Interface("response", response).
		Str("domain", domain).
		Msg("nameserver.info response")

	// Debug the response structure
	if resData, ok := response["resData"].(map[string]interface{}); ok {
		log.Debug().
			Interface("resData", resData).
			Msg("Processing resData")
		// Try different possible field names for the record list
		var recordList []interface{}

		if rl, ok := resData["record"].([]interface{}); ok {
			recordList = rl
		} else if rl, ok := resData["records"].([]interface{}); ok {
			recordList = rl
		} else if rl, ok := resData["nameserver"].([]interface{}); ok {
			recordList = rl
		} else {
			// If no array found, check if resData itself contains record data
			for _, value := range resData {
				if rl, ok := value.([]interface{}); ok && len(rl) > 0 {
					// Check if this looks like a record array
					if len(rl) > 0 {
						if record, ok := rl[0].(map[string]interface{}); ok {
							if _, hasType := record["type"]; hasType {
								recordList = rl
								break
							}
						}
					}
				}
			}
		}

		log.Debug().
			Int("record_count", len(recordList)).
			Msg("Found record list")

		for i, r := range recordList {
			if record, ok := r.(map[string]interface{}); ok {
				log.Debug().
					Int("index", i).
					Interface("record", record).
					Msg("Processing record")

				dnsRecord := DNSRecord{}
				if id, ok := record["id"].(float64); ok {
					dnsRecord.ID = int(id)
				}
				if name, ok := record["name"].(string); ok {
					if name == domain {
						dnsRecord.Name = "@"
					} else {
						dnsRecord.Name = strings.TrimSuffix(name, "."+domain)
					}
				}
				if recordType, ok := record["type"].(string); ok {
					dnsRecord.Type = recordType
				}
				if content, ok := record["content"].(string); ok {
					dnsRecord.Content = content
				}
				if ttl, ok := record["ttl"].(float64); ok {
					dnsRecord.TTL = int(ttl)
				}
				if prio, ok := record["prio"].(float64); ok {
					dnsRecord.Prio = int(prio)
				}
				if recordDomain, ok := record["domain"].(string); ok {
					dnsRecord.Domain = recordDomain
				} else {
					// Use the provided domain if not specified in record
					dnsRecord.Domain = domain
				}

				log.Debug().
					Interface("parsed_record", dnsRecord).
					Msg("Parsed DNS record")

				records = append(records, dnsRecord)
			}
		}
	}

	return records
}

func (s *DNSService) CreateRecord(ctx context.Context, record DNSRecord) (*DNSRecord, error) {
	params := map[string]interface{}{
		"domain":  record.Domain,
		"type":    record.Type,
		"name":    record.Name,
		"content": record.Content,
		"ttl":     record.TTL,
	}

	if record.Prio > 0 {
		params["prio"] = record.Prio
	}

	// Use atomic backup if available
	if s.backupStore != nil {
		context := map[string]interface{}{
			"command": "create",
			"params":  params,
		}

		var createdRecord DNSRecord
		_, err := s.backupStore.AtomicChange(OperationCreate, record, context, func() error {
			response, err := s.client.transport.Call(ctx, "nameserver.createRecord", params)
			if err != nil {
				return err
			}

			createdRecord = record
			if resData, ok := response["resData"].(map[string]interface{}); ok {
				if id, ok := resData["id"].(float64); ok {
					createdRecord.ID = int(id)
				}
			}

			return nil
		})

		if err != nil {
			return nil, err
		}

		return &createdRecord, nil
	}

	// Fallback to non-atomic operation
	response, err := s.client.transport.Call(ctx, "nameserver.createRecord", params)
	if err != nil {
		return nil, err
	}

	if resData, ok := response["resData"].(map[string]interface{}); ok {
		if id, ok := resData["id"].(float64); ok {
			record.ID = int(id)
		}
	}

	return &record, nil
}

func (s *DNSService) UpdateRecord(ctx context.Context, id int, updates DNSRecord) (*DNSRecord, error) {
	params := map[string]interface{}{
		"id": id,
	}

	if updates.Name != "" {
		params["name"] = updates.Name
	}
	if updates.Type != "" {
		params["type"] = updates.Type
	}
	if updates.Content != "" {
		params["content"] = updates.Content
	}
	if updates.TTL > 0 {
		params["ttl"] = updates.TTL
	}
	if updates.Prio > 0 {
		params["prio"] = updates.Prio
	}

	// Use atomic backup if available
	if s.backupStore != nil {
		// Get the current record state before making changes
		originalRecord, err := s.getRecordByID(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("failed to get original record for backup: %w", err)
		}

		context := map[string]interface{}{
			"command": "update",
			"params":  params,
			"id":      id,
		}

		_, err = s.backupStore.AtomicChange(OperationUpdate, *originalRecord, context, func() error {
			_, err := s.client.transport.Call(ctx, "nameserver.updateRecord", params)
			return err
		})

		if err != nil {
			return nil, err
		}

		return &updates, nil
	}

	// Fallback to non-atomic operation
	_, err := s.client.transport.Call(ctx, "nameserver.updateRecord", params)
	if err != nil {
		return nil, err
	}

	return &updates, nil
}

func (s *DNSService) DeleteRecord(ctx context.Context, id int) error {
	params := map[string]interface{}{
		"id": id,
	}

	// Use atomic backup if available
	if s.backupStore != nil {
		// Get the current record state before deletion
		recordToDelete, err := s.getRecordByID(ctx, id)
		if err != nil {
			return fmt.Errorf("failed to get record for backup: %w", err)
		}

		context := map[string]interface{}{
			"command": "delete",
			"params":  params,
			"id":      id,
		}

		_, err = s.backupStore.AtomicChange(OperationDelete, *recordToDelete, context, func() error {
			_, err := s.client.transport.Call(ctx, "nameserver.deleteRecord", params)
			return err
		})

		return err
	}

	// Fallback to non-atomic operation
	_, err := s.client.transport.Call(ctx, "nameserver.deleteRecord", params)
	return err
}

func exportJSON(records []DNSRecord) ([]byte, error) {
	return json.MarshalIndent(records, "", "  ")
}

func importJSON(data []byte) ([]DNSRecord, error) {
	var records []DNSRecord
	err := json.Unmarshal(data, &records)
	return records, err
}

func (s *DNSService) ExportRecords(ctx context.Context, format ExportFormat) ([]byte, error) {
	records, err := s.ListRecords(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list records for export: %w", err)
	}

	switch format {
	case ExportJSON:
		return exportJSON(records)
	case ExportZonefileFormat:
		return ExportZonefile(records, s.domain)
	default:
		return nil, fmt.Errorf("unsupported export format")
	}
}

func (s *DNSService) ImportRecords(ctx context.Context, data []byte, format ImportFormat) error {
	var records []DNSRecord
	var err error

	switch format {
	case ImportJSON:
		records, err = importJSON(data)
	case ImportZonefileFormat:
		records, err = ImportZonefile(data, s.domain)
	default:
		return fmt.Errorf("unsupported import format")
	}

	if err != nil {
		return err
	}

	for _, record := range records {
		if record.Domain == "" {
			record.Domain = s.domain
		}
		if record.TTL == 0 {
			record.TTL = s.defaultTTL
		}
		_, err := s.CreateRecord(ctx, record)
		if err != nil {
			return err
		}
	}

	return nil
}

// getRecordByID retrieves a specific DNS record by its ID
func (s *DNSService) getRecordByID(ctx context.Context, id int) (*DNSRecord, error) {
	// Get all records and find the one with matching ID
	// This is not optimal but necessary since the API doesn't support direct ID lookup
	records, err := s.ListRecords(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list records: %w", err)
	}

	for _, record := range records {
		if record.ID == id {
			return &record, nil
		}
	}

	return nil, fmt.Errorf("record with ID %d not found", id)
}

func (s *DNSService) ImportRecordsWithSync(ctx context.Context, records []DNSRecord, format ImportFormat) error {
	// Get existing records
	existingRecords, err := s.ListRecords(ctx)
	if err != nil {
		return fmt.Errorf("failed to list existing records: %w", err)
	}

	// Create a map of records to import for quick lookup
	importMap := make(map[string]DNSRecord)
	for _, record := range records {
		if record.Domain == "" {
			record.Domain = s.domain
		}
		if record.TTL == 0 {
			record.TTL = s.defaultTTL
		}

		// Create a key for comparison (excluding ID)
		key := fmt.Sprintf("%s|%s|%s|%s", record.Domain, record.Name, record.Type, record.Content)
		importMap[key] = record
	}

	// Delete records that are not in the import
	for _, existing := range existingRecords {
		key := fmt.Sprintf("%s|%s|%s|%s", existing.Domain, existing.Name, existing.Type, existing.Content)
		if _, found := importMap[key]; !found {
			// Skip SOA records as they shouldn't be deleted
			if existing.Type == "SOA" {
				continue
			}

			log.Debug().
				Int("id", existing.ID).
				Str("name", existing.Name).
				Str("type", existing.Type).
				Msg("Deleting record not present in import")

			if err := s.DeleteRecord(ctx, existing.ID); err != nil {
				log.Warn().
					Err(err).
					Int("id", existing.ID).
					Msg("Failed to delete record")
			}
		}
	}

	// Create/update records from import
	for _, record := range records {
		_, err := s.CreateRecord(ctx, record)
		if err != nil {
			log.Warn().
				Err(err).
				Str("name", record.Name).
				Str("type", record.Type).
				Msg("Failed to create record (may already exist)")
		}
	}

	return nil
}
