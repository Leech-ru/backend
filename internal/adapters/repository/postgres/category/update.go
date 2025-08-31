package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"fmt"
)

// Update updates an existing category
func (s *categoryRepo) Update(ctx context.Context, entity ent.Category) (*ent.Category, error) {
	updated, err := s.client.Category.
		UpdateOneID(entity.ID).
		SetName(entity.Name).
		SetImageID(entity.ImageID).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.CategoryNotFound
	case ent.IsConstraintError(err):
		return nil, errorz.InvalidCategoryFormat
	case err != nil:
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	return updated, nil
}
