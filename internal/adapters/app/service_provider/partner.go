package service_provider

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/service/partner"
	"context"
)

type partnerService interface {
	Create(ctx context.Context, req *dto.CreatePartnerRequest) (*dto.CreatePartnerResponse, error)
	GetByID(ctx context.Context, req *dto.GetByIdPartnerRequest) (*dto.GetByIdPartnerResponse, error)
	GetAll(ctx context.Context, req *dto.GetAllPartnerRequest) (*dto.GetAllPartnerResponse, error)
	Update(ctx context.Context, req *dto.UpdatePartnerRequest) (*dto.UpdatePartnerResponse, error)
	Delete(ctx context.Context, req *dto.DeletePartnerRequest) error
}

func (s *ServiceProvider) PartnerService() partnerService {
	if s.partnerService == nil {
		s.partnerService = partner.NewPartnerService(s.DB())
	}
	return s.partnerService
}
