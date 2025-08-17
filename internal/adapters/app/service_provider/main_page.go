package service_provider

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/service/main_page"
	"context"
)

type mainPageService interface {
	Create(ctx context.Context, req *dto.CreateMainPageRequest) (*dto.CreateMainPageResponse, error)
	GetByID(ctx context.Context, req *dto.GetByIdMainPageRequest) (*dto.GetByIdMainPageResponse, error)
	GetAll(ctx context.Context, req *dto.GetAllMainPageRequest) (dto.GetAllMainPageResponse, error)
	Update(ctx context.Context, req *dto.UpdateMainPageRequest) (*dto.UpdateMainPageResponse, error)
	Delete(ctx context.Context, req *dto.DeleteMainPageRequest) error
}

func (s *ServiceProvider) MainPageService() mainPageService {
	if s.mainPageService == nil {
		s.mainPageService = mainpage.NewMainPageService(s.DB(), s.imageService)
	}
	return s.mainPageService
}
