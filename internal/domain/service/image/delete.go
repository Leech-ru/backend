package image

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

func (s *imageService) Delete(ctx context.Context, req *dto.DeleteImageRequest) error {
	err := s.imageMinIORepo.Delete(ctx, req.ID)
	if err != nil {
		return err
	}
	err = s.imageMinIORepo.Delete(ctx, req.ID)
	if err != nil {
		return err
	}

	return nil
}
