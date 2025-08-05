package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

// Delete removes a category by ID
func (s *categoryRepo) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.client.Category.
		DeleteOneID(id).
		Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		return errorz.CategoryNotFound
	case err != nil:
		return err
	}

	return nil
}
