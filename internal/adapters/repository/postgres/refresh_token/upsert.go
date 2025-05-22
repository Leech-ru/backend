package refresh_token

import (
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/refreshtoken"
	"Leech-ru/pkg/ent/user"
	"context"
)

// Upsert update or create token by user
func (s *tokenRepo) Upsert(ctx context.Context, entity ent.RefreshToken) (*ent.RefreshToken, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	userID := entity.Edges.User.ID

	existing, err := tx.RefreshToken.
		Query().
		Where(refreshtoken.HasUserWith(user.IDEQ(userID))).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		_ = tx.Rollback()
		return nil, err
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
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}
