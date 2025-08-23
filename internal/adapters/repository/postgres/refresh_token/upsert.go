package refresh_token

import (
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/refreshtoken"
	"Leech-ru/pkg/ent/user"
	"context"
	"fmt"
)

// Upsert update or create token by user
func (s *tokenRepo) Upsert(ctx context.Context, entity ent.RefreshToken) (*ent.RefreshToken, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin tx: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
		}
	}()

	userID := entity.Edges.User.ID

	existing, err := tx.RefreshToken.
		Query().
		Where(refreshtoken.HasUserWith(user.IDEQ(userID))).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		_ = tx.Rollback()
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	var result *ent.RefreshToken

	if existing != nil {
		result, err = tx.RefreshToken.
			UpdateOneID(existing.ID).
			SetJti(entity.Jti).
			Save(ctx)
	} else {
		result, err = tx.RefreshToken.
			Create().
			SetJti(entity.Jti).
			SetUserID(userID).
			Save(ctx)
	}

	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit tx: %w", err)
	}

	return result, nil
}
