package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Delete deletes a category product by ID.
//
// @Summary      Delete category by ID
// @Description  Deletes the category product with the given UUID.
// @Tags         category
// @Security CookieAuth
// @Param        id   path      string  true  "Category ID (UUID)"  Format(uuid)
// @Success      204  "Successfully deleted"
// @Failure      400  {object}  dto.HTTPStatus "Validation error"
// @Failure      404  {object}  dto.HTTPStatus "Category not found"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/category/{id} [delete]
func (h *handler) Delete(c echo.Context) error {
	id := c.Param("id")
	cosmeticID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.CategoryNotFound.Error(),
		})
	}
	var req dto.DeleteCategoryRequest
	req.ID = cosmeticID

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = h.categoryService.Delete(c.Request().Context(), &req)
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

	return c.NoContent(http.StatusNoContent)
}
