package info

import (
	"encoding/json"
	"os"
)

// Write serializes the structure in JSON, writes it to the file, and returns the written model
func (r *Repository) Write(info *Info) (*Info, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	bytes, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(r.filePath, bytes, 0644); err != nil {
		return nil, err
	}

	copied := *info
	return &copied, nil
}
