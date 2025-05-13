package user

import (
	"Leech-ru/internal/adapters/repository/postgres/user"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

type UserRepo interface {
	Create(ctx context.Context, userEntity ent.User) (*ent.User, error)
	GetById(ctx context.Context, id uuid.UUID) (*ent.User, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	GetAll(ctx context.Context, limit, offset int) ([]*ent.User, error)
	Update(ctx context.Context, userEntity ent.User) (*ent.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type userService struct {
	userRepo UserRepo
}

func NewUserService(client *ent.Client) *userService {
	return &userService{
		userRepo: user.NewUserRepo(client),
	}
}
