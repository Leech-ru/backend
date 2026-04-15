package partner

import (
	"Leech-ru/pkg/ent"
	"context"
	"fmt"
)

// GetAll retrieves all partner with optional pagination.
func (s *partnersRepo) GetAll(ctx context.Context, limit, offset int) ([]*ent.Partner, int, error) {
	baseQuery := s.client.Partner.Query().WithLinks()

	totalItems, err := baseQuery.Clone().Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count partners in db: %w", err)
	}

	partners, err := baseQuery.
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to query db: %w", err)
	}

	return partners, totalItems, nil
}
