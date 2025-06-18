package cosmetics

import (
	"Leech-ru/internal/adapters/repository/postgres/cosmetics"
	"Leech-ru/internal/domain/types"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

type cosmeticsRepo interface {
	Create(ctx context.Context, entity ent.Cosmetics) (*ent.Cosmetics, error)
	GetById(ctx context.Context, id uuid.UUID) (*ent.Cosmetics, error)
	GetAll(ctx context.Context, limit, offset int) ([]*ent.Cosmetics, error)
	GetAllByFilter(ctx context.Context, limit, offset int, category *types.Category, titlePrefix *string, volume *int) ([]*ent.Cosmetics, error)
	Update(ctx context.Context, userEntity ent.Cosmetics) (*ent.Cosmetics, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type cosmeticsService struct {
	cosmeticsRepo cosmeticsRepo
}

func NewCosmeticsService(entClient *ent.Client) *cosmeticsService {
	return &cosmeticsService{
		cosmeticsRepo: cosmetics.NewCosmeticsRepo(entClient),
	}
}
