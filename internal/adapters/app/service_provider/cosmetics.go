package service_provider

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/service/cosmetics"
	"context"
	"mime/multipart"
)

type cosmeticsService interface {
	Create(ctx context.Context, req *dto.CreateCosmeticsRequest, file *multipart.FileHeader) (*dto.CreateCosmeticsResponse, error)
	GetByID(ctx context.Context, req *dto.GetByIdCosmeticsRequest) (*dto.GetByIdCosmeticsResponse, error)
	GetAllByFilter(ctx context.Context, req *dto.GetAllByFilterCosmeticsRequest) (*dto.GetAllByFilterCosmeticsResponse, error)
	GetAllByFilterForAdmin(ctx context.Context, req *dto.GetAllByFilterForAdminCosmeticsRequest) (*dto.GetAllByFilterForAdminCosmeticsResponse, error)
	Update(ctx context.Context, req *dto.UpdateCosmeticsRequest, file *multipart.FileHeader) (*dto.UpdateCosmeticsResponse, error)
	Delete(ctx context.Context, req *dto.DeleteCosmeticsRequest) error
}

func (s *ServiceProvider) CosmeticsService() cosmeticsService {
	if s.cosmeticsService == nil {
		s.cosmeticsService = cosmetics.NewCosmeticsService(s.DB(), s.ImageService())
	}
	return s.cosmeticsService
}
