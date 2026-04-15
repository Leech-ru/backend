package news

import (
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/news"
	"context"
	"fmt"
)

// GetAllByFilter retrieves all category with optional pagination and filter.
func (s *newsRepo) GetAllByFilter(ctx context.Context, limit, offset int, isHidden *bool) ([]*ent.News, int, error) {
	baseQuery := s.client.News.Query()

	if isHidden != nil {
		baseQuery = baseQuery.Where(news.IsHiddenEQ(*isHidden))
	}

	totalItems, err := baseQuery.Clone().Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count news in db: %w", err)
	}

	news, err := baseQuery.
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to query db: %w", err)
	}

	return news, totalItems, nil
}
