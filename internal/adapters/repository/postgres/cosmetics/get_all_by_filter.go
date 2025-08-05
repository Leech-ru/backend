package cosmetics

import (
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/category"
	"Leech-ru/pkg/ent/cosmetics"
	"context"
	"github.com/google/uuid"
)

// GetAllByFilter retrieves all cosmetics with optional pagination and filter.
func (s *cosmeticsRepo) GetAllByFilter(ctx context.Context, limit, offset int, categoryID *uuid.UUID, titlePrefix *string, volume *int, isHidden *bool) ([]*ent.Cosmetics, error) {
	query := s.client.Cosmetics.Query()

	if categoryID != nil {
		query = query.Where(cosmetics.HasCategoryWith(category.IDEQ(*categoryID)))
	}
	if titlePrefix != nil {
		query = query.Where(cosmetics.TitleHasPrefix(*titlePrefix))
	}
	if volume != nil {
		query = query.Where(cosmetics.VolumeEQ(*volume))
	}
	if isHidden != nil {
		query = query.Where(cosmetics.IsHiddenEQ(*isHidden))
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
