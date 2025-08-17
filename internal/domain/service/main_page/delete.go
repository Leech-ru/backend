package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Delete removes a main page
func (s *mainPageService) Delete(ctx context.Context, req *dto.DeleteMainPageRequest) error {
	mainPage, err := s.mainPageRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.MainPageNotFound):
		return errorz.MainPageNotFound
	case err != nil:
		return err
	}

	if err := s.imageService.Delete(ctx, &dto.DeleteImageRequest{ID: mainPage.ImageID}); err != nil {
		return err
	}

	if err := s.mainPageRepo.Delete(ctx, mainPage.ID); err != nil {
		return err
	}

	return nil
}
