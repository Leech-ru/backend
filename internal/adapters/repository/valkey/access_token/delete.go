package access_token

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// Delete удаляет jti по заданному userID.
func (r *repo) Delete(ctx context.Context, userID uuid.UUID) error {
	key := fmt.Sprintf(KeyTemplate, userID.String())
	return r.client.Del(ctx, key).Err()
}
