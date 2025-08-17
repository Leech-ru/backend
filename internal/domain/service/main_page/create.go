package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
)

// Create creates a new main page content
func (s *mainPageService) Create(ctx context.Context, req *dto.CreateMainPageRequest) (*dto.CreateMainPageResponse, error) {
	mainPage := &ent.MainPage{
		ImageID:  req.ImageID,
		Title:    req.Title,
		Content:  req.Content,
		Href:     req.Href,
		IsHidden: req.IsHidden,
		Fluid:    req.Fluid,
	}

	if cond, err := s.imageService.Exists(ctx, req.ImageID); err != nil {
		switch {
		case !cond:
			return nil, errorz.ImageNotFound
		case err != nil:
			return nil, err
		}
	}

	mainPage, err := s.mainPageRepo.Create(ctx, *mainPage)
	if err != nil {
		return nil, err
	}

	return &dto.CreateMainPageResponse{
		ID:       mainPage.ID,
		ImageID:  mainPage.ImageID,
		Title:    mainPage.Title,
		Content:  mainPage.Content,
		Href:     mainPage.Href,
		IsHidden: mainPage.IsHidden,
		Fluid:    mainPage.Fluid,
	}, nil
}
