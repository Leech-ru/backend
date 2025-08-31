package mainpage

import (
	"Leech-ru/pkg/ent"
	"context"
	"fmt"
)

// GetAll retrieves all main pages with optional pagination and filter.
func (s *mainPageRepo) GetAll(ctx context.Context, limit, offset int) ([]*ent.MainPage, error) {
	pages, err := s.client.MainPage.Query().
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	return pages, nil
}
