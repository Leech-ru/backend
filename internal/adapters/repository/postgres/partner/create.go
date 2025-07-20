package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Create creates a new partner in the database
func (s *partnersRepo) Create(ctx context.Context, entity ent.Partner) (*ent.Partner, error) {
	created, err := s.client.Partner.
		Create().
		SetName(entity.Name).
		SetNillableDescription(entity.Description).
		Save(ctx)

	switch {
	case ent.IsConstraintError(err):
		return nil, errorz.InvalidPartnerFormat
	case err != nil:
		return nil, err
	}
	return created, nil
}
