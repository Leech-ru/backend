package cosmetics

import (
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/category"
	"Leech-ru/pkg/ent/cosmetics"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// GetAllByFilter retrieves all cosmetics with optional pagination and filter.
func (s *cosmeticsRepo) GetAllByFilter(ctx context.Context, limit, offset int, categoryID *uuid.UUID, titlePrefix *string, volume *int, isHidden *bool) ([]*ent.Cosmetics, int, error) {
	baseQuery := s.client.Cosmetics.Query().WithCategory()

	if categoryID != nil {
		baseQuery = baseQuery.Where(cosmetics.HasCategoryWith(category.IDEQ(*categoryID)))
	}
	if titlePrefix != nil {
		baseQuery = baseQuery.Where(cosmetics.TitleHasPrefix(*titlePrefix))
	}
	if volume != nil {
		baseQuery = baseQuery.Where(cosmetics.VolumeEQ(*volume))
	}
	if isHidden != nil {
		baseQuery = baseQuery.Where(cosmetics.IsHiddenEQ(*isHidden))
	}

	totalItems, err := baseQuery.Clone().Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count cosmetics in db: %w", err)
	}

	cosmetics, err := baseQuery.
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to query db: %w", err)
	}

	return cosmetics, totalItems, nil
}
