package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"fmt"
)

// Create creates a new category in the database
func (s *categoryRepo) Create(ctx context.Context, entity ent.Category) (*ent.Category, error) {
	created, err := s.client.Category.
		Create().
		SetName(entity.Name).
		Save(ctx)

	switch {
	case ent.IsConstraintError(err):
		return nil, errorz.InvalidCategoryFormat
	case err != nil:
		return nil, fmt.Errorf("failed to query db: %w", err)
	}
	return created, nil
}
