package cosmetics

import (
	"Leech-ru/pkg/ent"
	"context"
)

// GetAll retrieves all cosmetics with optional pagination
func (s *cosmeticsRepo) GetAll(ctx context.Context, limit, offset int) ([]*ent.Cosmetics, error) {
	cosmetics, err := s.client.Cosmetics.
		Query().
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return cosmetics, nil
}

//TODO GetAllByFilter
