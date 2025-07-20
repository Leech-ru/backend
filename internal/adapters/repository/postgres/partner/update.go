package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Update updates an existing partner
func (s *partnersRepo) Update(ctx context.Context, entity ent.Partner) (*ent.Partner, error) {
	updated, err := s.client.Partner.
		UpdateOneID(entity.ID).
		SetName(entity.Name).SetNillableDescription(entity.Description).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.PartnerNotFound
	case err != nil:
		return nil, err
	}

	return updated, nil
}
