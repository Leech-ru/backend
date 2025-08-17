package news

import (
	"Leech-ru/pkg/ent"
	"context"
)

// GetAll retrieves all category with optional pagination and filter.
func (s *newsRepo) GetAll(ctx context.Context, limit, offset int) ([]*ent.News, error) {
	news, err := s.client.News.Query().
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return news, nil
}
