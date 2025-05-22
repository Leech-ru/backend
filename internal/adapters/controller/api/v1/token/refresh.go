package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/utils/cookie"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) Refresh(c echo.Context) error {
	var req dto.RefreshRequest
	if err := h.formDecoder.Decode(&req, c.QueryParams()); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: "GOOOOOOOOOOOOOOOYDA IS REQUIRED!",
		})
	}

	token, err := cookie.ReadRefreshTokenCookie(c.Request())
	switch {
	case errors.Is(err, errorz.NoCookie):
		return c.JSON(http.StatusUnauthorized, dto.HTTPStatus{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	userID, err := h.tokenService.ParseRefreshToken(c.Request().Context(), token)
	switch {
	case errors.Is(err, errorz.Unauthorized):
		return c.JSON(http.StatusUnauthorized, dto.HTTPStatus{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
	case errors.Is(err, errorz.InvalidToken):
		return c.JSON(http.StatusUnauthorized, dto.HTTPStatus{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	newToken, err := h.tokenService.GenerateAccessToken(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	createdCookie := cookie.NewAccessTokenCookie(newToken, h.jwtConfig.AccessTokenExpires(), true)
	c.SetCookie(createdCookie)

	return c.NoContent(http.StatusNoContent)
}
