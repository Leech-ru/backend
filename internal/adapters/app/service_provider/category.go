package service_provider

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/service/category"
	"context"
)

type categoryService interface {
	Create(ctx context.Context, req *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, error)
	GetByID(ctx context.Context, req *dto.GetByIdCategoryRequest) (*dto.GetByIdCategoryResponse, error)
	GetAll(ctx context.Context, req *dto.GetAllCategoriesRequest) (*dto.GetAllCategoriesResponse, error)
	Update(ctx context.Context, req *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, error)
	Delete(ctx context.Context, req *dto.DeleteCategoryRequest) error
}

func (s *ServiceProvider) CategoryService() categoryService {
	if s.categoryService == nil {
		s.categoryService = category.NewCategoryService(s.DB())
	}
	return s.categoryService
}
