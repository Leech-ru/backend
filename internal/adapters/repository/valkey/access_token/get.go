package access_token

import (
	"Leech-ru/internal/domain/common/errorz"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// Get returns jti for user_id matches the provided jti.
func (r *repo) Get(ctx context.Context, userID uuid.UUID) (string, error) {
	key := fmt.Sprintf(KeyTemplate, userID)
	stored, err := r.client.Get(ctx, key).Result()

	switch {
	case errors.Is(err, redis.Nil):
		return "", errorz.TokenNotFound
	case err != nil:
		return "", err
	}

	return stored, nil
}
