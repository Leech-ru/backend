package news

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Update updates a news article by ID.
//
// @Summary      Update news
// @Description  Updates news article fields by given ID.
// @Tags         news
// @Security     CookieAuth
// @Accept       json
// @Produce      json
// @Param        id      path      string                 true  "News ID (UUID)"  Format(uuid)
// @Param        request body      dto.UpdateNewsRequest  true  "Updated news fields"
// @Success      200     {object}  dto.UpdateNewsResponse
// @Failure      400     {object}  dto.HTTPStatus "Validation or binding error"
// @Failure      404     {object}  dto.HTTPStatus "News not found"
// @Failure      500     {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/news/{id} [patch]
func (h *handler) Update(c echo.Context) error {
	id := c.Param("id")
	newsID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.NewsNotFound.Error(),
		})
	}

	var req dto.UpdateNewsRequest
	req.ID = newsID

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

	resp, err := h.newsService.Update(c.Request().Context(), &req)
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

	return c.JSON(http.StatusOK, resp)
}
