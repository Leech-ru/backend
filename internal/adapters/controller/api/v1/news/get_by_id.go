package news

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetById returns a news article by its ID.
//
// @Summary      Get news by ID
// @Description  Retrieves a news article using its UUID.
// @Tags         news
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "News ID (UUID)"  Format(uuid)
// @Success      200  {object}  dto.GetByIdNewsResponse
// @Failure      400  {object}  dto.HTTPStatus "Validation error"
// @Failure      404  {object}  dto.HTTPStatus "News not found"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/news/{id} [get]
func (h *handler) GetById(c echo.Context) error {
	id := c.Param("id")
	newsID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.NewsNotFound.Error(),
		})
	}

	var req dto.GetByIdNewsRequest
	req.ID = newsID

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	resp, err := h.newsService.GetByID(c.Request().Context(), &req)
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
