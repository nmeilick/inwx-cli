package backup

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/nmeilick/inwx-cli/pkg/inwx"
	"github.com/rs/zerolog/log"
)

// AtomicStore provides atomic backup operations with proper error handling
type AtomicStore struct {
	basePath string
	mutex    sync.RWMutex
}

// NewAtomicStore creates a new atomic backup store
func NewAtomicStore() (*AtomicStore, error) {
	basePath, err := locateBackupBase()
	if err != nil {
		return nil, fmt.Errorf("failed to locate backup directory: %w", err)
	}
	return &AtomicStore{basePath: basePath}, nil
}

// AtomicChange performs an atomic backup and change operation
// It creates a backup entry, executes the callback, and only keeps the backup if the callback succeeds
func (s *AtomicStore) AtomicChange(operation inwx.OperationType, record inwx.DNSRecord, context map[string]interface{}, callback func() error) (*inwx.BackupEntry, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Generate secure ID
	id, err := generateSecureID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate backup ID: %w", err)
	}

	// Create backup entry
	entry := &inwx.BackupEntry{
		ID:        id,
		Timestamp: time.Now(),
		Operation: operation,
		Record:    record,
		Context:   context,
	}

	// Write backup atomically
	backupPath, err := s.writeBackupAtomic(entry)
	if err != nil {
		return nil, fmt.Errorf("failed to create backup: %w", err)
	}

	log.Debug().
		Str("backup_id", entry.ID).
		Str("operation", string(operation)).
		Str("record_name", record.Name).
		Str("record_type", record.Type).
		Msg("Created backup entry")

	// Execute the callback
	if err := callback(); err != nil {
		// Callback failed, remove the backup entry
		if removeErr := os.Remove(backupPath); removeErr != nil {
			log.Warn().
				Err(removeErr).
				Str("backup_path", backupPath).
				Msg("Failed to remove backup after callback failure")
		}
		return nil, fmt.Errorf("operation failed: %w", err)
	}

	log.Debug().
		Str("backup_id", entry.ID).
		Msg("Operation succeeded, backup entry preserved")

	return entry, nil
}

// writeBackupAtomic writes a backup entry atomically using a temporary file
func (s *AtomicStore) writeBackupAtomic(entry *inwx.BackupEntry) (string, error) {
	data, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal backup entry: %w", err)
	}

	filename := fmt.Sprintf("%s_%s.json", entry.Timestamp.Format("20060102_150405"), entry.ID)
	finalPath := filepath.Join(s.basePath, filename)
	tempPath := finalPath + ".tmp"

	// Write to temporary file first
	if err := os.WriteFile(tempPath, data, 0644); err != nil {
		return "", fmt.Errorf("failed to write temporary backup file: %w", err)
	}

	// Atomically move to final location
	if err := os.Rename(tempPath, finalPath); err != nil {
		// Clean up temp file on failure
		os.Remove(tempPath)
		return "", fmt.Errorf("failed to move backup to final location: %w", err)
	}

	return finalPath, nil
}

// Save creates a backup entry (non-atomic version for compatibility)
func (s *AtomicStore) Save(operation inwx.OperationType, record inwx.DNSRecord, context map[string]interface{}) (*inwx.BackupEntry, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Generate secure ID
	id, err := generateSecureID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate backup ID: %w", err)
	}

	entry := &inwx.BackupEntry{
		ID:        id,
		Timestamp: time.Now(),
		Operation: operation,
		Record:    record,
		Context:   context,
	}

	_, err = s.writeBackupAtomic(entry)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

// Remove removes a backup entry by ID
func (s *AtomicStore) Remove(entryID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	files, err := filepath.Glob(filepath.Join(s.basePath, "*_"+entryID+".json"))
	if err != nil {
		return fmt.Errorf("failed to search for backup files: %w", err)
	}

	if len(files) == 0 {
		return fmt.Errorf("backup entry not found: %s", entryID)
	}

	for _, file := range files {
		if err := os.Remove(file); err != nil {
			return fmt.Errorf("failed to remove backup file %s: %w", file, err)
		}
	}

	return nil
}

// List returns all backup entries, with proper error handling
func (s *AtomicStore) List() ([]*inwx.BackupEntry, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	files, err := filepath.Glob(filepath.Join(s.basePath, "*.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to list backup files: %w", err)
	}

	var entries []*inwx.BackupEntry
	var errors []error

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to read backup file %s: %w", file, err))
			continue
		}

		var entry inwx.BackupEntry
		if err := json.Unmarshal(data, &entry); err != nil {
			errors = append(errors, fmt.Errorf("failed to parse backup file %s: %w", file, err))
			continue
		}

		entries = append(entries, &entry)
	}

	// Log any errors encountered but don't fail the entire operation
	for _, err := range errors {
		log.Warn().Err(err).Msg("Error processing backup file")
	}

	return entries, nil
}

