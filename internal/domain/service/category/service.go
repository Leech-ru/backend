package category

import (
	"Leech-ru/internal/adapters/repository/postgres/category"
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

type categoryService struct {
	categoryRepo categoryRepo
}

func NewCategoryService(entClient *ent.Client) *categoryService {
	return &categoryService{
		categoryRepo: category.NewCategoryRepo(entClient),
	}
}
