package news

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Delete deletes a news article by ID.
//
// @Summary      Delete news by ID
// @Description  Deletes the news article with the given UUID.
// @Tags         news
// @Security     CookieAuth
// @Param        id   path      string  true  "News ID (UUID)"  Format(uuid)
// @Success      204  "Successfully deleted"
// @Failure      400  {object}  dto.HTTPStatus "Validation error"
// @Failure      404  {object}  dto.HTTPStatus "News not found"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/news/{id} [delete]
func (h *handler) Delete(c echo.Context) error {
	id := c.Param("id")
	newsID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.NewsNotFound.Error(),
		})
	}

	var req dto.DeleteNewsRequest
	req.ID = newsID

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = h.newsService.Delete(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.NewsNotFound):
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
