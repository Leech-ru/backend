package category

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

type categoryService interface {
	Create(ctx context.Context, req *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, error)
	GetByID(ctx context.Context, req *dto.GetByIdCategoryRequest) (*dto.GetByIdCategoryResponse, error)
	GetAll(ctx context.Context, req *dto.GetAllCategoriesRequest) (*dto.GetAllCategoriesResponse, error)
	Update(ctx context.Context, req *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, error)
	Delete(ctx context.Context, req *dto.DeleteCategoryRequest) error
}

type handler struct {
	categoryService categoryService
	authMiddleware  *auth.Middleware
	roleMiddleware  *role.Middleware
	validator       *validator.Validator
	formDecoder     *form.Decoder
}

func NewHandler(
	categoryService categoryService,
	authMiddleware *auth.Middleware,
	roleMiddleware *role.Middleware,
	validator *validator.Validator,
	formDecoder *form.Decoder,

) *handler {
	return &handler{
		categoryService: categoryService,
		authMiddleware:  authMiddleware,
		roleMiddleware:  roleMiddleware,
		validator:       validator,
		formDecoder:     formDecoder,
	}
}

func (h *handler) Setup(router *echo.Group) {
	router.GET("/category", h.GetAll)
	router.GET("/category/:id", h.GetById)
	router.POST("/category", h.Create, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
	router.PATCH("/category/:id", h.Update, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
	router.DELETE("/category/:id", h.Delete, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
}
