package service_provider

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/service/info"
	"context"
)

type infoService interface {
	Get(_ context.Context) (*dto.GetInfoResponse, error)
	Update(_ context.Context, req *dto.UpdateInfoRequest) (*dto.UpdateInfoResponse, error)
}

func (s *ServiceProvider) InfoService() infoService {
	if s.infoService == nil {
		s.infoService = info.NewService(s.JsonInfoConfig())
	}

	return s.infoService
}
