package mainpage

import (
	"Leech-ru/pkg/ent"
	"context"
	"fmt"
)

// GetAll retrieves all main pages with optional pagination and filter.
func (s *mainPageRepo) GetAll(ctx context.Context, limit, offset int) ([]*ent.MainPage, int, error) {
	baseQuery := s.client.MainPage.Query()

	totalItems, err := baseQuery.Clone().Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count main pages in db: %w", err)
	}

	pages, err := baseQuery.
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to query db: %w", err)
	}

	return pages, totalItems, nil
}
