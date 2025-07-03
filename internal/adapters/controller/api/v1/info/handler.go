package info

import (
	"Leech-ru/internal/adapters/controller/api/middleware/auth"
	"Leech-ru/internal/adapters/controller/api/middleware/role"
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/types"
	"context"
	"github.com/labstack/echo/v4"
)

type infoService interface {
	Get(_ context.Context) (*dto.GetInfoResponse, error)
	Update(_ context.Context, req *dto.UpdateInfoRequest) (*dto.UpdateInfoResponse, error)
}

type handler struct {
	infoService    infoService
	authMiddleware *auth.Middleware
	roleMiddleware *role.Middleware
	validator      *validator.Validator
}

func NewHandler(
	infoService infoService,
	authMiddleware *auth.Middleware,
	roleMiddleware *role.Middleware,
	validator *validator.Validator,

) *handler {
	return &handler{
		infoService:    infoService,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
		validator:      validator,
	}
}

func (h *handler) Setup(router *echo.Group) {
	router.GET("/info/corporation", h.GetInfo)
	router.PATCH("/info/corporation", h.UpdateInfo, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
}
