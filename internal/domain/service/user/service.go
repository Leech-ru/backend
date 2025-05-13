package user

import (
	"Leech-ru/internal/adapters/repository/postgres/user"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

type userRepo interface {
	Create(ctx context.Context, userEntity ent.User) (*ent.User, error)
	GetById(ctx context.Context, id uuid.UUID) (*ent.User, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	GetAll(ctx context.Context, limit, offset int) ([]*ent.User, error)
	Update(ctx context.Context, userEntity ent.User) (*ent.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type tokenService interface {
	NewToken(ctx context.Context, id uuid.UUID) (string, error)
	UpdateToken(ctx context.Context, id uuid.UUID) (string, error)
	GetIDByToken(ctx context.Context, token string) (uuid.UUID, error)
	GetTokenByID(ctx context.Context, id uuid.UUID) (string, error)
}

type userService struct {
	userRepo     userRepo
	tokenService tokenService
}

func NewUserService(client *ent.Client, tokenService tokenService) *userService {
	return &userService{
		userRepo:     user.NewUserRepo(client),
		tokenService: tokenService,
	}
}
