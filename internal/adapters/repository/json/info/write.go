package info

import (
	"encoding/json"
	"os"
)

// Write serializes the structure in JSON and writes to the file
func (r *Repository) Write(info *Info) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	bytes, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, bytes, 0644)
}
