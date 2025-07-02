package info

import (
	"sync"
)

type Repository struct {
	filePath string
	mu       sync.Mutex
}

// NewRepository creates a repository indicating the path to the file
func NewRepository(filePath string) *Repository {
	return &Repository{filePath: filePath}
}
