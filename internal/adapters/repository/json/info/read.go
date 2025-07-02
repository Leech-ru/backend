package info

import (
	"encoding/json"
	"os"
)

// Read always opens the file, reads and parses data again.
func (r *Repository) Read() (*Info, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	bytes, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var info Info
	if err := json.Unmarshal(bytes, &info); err != nil {
		return nil, err
	}

	return &info, nil
}
