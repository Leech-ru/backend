package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register new user
//
// @Summary Register new user
// @Tags user
// @Accept json
// @Produce json
// @Param request body dto.RegisterUserRequest true "Registration data"
// @Success 201 {object} dto.RegisterUserResponse
// @Header 201 {string} Set-Cookie "user_auth_access_token=token; Path=/; HttpOnly; Secure; SameSite=Strict"
// @Header 201 {string} Set-Cookie "user_auth_refresh_token=token; Path=/; HttpOnly; Secure; SameSite=Strict"
// @Failure 400 {object} dto.HTTPStatus
// @Failure 409 {object} dto.HTTPStatus
// @Router /api/v1/user/register [post]
func (h *handler) Register(c echo.Context) error {
	var req dto.RegisterUserRequest
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
	resp, err := h.userService.Register(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.EmailAlreadyExist):
		return c.JSON(http.StatusConflict, dto.HTTPStatus{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})

	}

	h.cookieService.SetRefreshTokenCookie(c, resp.RefreshToken, h.jwtConfig.RefreshTokenExpires(), h.serverConfig.DevMode())
	h.cookieService.ClearAccessTokenCookie(c, h.serverConfig.DevMode())

	return c.JSON(http.StatusCreated, resp)

}
