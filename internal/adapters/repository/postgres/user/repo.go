package user

import (
	"Leech-ru/pkg/ent"
)

type userRepo struct {
	client *ent.Client
}

// NewUserRepo creates a new copy of the repository for Client
func NewUserRepo(client *ent.Client) *userRepo {
	return &userRepo{client: client}
}