// Get retrieves a specific backup entry by ID
func (s *AtomicStore) Get(entryID string) (*inwx.BackupEntry, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// First try exact match
	files, err := filepath.Glob(filepath.Join(s.basePath, "*_"+entryID+".json"))
	if err != nil {
		return nil, fmt.Errorf("failed to search for backup files: %w", err)
	}

	// If no exact match found, try partial match (for truncated IDs)
	if len(files) == 0 {
		allFiles, err := filepath.Glob(filepath.Join(s.basePath, "*.json"))
		if err != nil {
			return nil, fmt.Errorf("failed to search for backup files: %w", err)
		}

		var matchingFiles []string
		for _, file := range allFiles {
			// Extract the ID part from filename (format: timestamp_ID.json)
			basename := filepath.Base(file)
			parts := strings.Split(basename, "_")
			if len(parts) >= 2 {
				fileID := strings.TrimSuffix(parts[len(parts)-1], ".json")
				if strings.HasPrefix(fileID, entryID) {
					matchingFiles = append(matchingFiles, file)
				}
			}
		}

		if len(matchingFiles) == 0 {
			return nil, fmt.Errorf("backup entry not found: %s", entryID)
		}

		if len(matchingFiles) > 1 {
			return nil, fmt.Errorf("ambiguous backup ID %s: matches %d entries", entryID, len(matchingFiles))
		}

		files = matchingFiles
	}

	if len(files) > 1 {
		log.Warn().
			Str("entry_id", entryID).
			Int("file_count", len(files)).
			Msg("Multiple backup files found for entry ID")
	}

	data, err := os.ReadFile(files[0])
	if err != nil {
		return nil, fmt.Errorf("failed to read backup file: %w", err)
	}

	var entry inwx.BackupEntry
	if err := json.Unmarshal(data, &entry); err != nil {
		return nil, fmt.Errorf("failed to parse backup file: %w", err)
	}

	return &entry, nil
}

// PurgeOlderThan removes backup entries older than the specified duration
func (s *AtomicStore) PurgeOlderThan(duration time.Duration) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cutoff := time.Now().Add(-duration)
	entries, err := s.listUnsafe() // Use unsafe version since we already hold the lock
	if err != nil {
		return fmt.Errorf("failed to list backup entries: %w", err)
	}

	var errors []error
	purgedCount := 0

	for _, entry := range entries {
		if entry.Timestamp.Before(cutoff) {
			if err := s.removeUnsafe(entry.ID); err != nil {
				errors = append(errors, fmt.Errorf("failed to remove entry %s: %w", entry.ID, err))
			} else {
				purgedCount++
			}
		}
	}

	log.Info().
		Int("purged_count", purgedCount).
		Dur("older_than", duration).
		Msg("Purged old backup entries")

	if len(errors) > 0 {
		return fmt.Errorf("encountered %d errors during purge: %v", len(errors), errors[0])
	}

	return nil
}

// Verify checks the integrity of all backup files
func (s *AtomicStore) Verify() error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	files, err := filepath.Glob(filepath.Join(s.basePath, "*.json"))
	if err != nil {
		return fmt.Errorf("failed to list backup files: %w", err)
	}

	var errors []error
	validCount := 0

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			errors = append(errors, fmt.Errorf("cannot read %s: %w", file, err))
			continue
		}

		var entry inwx.BackupEntry
		if err := json.Unmarshal(data, &entry); err != nil {
			errors = append(errors, fmt.Errorf("invalid JSON in %s: %w", file, err))
			continue
		}

		// Basic validation
		if entry.ID == "" {
			errors = append(errors, fmt.Errorf("missing ID in %s", file))
			continue
		}

		if entry.Timestamp.IsZero() {
			errors = append(errors, fmt.Errorf("missing timestamp in %s", file))
			continue
		}

		if entry.Operation == "" {
			errors = append(errors, fmt.Errorf("missing operation in %s", file))
			continue
		}

		validCount++
	}

	log.Info().
		Int("total_files", len(files)).
		Int("valid_files", validCount).
		Int("error_count", len(errors)).
		Msg("Backup verification completed")

	if len(errors) > 0 {
		return fmt.Errorf("found %d corrupted backup files: %v", len(errors), errors[0])
	}

	return nil
}

// listUnsafe is the internal version of List that doesn't acquire locks
func (s *AtomicStore) listUnsafe() ([]*inwx.BackupEntry, error) {
	files, err := filepath.Glob(filepath.Join(s.basePath, "*.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to list backup files: %w", err)
	}

	var entries []*inwx.BackupEntry
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			continue // Skip unreadable files
		}

		var entry inwx.BackupEntry
		if err := json.Unmarshal(data, &entry); err != nil {
			continue // Skip corrupted files
		}

		entries = append(entries, &entry)
	}

	return entries, nil
}

// removeUnsafe is the internal version of Remove that doesn't acquire locks
func (s *AtomicStore) removeUnsafe(entryID string) error {
	files, err := filepath.Glob(filepath.Join(s.basePath, "*_"+entryID+".json"))
	if err != nil {
		return fmt.Errorf("failed to search for backup files: %w", err)
	}

	for _, file := range files {
		if err := os.Remove(file); err != nil {
			return fmt.Errorf("failed to remove backup file %s: %w", file, err)
		}
	}

	return nil
}

// generateSecureID generates a cryptographically secure random ID
func generateSecureID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate secure ID: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}
