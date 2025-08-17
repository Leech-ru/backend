package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Update changes main page content data and optionally uploads a new image file
func (s *mainPageService) Update(ctx context.Context, req *dto.UpdateMainPageRequest) (*dto.UpdateMainPageResponse, error) {
	pageToUpdate, err := s.mainPageRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.MainPageNotFound):
		return nil, errorz.MainPageNotFound
	case err != nil:
		return nil, err
	}

	if req.ImageID != nil {
		if cond, err := s.imageService.Exists(ctx, *req.ImageID); err != nil || !cond {
			switch {
			case !cond:
				return nil, errorz.MainPageNotFound
			default:
				return nil, err
			}
		}
		if err := s.imageService.Delete(ctx, &dto.DeleteImageRequest{ID: pageToUpdate.ImageID}); err != nil {
			switch {
			case errors.Is(err, errorz.ImageNotFound):
				return nil, errorz.ImageNotFound
			case err != nil:
				return nil, err
			}
		}
		pageToUpdate.ImageID = *req.ImageID
	}

	if req.Title != nil {
		pageToUpdate.Title = *req.Title
	}

	if req.Content != nil {
		pageToUpdate.Content = *req.Content
	}

	if req.Href != nil {
		pageToUpdate.Href = *req.Href
	}

	if req.IsHidden != nil {
		pageToUpdate.IsHidden = *req.IsHidden
	}

	if req.Fluid != nil {
		pageToUpdate.Fluid = *req.Fluid
	}

	updatedPage, err := s.mainPageRepo.Update(ctx, *pageToUpdate)
	switch {
	case errors.Is(err, errorz.InvalidMainPageFormat):
		return nil, errorz.InvalidMainPageFormat
	case errors.Is(err, errorz.MainPageNotFound):
		return nil, errorz.MainPageNotFound
	case errors.Is(err, errorz.ImageNotFound):
		return nil, errorz.ImageNotFound
	}

	return &dto.UpdateMainPageResponse{
		ID:       updatedPage.ID,
		ImageID:  updatedPage.ImageID,
		Title:    updatedPage.Title,
		Content:  updatedPage.Content,
		Href:     updatedPage.Href,
		IsHidden: updatedPage.IsHidden,
		Fluid:    updatedPage.Fluid,
	}, nil
}
