package access_token

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// Set stores the jti for the given user_id with a TTL.
func (r *repo) Set(ctx context.Context, userID uuid.UUID, value string, exp time.Time) error {
	key := fmt.Sprintf(KeyTemplate, userID.String())
	ttl := time.Until(exp)
	if ttl <= 0 {
		ttl = time.Minute
	}
	return r.client.SetEx(ctx, key, value, ttl).Err()
}
