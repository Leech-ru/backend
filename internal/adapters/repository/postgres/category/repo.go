package category

import (
	"Leech-ru/pkg/ent"
)

type categoryRepo struct {
	client *ent.Client
}

// NewCategoryRepo creates a new copy of the repository for Client
func NewCategoryRepo(client *ent.Client) *categoryRepo {

	return &categoryRepo{client: client}
}
