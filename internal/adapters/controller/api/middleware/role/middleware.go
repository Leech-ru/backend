package role

import (
	"Leech-ru/internal/domain/types"
	"context"
	"github.com/google/uuid"
)

type userService interface {
	GetRoleByID(ctx context.Context, userID uuid.UUID) (types.Role, error)
}

type Middleware struct {
	userService userService
}

func NewRoleMiddleware(userService userService) *Middleware {
	return &Middleware{
		userService: userService,
	}
}
