package service_provider

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/service/cosmetics"
	"context"
)

type cosmeticsService interface {
	Create(ctx context.Context, req *dto.CreateCosmeticsRequest) (*dto.CreateCosmeticsResponse, error)
	GetByID(ctx context.Context, req *dto.GetCosmeticsRequest) (*dto.GetCosmeticsResponse, error)
	GetAllByFilter(ctx context.Context, req *dto.GetAllByFilterCosmeticsRequest) (*dto.GetAllByFilterCosmeticsResponse, error)
	Update(ctx context.Context, req *dto.UpdateCosmeticsRequest) (*dto.UpdateCosmeticsResponse, error)
	Delete(ctx context.Context, req *dto.DeleteCosmeticsRequest) error
}

func (s *ServiceProvider) CosmeticsService() cosmeticsService {
	if s.cosmeticsService == nil {
		s.cosmeticsService = cosmetics.NewCosmeticsService(s.DB())
	}
	return s.cosmeticsService
}
