package news

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
)

// Create news and returns it.
func (s *newsService) Create(ctx context.Context, req *dto.CreateNewsRequest) (*dto.CreateNewsResponse, error) {
	news := &ent.News{
		ImageID:  req.ImageID,
		Title:    req.Title,
		Content:  req.Content,
		Href:     req.Href,
		IsHidden: req.IsHidden,
	}
	if cond, err := s.imageService.Exists(ctx, req.ImageID); err != nil || !cond {
		switch {
		case !cond:
			return nil, errorz.ImageNotFound
		case err != nil:
			return nil, err
		}
	}

	news, err := s.newsRepo.Create(ctx, *news)
	if err != nil {
		return nil, err
	}

	return &dto.CreateNewsResponse{
		ID:       news.ID,
		ImageID:  news.ImageID,
		Title:    news.Title,
		Content:  news.Content,
		Href:     news.Href,
		IsHidden: news.IsHidden,
	}, nil
}
