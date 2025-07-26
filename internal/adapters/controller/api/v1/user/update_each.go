package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/types"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// UpdateEach user
//
// @Summary UpdateEach user information
// @Description Information updating the current user under which the input is executed
// @Tags user
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param id path string true "User ID (UUID)"
// @Param request body dto.UpdateEachUserRequest true "User data"
// @Success 200 {object} dto.UpdateEachUserResponse
// @Failure 400 {object} dto.HTTPStatus
// @Failure 401 {object} dto.HTTPStatus
// @Router /api/v1/user/{id} [patch]
func (h *handler) UpdateEach(c echo.Context) error {
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.UserNotFound.Error(),
		})
	}
	userRole, ok := c.Get("user_role").(types.Role)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.HTTPStatus{
			Code:    http.StatusForbidden,
			Message: errorz.PermissionDenied.Error(),
		})
	}

	var req dto.UpdateEachUserRequest
	req.ID = userID
	req.RequesterRole = userRole

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

	resp, err := h.userService.UpdateEach(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.UserNotFound):
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	case errors.Is(err, errorz.PermissionDenied):
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

	return c.JSON(http.StatusOK, resp)
}
