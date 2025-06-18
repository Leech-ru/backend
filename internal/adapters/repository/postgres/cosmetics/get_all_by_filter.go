package cosmetics

import (
	"Leech-ru/internal/domain/types"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/cosmetics"
	"context"
)

// GetAllByFilter retrieves all cosmetics with optional pagination and filter.
func (s *cosmeticsRepo) GetAllByFilter(ctx context.Context, limit, offset int, category *types.Category, titlePrefix *string, volume *int) ([]*ent.Cosmetics, error) {
	query := s.client.Cosmetics.Query()

	if category != nil {
		query = query.Where(cosmetics.CategoryEQ(*category))
	}
	if titlePrefix != nil {
		query = query.Where(cosmetics.TitleHasPrefix(*titlePrefix))
	}
	if volume != nil {
		query = query.Where(cosmetics.VolumeEQ(*volume))
	}

	cosmetics, err := query.
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return cosmetics, nil
}
