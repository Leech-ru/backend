package image

import (
	"Leech-ru/pkg/ent"
)

type imageRepo struct {
	client *ent.Client
}

// NewImageRepo creates a new copy of the repository for Client
func NewImageRepo(client *ent.Client) *imageRepo {
	return &imageRepo{client: client}
}
