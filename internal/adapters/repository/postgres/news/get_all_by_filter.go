package news

import (
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/news"
	"context"
	"fmt"
)

// GetAllByFilter retrieves all category with optional pagination and filter.
func (s *newsRepo) GetAllByFilter(ctx context.Context, limit, offset int, isHidden *bool) ([]*ent.News, error) {
	query := s.client.News.Query()

	if isHidden != nil {
		query = query.Where(news.IsHiddenEQ(*isHidden))
	}

	news, err := query.
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	return news, nil
}
