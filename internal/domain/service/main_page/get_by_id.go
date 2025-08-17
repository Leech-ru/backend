package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// GetByID returns main page content by ID
func (s *mainPageService) GetByID(ctx context.Context, req *dto.GetByIdMainPageRequest) (*dto.GetByIdMainPageResponse, error) {
	mainPage, err := s.mainPageRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.MainPageNotFound):
		return nil, errorz.MainPageNotFound
	case err != nil:
		return nil, err
	}

	return &dto.GetByIdMainPageResponse{
		ID:       mainPage.ID,
		ImageID:  mainPage.ImageID,
		Title:    mainPage.Title,
		Content:  mainPage.Content,
		Href:     mainPage.Href,
		IsHidden: mainPage.IsHidden,
		Fluid:    mainPage.Fluid,
	}, nil
}
