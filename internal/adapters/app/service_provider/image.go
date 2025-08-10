package service_provider

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/service/image"
	"context"
	"github.com/google/uuid"
)

type imageService interface {
	Create(ctx context.Context, req *dto.CreateImageRequest) (*dto.CreateImageResponse, error)
	GetById(ctx context.Context, req *dto.GetByIdImageRequest) (*dto.GetByIdImageResponse, error)
	Delete(ctx context.Context, req *dto.DeleteImageRequest) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}

func (s *ServiceProvider) ImageService() imageService {
	if s.imageService == nil {
		s.imageService = image.NewImageService(s.MinIO(), s.MinIOConfig())
	}
	return s.imageService
}
