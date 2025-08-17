package news

import (
	"Leech-ru/internal/domain/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetAllByFilter returns a list of news with optional preview length.
//
// @Summary      Get news by filters
// @Description  Retrieves a list of news with pagination and content preview length option.
// @Tags         news
// @Accept       json
// @Produce      json
// @Param        limit           query     int  false  "Max number of items"          minimum(1) maximum(100)
// @Param        offset          query     int  false  "Offset for pagination"        minimum(0)
// @Param        preview_length  query     int  false  "Number of characters for content preview" minimum(1) maximum(500)
// @Success      200  {array}   dto.NewsListItem
// @Failure      400  {object}  dto.HTTPStatus "Invalid query parameters"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/news [get]
func (h *handler) GetAllByFilter(c echo.Context) error {
	var req dto.GetAllByFilterNewsRequest

	if err := h.formDecoder.Decode(&req, c.QueryParams()); err != nil {
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

	resp, err := h.newsService.GetAllByFilter(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
