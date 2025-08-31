package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/mainpage"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// GetById retrieves a page content by ID
func (s *mainPageRepo) GetById(ctx context.Context, id uuid.UUID) (*ent.MainPage, error) {
	c, err := s.client.MainPage.
		Query().
		Where(mainpage.ID(id)).
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.MainPageNotFound
	case err != nil:
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	return c, nil
}
