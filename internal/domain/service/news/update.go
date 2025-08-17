package news

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Update change news data and optionally upload a new image file.
func (s *newsService) Update(ctx context.Context, req *dto.UpdateNewsRequest) (*dto.UpdateNewsResponse, error) {
	newsToUpdate, err := s.newsRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.NewsNotFound):
		return nil, errorz.NewsNotFound
	case err != nil:
		return nil, err
	}
	if req.ImageID != nil {
		if cond, err := s.imageService.Exists(ctx, *req.ImageID); err != nil || !cond {
			switch {
			case !cond:
				return nil, errorz.ImageNotFound
			default:
				return nil, err
			}
		}
		if err := s.imageService.Delete(ctx, &dto.DeleteImageRequest{ID: newsToUpdate.ImageID}); err != nil {
			switch {
			case errors.Is(err, errorz.ImageNotFound):
			case err != nil:
				return nil, err
			}
		}
		newsToUpdate.ImageID = *req.ImageID
	}
	if req.Title != nil {
		newsToUpdate.Title = *req.Title
	}
	if req.Content != nil {
		newsToUpdate.Title = *req.Title
	}
	if req.Href != nil {
		newsToUpdate.Title = *req.Title
	}
	if req.IsHidden != nil {
		newsToUpdate.IsHidden = *req.IsHidden
	}

	updatedNews, err := s.newsRepo.Update(ctx, *newsToUpdate)
	switch {
	case errors.Is(err, errorz.InvalidCosmeticsFormat):
		return nil, errorz.InvalidCosmeticsFormat
	case errors.Is(err, errorz.CategoryNotFound):
		return nil, errorz.CategoryNotFound
	case errors.Is(err, errorz.ImageNotFound):
		return nil, errorz.ImageNotFound
	case err != nil:
		return nil, err
	}

	return &dto.UpdateNewsResponse{
		ID:       updatedNews.ID,
		ImageID:  updatedNews.ImageID,
		Title:    updatedNews.Title,
		Content:  updatedNews.Content,
		Href:     updatedNews.Href,
		IsHidden: updatedNews.IsHidden,
	}, nil
}
