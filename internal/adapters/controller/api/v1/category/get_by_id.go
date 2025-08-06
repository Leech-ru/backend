package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetById returns category product by its ID.
//
// @Summary      Get category by ID
// @Description  Retrieves a category product using its UUID.
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Cosmetic ID (UUID)"  Format(uuid)
// @Success      200  {object}  dto.GetByIdCategoryResponse
// @Failure      400  {object}  dto.HTTPStatus "Validation error"
// @Failure      404  {object}  dto.HTTPStatus "Cosmetic not found"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/category/{id} [get]
func (h *handler) GetById(c echo.Context) error {
	id := c.Param("id")
	cosmeticID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.CategoryNotFound.Error(),
		})
	}
	var req dto.GetByIdCategoryRequest
	req.ID = cosmeticID

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	resp, err := h.categoryService.GetByID(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.CategoryNotFound):
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
