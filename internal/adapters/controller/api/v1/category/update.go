package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Update updates a category product by ID.
//
// @Summary      Update category
// @Description  Updates category product fields by given ID.
// @Tags         category
// @Security CookieAuth
// @Accept       json
// @Produce      json
// @Param        id      path      string                    true  "Cosmetic ID (UUID)"  Format(uuid)
// @Param        request body      dto.UpdateCategoryRequest true  "Updated category fields"
// @Success      200     {object}  dto.UpdateCategoryResponse
// @Failure      400     {object}  dto.HTTPStatus "Validation or binding error"
// @Failure      404     {object}  dto.HTTPStatus "Cosmetic not found"
// @Failure      500     {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/category/{id} [patch]
func (h *handler) Update(c echo.Context) error {
	id := c.Param("id")
	cosmeticID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.CategoryNotFound.Error(),
		})
	}
	var req dto.UpdateCategoryRequest
	req.ID = cosmeticID

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	resp, err := h.categoryService.Update(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.CategoryNotFound):
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	case errors.Is(err, errorz.ImageNotFound):
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
