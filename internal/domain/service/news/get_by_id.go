package news

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// GetByID returns the news by ID.
func (s *newsService) GetByID(ctx context.Context, req *dto.GetByIdNewsRequest) (*dto.GetByIdNewsResponse, error) {
	news, err := s.newsRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.NewsNotFound):
		return nil, errorz.NewsNotFound
	case err != nil:
		return nil, err
	}

	return &dto.GetByIdNewsResponse{
		ID:       news.ID,
		ImageID:  news.ImageID,
		Title:    news.Title,
		Content:  news.Content,
		IsHidden: news.IsHidden,
	}, nil
}
