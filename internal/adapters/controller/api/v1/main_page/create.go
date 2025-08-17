package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Create handles the creation of new main page content
//
//	@Summary		Create a new main page content
//	@Description	Creates a new page content with provided details
//	@Tags			main
//	@Security		CookieAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.CreateMainPageRequest	true	"Content data"
//	@Success		201		{object}	dto.CreateMainPageResponse
//	@Failure		400		{object}	dto.HTTPStatus	"Invalid request body or validation error"
//	@Failure		409		{object}	dto.HTTPStatus	"Conflict: invalid content format"
//	@Failure		500		{object}	dto.HTTPStatus	"Internal server error"
//	@Router			/api/v1/main [post]
func (h *handler) Create(c echo.Context) error {
	var req dto.CreateMainPageRequest
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

	resp, err := h.mainPageService.Create(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.InvalidMainPageFormat):
		return c.JSON(http.StatusConflict, dto.HTTPStatus{
			Code:    http.StatusConflict,
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

	return c.JSON(http.StatusCreated, resp)
}
