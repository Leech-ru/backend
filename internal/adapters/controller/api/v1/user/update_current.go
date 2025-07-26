package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// UpdateCurrent user
//
// @Summary UpdateCurrent user information
// @Description Information updating the current user under which the input is executed
// @Tags user
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param request body dto.UpdateCurrentUserRequest true "User data"
// @Success 200 {object} dto.UpdateCurrentUserResponse
// @Failure 400 {object} dto.HTTPStatus
// @Failure 401 {object} dto.HTTPStatus
// @Router /api/v1/user [patch]
func (h *handler) UpdateCurrent(c echo.Context) error {
	var req dto.UpdateCurrentUserRequest
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

	resp, err := h.userService.UpdateCurrent(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.UserNotFound):
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

	return c.JSON(http.StatusOK, resp)

}
