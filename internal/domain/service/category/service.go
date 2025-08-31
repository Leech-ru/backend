package category

import (
	"Leech-ru/internal/adapters/repository/postgres/category"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

type categoryRepo interface {
	Create(ctx context.Context, entity ent.Category) (*ent.Category, error)
	GetById(ctx context.Context, id uuid.UUID) (*ent.Category, error)
	GetAll(ctx context.Context, limit, offset int) ([]*ent.Category, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, entity ent.Category) (*ent.Category, error)
}

type imageService interface {
	Create(ctx context.Context, req *dto.CreateImageRequest) (*dto.CreateImageResponse, error)
	GetById(ctx context.Context, req *dto.GetByIdImageRequest) (*dto.GetByIdImageResponse, error)
	Delete(ctx context.Context, req *dto.DeleteImageRequest) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}

type categoryService struct {
	categoryRepo categoryRepo
	imageService imageService
}

func NewCategoryService(entClient *ent.Client, imageService imageService) *categoryService {
	return &categoryService{
		categoryRepo: category.NewCategoryRepo(entClient),
		imageService: imageService,
	}
}
