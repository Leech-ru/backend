package news

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Delete removes news and its image (if any)
func (s *newsService) Delete(ctx context.Context, req *dto.DeleteNewsRequest) error {
	news, err := s.newsRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.NewsNotFound):
		return errorz.NewsNotFound
	case err != nil:
		return err
	}

	if err := s.imageService.Delete(ctx, &dto.DeleteImageRequest{ID: news.ImageID}); err != nil {
		return err
	}

	if err := s.newsRepo.Delete(ctx, req.ID); err != nil {
		return err
	}

	return nil
}
