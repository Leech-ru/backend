package partner

import (
	"Leech-ru/pkg/ent"
)

type partnersRepo struct {
	client *ent.Client
}

// NewPartnersRepo creates a new copy of the repository for Client
func NewPartnersRepo(client *ent.Client) *partnersRepo {

	return &partnersRepo{client: client}
}
