package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Delete removes a main page content by ID
func (s *mainPageRepo) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.client.MainPage.
		DeleteOneID(id).
		Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		return errorz.MainPageNotFound
	case err != nil:
		return fmt.Errorf("failed to query db: %w", err)
	}

	return nil
}
