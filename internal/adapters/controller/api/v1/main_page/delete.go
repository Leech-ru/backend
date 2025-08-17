package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Delete deletes a main page content by ID.
//
// @Summary      Delete content by ID
// @Description  Deletes the main page content with the given UUID.
// @Tags         main
// @Security CookieAuth
// @Param        id   path      string  true  "Content ID (UUID)"  Format(uuid)
// @Success      204  "Successfully deleted"
// @Failure      400  {object}  dto.HTTPStatus "Validation error"
// @Failure      404  {object}  dto.HTTPStatus "Content not found"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/main/{id} [delete]
func (h *handler) Delete(c echo.Context) error {
	id := c.Param("id")
	pageID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.MainPageNotFound.Error(),
		})
	}

	var req dto.DeleteMainPageRequest
	req.ID = pageID

	if err = h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = h.mainPageService.Delete(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.MainPageNotFound):
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
