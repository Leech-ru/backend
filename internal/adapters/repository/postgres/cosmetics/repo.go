package cosmetics

import (
	"Leech-ru/pkg/ent"
)

// TODO все обновления должны быть в одной транзакции
type cosmeticsRepo struct {
	client *ent.Client
}

// NewCosmeticsRepo creates a new copy of the repository for Client
func NewCosmeticsRepo(client *ent.Client) *cosmeticsRepo {

	return &cosmeticsRepo{client: client}
}
