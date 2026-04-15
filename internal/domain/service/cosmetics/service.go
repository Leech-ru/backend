package cosmetics

import (
	"Leech-ru/internal/adapters/repository/postgres/cosmetics"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

type cosmeticsRepo interface {
	Create(ctx context.Context, entity ent.Cosmetics) (*ent.Cosmetics, error)
	GetById(ctx context.Context, id uuid.UUID) (*ent.Cosmetics, error)
	GetAllByFilter(ctx context.Context, limit, offset int, categoryID *uuid.UUID, titlePrefix *string, volume *int, isHidden *bool) ([]*ent.Cosmetics, int, error)
	Update(ctx context.Context, userEntity ent.Cosmetics) (*ent.Cosmetics, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type imageService interface {
	Create(ctx context.Context, req *dto.CreateImageRequest) (*dto.CreateImageResponse, error)
	GetById(ctx context.Context, req *dto.GetByIdImageRequest) (*dto.GetByIdImageResponse, error)
	Delete(ctx context.Context, req *dto.DeleteImageRequest) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}

type cosmeticsService struct {
	cosmeticsRepo cosmeticsRepo
	imageService  imageService
}

func NewCosmeticsService(entClient *ent.Client, imageService imageService) *cosmeticsService {
	return &cosmeticsService{
		cosmeticsRepo: cosmetics.NewCosmeticsRepo(entClient),
		imageService:  imageService,
	}
}
