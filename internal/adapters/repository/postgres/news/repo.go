package news

import (
	"Leech-ru/pkg/ent"
)

type newsRepo struct {
	client *ent.Client
}

// NewNewsRepo creates a new copy of the repository for Client
func NewNewsRepo(client *ent.Client) *newsRepo {

	return &newsRepo{client: client}
}
