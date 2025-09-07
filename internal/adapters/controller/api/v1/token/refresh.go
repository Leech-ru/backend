package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Refresh refreshes access and refresh tokens.
//
// @Summary      Refresh tokens
// @Description  Refreshes access and refresh tokens using valid refresh token from cookies
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        goyda    query     string                  true   "Must be 'true'"  Enums(true)
// @Success      204      "Successful updated"
// @Failure      400      {object}  dto.HTTPStatus          "Invalid request or validation error"
// @Failure      401      {object}  dto.HTTPStatus          "Unauthorized - invalid/missing refresh token"
// @Failure      500      {object}  dto.HTTPStatus          "Internal server error"
// @Router       /api/v1/auth/refresh [post]
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

	token, err := h.cookieService.ReadRefreshTokenCookie(c.Request())
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
	refreshToken, err := h.tokenService.GenerateRefreshToken(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	accessToken, err := h.tokenService.GenerateAccessToken(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	h.cookieService.SetRefreshTokenCookie(c, refreshToken, h.jwtConfig.RefreshTokenExpires(), h.serverConfig.DevMode())
	h.cookieService.SetAccessTokenCookie(c, accessToken, h.jwtConfig.AccessTokenExpires(), h.serverConfig.DevMode())

	return c.NoContent(http.StatusNoContent)
}
