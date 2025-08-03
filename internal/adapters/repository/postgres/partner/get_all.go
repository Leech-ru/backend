package partner

import (
	"Leech-ru/pkg/ent"
	"context"
)

// GetAll retrieves all partner with optional pagination.
func (s *partnersRepo) GetAll(ctx context.Context, limit, offset int) ([]*ent.Partner, error) {
	partners, err := s.client.Partner.
		Query().
		WithLinks().
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return partners, nil
}
