package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/partner"
	"context"
	"github.com/google/uuid"
)

// GetById retrieves a partner by ID
func (s *partnersRepo) GetById(ctx context.Context, id uuid.UUID) (*ent.Partner, error) {
	p, err := s.client.Partner.
		Query().
		WithLinks().
		Where(partner.ID(id)).
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.PartnerNotFound
	case err != nil:
		return nil, err
	}

	return p, nil
}
