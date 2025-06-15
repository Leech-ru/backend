package cosmetics

import (
	"Leech-ru/pkg/ent"
)

type cosmeticsRepo struct {
	client *ent.Client
}

// NewCosmeticsRepo creates a new copy of the repository for Client
func NewCosmeticsRepo(client *ent.Client) *cosmeticsRepo {

	return &cosmeticsRepo{client: client}
}
