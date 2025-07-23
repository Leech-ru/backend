package partner

import (
	"Leech-ru/internal/adapters/repository/postgres/partner"
	"Leech-ru/pkg/ent"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type partnerRepo interface {
	Create(ctx context.Context, entity ent.Partner) (*ent.Partner, error)
	GetById(ctx context.Context, id uuid.UUID) (*ent.Partner, error)
	GetAll(ctx context.Context, limit, offset int) ([]*ent.Partner, error)
	Update(ctx context.Context, entity ent.Partner) (*ent.Partner, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type partnerService struct {
	partnerRepo partnerRepo
}

func NewPartnerService(entClient *ent.Client) *partnerService {
	return &partnerService{
		partnerRepo: partner.NewPartnersRepo(entClient),
	}
}
