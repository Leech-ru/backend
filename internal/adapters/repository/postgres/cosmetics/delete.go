package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Delete removes a cosmetics by ID
func (s *cosmeticsRepo) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.client.Cosmetics.
		DeleteOneID(id).
		Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		return errorz.CosmeticsNotFound
	case err != nil:
		return fmt.Errorf("failed to query db: %w", err)
	}

	return nil
}
