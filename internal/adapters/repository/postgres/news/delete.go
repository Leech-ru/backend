package news

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

// Delete removes a category by ID
func (s *newsRepo) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.client.News.
		DeleteOneID(id).
		Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		return errorz.NewsNotFound
	case err != nil:
		return err
	}

	return nil
}
