package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Login User login
//
// @Summary Authenticate user
// @Tags user
// @Accept json
// @Produce json
// @Param request body dto.LoginUserRequest true "Login credentials"
// @Success 200 {object} dto.LoginUserResponse
// @Header 200 {string} Set-Cookie "user_auth_refresh_token=token; Path=/; HttpOnly; Secure; SameSite=Strict"
// @Failure 400 {object} dto.HTTPStatus
// @Failure 401 {object} dto.HTTPStatus
// @Router /api/v1/user/login [post]
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

	h.cookieService.SetRefreshTokenCookie(c, resp.RefreshToken, h.jwtConfig.RefreshTokenExpires(), h.serverConfig.DevMode())
	h.cookieService.ClearAccessTokenCookie(c, h.serverConfig.DevMode())

	return c.JSON(http.StatusOK, resp)

}
