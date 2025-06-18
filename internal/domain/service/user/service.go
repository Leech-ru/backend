package user

import (
	"Leech-ru/internal/adapters/repository/postgres/user"
	"Leech-ru/internal/domain/types"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

type userRepo interface {
	Create(ctx context.Context, userEntity ent.User) (*ent.User, error)
	GetById(ctx context.Context, id uuid.UUID) (*ent.User, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	GetAllByFilter(
		ctx context.Context,
		limit, offset int,
		role *types.Role,
		namePrefix, surnamePrefix, emailPrefix *string,
	) ([]*ent.User, error)
	Update(ctx context.Context, userEntity ent.User) (*ent.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type tokenService interface {
	RevokeAccessToken(ctx context.Context, userID uuid.UUID) error
	GenerateRefreshToken(ctx context.Context, userID uuid.UUID) (string, error)
	LogoutAllSessions(ctx context.Context, userID uuid.UUID) error
}

type serverConfig interface {
	DevMode() bool
}

// TODO users search
type userService struct {
	userRepo     userRepo
	tokenService tokenService
	serverConfig serverConfig
}

func NewUserService(client *ent.Client, tokenService tokenService, serverConfig serverConfig) *userService {
	return &userService{
		userRepo:     user.NewUserRepo(client),
		tokenService: tokenService,
		serverConfig: serverConfig,
	}
}
