package image

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// TODO при удалении изображения, не очищается поле в косметкие
func (s *imageService) Delete(ctx context.Context, req *dto.DeleteImageRequest) error {
	err := s.imageMinIORepo.Delete(ctx, req.ID)
	if err != nil {
		return err
	}
	return nil
}
