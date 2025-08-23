package news

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/news"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// GetById retrieves a category by ID
func (s *newsRepo) GetById(ctx context.Context, id uuid.UUID) (*ent.News, error) {
	c, err := s.client.News.
		Query().
		Where(news.ID(id)).
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.NewsNotFound
	case err != nil:
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	return c, nil
}
