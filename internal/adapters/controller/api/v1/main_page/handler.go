package mainpage

import (
	"Leech-ru/internal/adapters/controller/api/middleware/auth"
	"Leech-ru/internal/adapters/controller/api/middleware/role"
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/types"
	"context"

	"github.com/go-playground/form"
	"github.com/labstack/echo/v4"
)

type mainPageService interface {
	Create(ctx context.Context, req *dto.CreateMainPageRequest) (*dto.CreateMainPageResponse, error)
	GetByID(ctx context.Context, req *dto.GetByIdMainPageRequest) (*dto.GetByIdMainPageResponse, error)
	GetAll(ctx context.Context, req *dto.GetAllMainPageRequest) (dto.GetAllMainPageResponse, error)
	Update(ctx context.Context, req *dto.UpdateMainPageRequest) (*dto.UpdateMainPageResponse, error)
	Delete(ctx context.Context, req *dto.DeleteMainPageRequest) error
}

type handler struct {
	mainPageService mainPageService
	authMiddleware  *auth.Middleware
	roleMiddleware  *role.Middleware
	validator       *validator.Validator
	formDecoder     *form.Decoder
}

func NewHandler(
	mainPageService mainPageService,
	authMiddleware *auth.Middleware,
	roleMiddleware *role.Middleware,
	validator *validator.Validator,
	formDecoder *form.Decoder,
) *handler {
	return &handler{
		mainPageService: mainPageService,
		authMiddleware:  authMiddleware,
		roleMiddleware:  roleMiddleware,
		validator:       validator,
		formDecoder:     formDecoder,
	}
}

func (h *handler) Setup(router *echo.Group) {
	router.POST("/main", h.Create, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
	router.PATCH("/main/:id", h.Update, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
	router.DELETE("/main/:id", h.Delete, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
	router.GET("/main", h.GetAll)
	router.GET("/main/:id", h.GetById)
}
