package category

import (
	"Leech-ru/pkg/ent"
	"context"
)

// GetAllByFilter retrieves all category with optional pagination and filter.
func (s *categoryRepo) GetAllByFilter(ctx context.Context, limit, offset int) ([]*ent.Category, error) {
	category, err := s.client.Category.Query().
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return category, nil
}
