package backup

import (
	"os"
	"path/filepath"
	"time"

	"github.com/adrg/xdg"
	"github.com/nmeilick/inwx-cli/pkg/inwx"
)

// BackupStore interface defines the contract for backup storage
type BackupStore interface {
	// AtomicChange performs an atomic backup and change operation
	AtomicChange(operation inwx.OperationType, record inwx.DNSRecord, context map[string]interface{}, callback func() error) (*inwx.BackupEntry, error)

	// Save creates a backup entry (for compatibility)
	Save(operation inwx.OperationType, record inwx.DNSRecord, context map[string]interface{}) (*inwx.BackupEntry, error)

	// Remove removes a backup entry by ID
	Remove(entryID string) error

	// List returns all backup entries
	List() ([]*inwx.BackupEntry, error)

	// Get retrieves a specific backup entry by ID
	Get(entryID string) (*inwx.BackupEntry, error)

	// PurgeOlderThan removes backup entries older than the specified duration
	PurgeOlderThan(duration time.Duration) error

	// Verify checks the integrity of all backup files
	Verify() error
}

// locateBackupBase determines the directory for storing backups using XDG data directory.
// It creates the directory if it does not exist.
func locateBackupBase() (string, error) {
	// Use XDG data home (defaults to $HOME/.local/share)
	basePath := filepath.Join(xdg.DataHome, "inwx", "backups")
	if err := os.MkdirAll(basePath, 0o755); err != nil {
		return "", err
	}
	return basePath, nil
}

func NewStore() (BackupStore, error) {
	return NewAtomicStore()
}
