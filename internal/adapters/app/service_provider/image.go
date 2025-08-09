package service_provider

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/service/image"
	"context"
)

type imageService interface {
	Create(ctx context.Context, req *dto.CreateImageRequest) (*dto.CreateImageResponse, error)
	GetById(ctx context.Context, req *dto.GetByIdImageRequest) (*dto.GetByIdImageResponse, error)
	Delete(ctx context.Context, req *dto.DeleteImageRequest) error
}

func (s *ServiceProvider) ImageService() imageService {
	if s.imageService == nil {
		s.imageService = image.NewImageService(s.MinIO(), s.MinIOConfig())
	}
	return s.imageService
}
