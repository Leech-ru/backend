package mainpage

import (
	"Leech-ru/pkg/ent"
)

type mainPageRepo struct {
	client *ent.Client
}

// NewMainPageRepo creates a new copy of the repository for Client
func NewMainPageRepo(client *ent.Client) *mainPageRepo {
	return &mainPageRepo{client: client}
}
