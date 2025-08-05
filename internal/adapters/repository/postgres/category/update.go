package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Update updates an existing category
func (s *categoryRepo) Update(ctx context.Context, entity ent.Category) (*ent.Category, error) {
	updated, err := s.client.Category.
		UpdateOneID(entity.ID).
		SetName(entity.Name).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.CategoryNotFound
	case err != nil:
		return nil, err
	}

	return updated, nil
}
