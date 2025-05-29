package user

import (
	"Leech-ru/internal/adapters/controller/api/middleware/auth"
	"Leech-ru/internal/adapters/controller/api/middleware/role"
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/types"
	"context"
	"github.com/go-playground/form"
	"github.com/labstack/echo/v4"
	"time"
)

type userService interface {
	Register(ctx context.Context, req *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error)
	Login(ctx context.Context, req *dto.LoginUserRequest) (*dto.LoginUserResponse, error)
	ChangePassword(ctx context.Context, req *dto.ChangePasswordRequest) (*dto.ChangePasswordResponse, error)
	GetByID(ctx context.Context, req *dto.GetUserRequest) (*dto.GetUserResponse, error)
	GetAll(ctx context.Context, req *dto.GetAllUsersRequest) (*dto.GetAllUsersResponse, error)
	Update(ctx context.Context, req *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
	Logout(ctx context.Context, req *dto.LogoutRequest) error
}

type jwtConfig interface {
	RefreshTokenExpires() time.Duration
}

type serverConfig interface {
	DevMode() bool
}

type handler struct {
	userService    userService
	jwtConfig      jwtConfig
	serverConfig   serverConfig
	authMiddleware *auth.Middleware
	roleMiddleware *role.Middleware
	validator      *validator.Validator
	formDecoder    *form.Decoder
}

func NewHandler(
	userService userService,
	jwtConfig jwtConfig,
	serverConfig serverConfig,
	authMiddleware *auth.Middleware,
	roleMiddleware *role.Middleware,
	validator *validator.Validator,
	formDecoder *form.Decoder,

) *handler {
	return &handler{
		userService:    userService,
		jwtConfig:      jwtConfig,
		serverConfig:   serverConfig,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
		validator:      validator,
		formDecoder:    formDecoder,
	}
}

func (h *handler) Setup(router *echo.Group) {
	router.POST("/user/register", h.Register)
	router.POST("/user/login", h.Login)
	router.POST("/user/password", h.ChangePassword, h.authMiddleware.RequireAuth)
	router.GET("/user/all", h.GetAll, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.Moderator))
	router.GET("/user", h.GetMe, h.authMiddleware.RequireAuth)
	router.PATCH("/user", h.Update, h.authMiddleware.RequireAuth)
	router.POST("/user/logout", h.Logout, h.authMiddleware.RequireAuth)
}
