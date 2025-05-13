package token

import (
	"Leech-ru/pkg/ent"
)

type tokenRepo struct {
	client *ent.Client
}

// NewTokenRepo creates a new copy of the repository for Client
func NewTokenRepo(client *ent.Client) *tokenRepo {
	return &tokenRepo{client: client}
}
