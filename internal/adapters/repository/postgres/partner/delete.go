package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Delete removes a partner by ID
func (s *partnersRepo) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.client.Partner.
		DeleteOneID(id).
		Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		return errorz.PartnerNotFound
	case err != nil:
		return fmt.Errorf("failed to query db: %w", err)
	}

	return nil
}
