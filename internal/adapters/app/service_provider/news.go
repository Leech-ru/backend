package service_provider

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/service/news"
	"context"
)

type newsService interface {
	Create(ctx context.Context, req *dto.CreateNewsRequest) (*dto.CreateNewsResponse, error)
	GetByID(ctx context.Context, req *dto.GetByIdNewsRequest) (*dto.GetByIdNewsResponse, error)
	GetAllByFilter(ctx context.Context, req *dto.GetAllByFilterNewsRequest) (*dto.GetAllByFilterNewsResponse, error)
	GetAllByFilterForAdmins(ctx context.Context, req *dto.GetAllByFilterForAdminsNewsRequest) (*dto.GetAllByFilterForAdminsNewsResponse, error)
	Update(ctx context.Context, req *dto.UpdateNewsRequest) (*dto.UpdateNewsResponse, error)
	Delete(ctx context.Context, req *dto.DeleteNewsRequest) error
}

func (s *ServiceProvider) NewsService() newsService {
	if s.newsService == nil {
		s.newsService = news.NewNewsService(s.DB(), s.ImageService())
	}
	return s.newsService
}
