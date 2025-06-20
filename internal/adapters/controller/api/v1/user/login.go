package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/utils/cookie"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) Login(c echo.Context) error {
	var req dto.LoginUserRequest
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

	resp, err := h.userService.Login(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.UserNotFound) || errors.Is(err, errorz.PasswordMismatch):
		return c.JSON(http.StatusUnauthorized, dto.HTTPStatus{
			Code:    http.StatusUnauthorized,
			Message: errorz.InvalidEmailOrPassword.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})

	}

	cookie.SetRefreshTokenCookie(c, resp.RefreshToken, h.jwtConfig.RefreshTokenExpires(), h.serverConfig.DevMode())
	cookie.ClearAccessTokenCookie(c, h.serverConfig.DevMode())

	return c.JSON(http.StatusOK, resp)

}
