package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/category"
	"context"
	"github.com/google/uuid"
)

// GetById retrieves a category by ID
func (s *categoryRepo) GetById(ctx context.Context, id uuid.UUID) (*ent.Category, error) {
	c, err := s.client.Category.
		Query().
		Where(category.ID(id)).
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.CategoryNotFound
	case err != nil:
		return nil, err
	}

	return c, nil
}
