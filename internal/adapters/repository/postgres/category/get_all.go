package category

import (
	"Leech-ru/pkg/ent"
	"context"
	"fmt"
)

// GetAll retrieves all category with optional pagination and filter.
func (s *categoryRepo) GetAll(ctx context.Context, limit, offset int) ([]*ent.Category, int, error) {
	baseQuery := s.client.Category.Query()

	totalItems, err := baseQuery.Clone().Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count categories in db: %w", err)
	}

	category, err := baseQuery.
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to query db: %w", err)
	}

	return category, totalItems, nil
}
