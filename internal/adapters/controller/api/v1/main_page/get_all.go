package mainpage

import (
	"Leech-ru/internal/domain/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAll returns a list of main page contents with pagination
//
// @Summary      Get all contents
// @Description  Retrieves a list of main page contents with limit and offset
// @Tags         main
// @Accept       json
// @Produce      json
// @Param        limit        query     int     false  "Max number of items"          minimum(1) maximum(100)
// @Param        offset       query     int     false  "Offset for pagination"        minimum(0)
// @Success      200  {array}   dto.MainPage
// @Failure      400  {object}  dto.HTTPStatus "Invalid query parameters"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/main [get]
func (h *handler) GetAll(c echo.Context) error {
	var req dto.GetAllMainPageRequest

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

	resp, err := h.mainPageService.GetAll(c.Request().Context(), &req)
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
