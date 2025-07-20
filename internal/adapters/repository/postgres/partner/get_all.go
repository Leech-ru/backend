package partner

import (
	"Leech-ru/pkg/ent"
	"context"
)

// GetAll retrieves all partner with optional pagination.
func (s *partnersRepo) GetAll(ctx context.Context, limit, offset int) ([]*ent.Partner, error) {
	query := s.client.Partner.Query()
	cosmetics, err := query.
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return cosmetics, nil
}
