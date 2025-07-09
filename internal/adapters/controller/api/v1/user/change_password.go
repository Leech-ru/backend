package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/utils/cookie"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ChangePassword Change user password
//
// @Summary Change user password
// @Tags user
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body dto.ChangePasswordRequest true "Password change data"
// @Success 204
// @Header 204 {string} Set-Cookie "user_auth_access_token=token; Path=/; HttpOnly; Secure; SameSite=Strict"
// @Header 204 {string} Set-Cookie "user_auth_refresh_token=token; Path=/; HttpOnly; Secure; SameSite=Strict"
// @Failure 400 {object} dto.HTTPStatus
// @Failure 403 {object} dto.HTTPStatus
// @Router /api/v1/user/password [post]
func (h *handler) ChangePassword(c echo.Context) error {
	var req dto.ChangePasswordRequest
	userID, _ := c.Get("user_id").(uuid.UUID)
	req.ID = userID

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

	resp, err := h.userService.ChangePassword(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.UserNotFound):
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.UserNotFound.Error(),
		})
	case errors.Is(err, errorz.PasswordMismatch):
		return c.JSON(http.StatusForbidden, dto.HTTPStatus{
			Code:    http.StatusForbidden,
			Message: errorz.PasswordMismatch.Error(),
		})
	case errors.Is(err, errorz.PasswordsCoincidence):
		return c.JSON(http.StatusConflict, dto.HTTPStatus{
			Code:    http.StatusConflict,
			Message: errorz.PasswordsCoincidence.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})

	}

	cookie.SetRefreshTokenCookie(c, resp.RefreshToken, h.jwtConfig.RefreshTokenExpires(), h.serverConfig.DevMode())
	cookie.ClearAccessTokenCookie(c, h.serverConfig.DevMode())

	return c.NoContent(http.StatusNoContent)

}
