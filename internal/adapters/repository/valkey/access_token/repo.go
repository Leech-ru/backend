package access_token

import (
	"github.com/redis/go-redis/v9"
)

const (
	KeyTemplate = "user_jti:%s"
)

type repo struct {
	client *redis.Client
}

// NewTokenRepo creates new valkey repo.
func NewTokenRepo(client *redis.Client) *repo {
	return &repo{client: client}
}
