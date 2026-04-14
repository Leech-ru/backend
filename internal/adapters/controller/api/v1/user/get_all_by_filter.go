package user

import (
	"Leech-ru/internal/domain/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetAllByFilter returns a list of users based on filter parameters.
//
// @Summary     Get users by filters
// @Description Retrieves a list of users filtered by role, one-line full name query, and email prefix. Only for admins
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       limit          query     int     false  "Max number of users to return"   minimum(1) maximum(100)  example(10)
// @Param       offset         query     int     false  "Pagination offset"               minimum(0)               example(0)
// @Param       role           query     int     false  "Role enum (0–2)"                 Enums(0,1,2)             example(0)
// @Param       q              query     string  false  "One-line search by full name tokens (name/surname)"       example("Иван Дима")
// @Param       email_prefix   query     string  false  "Filter by email prefix"                                  example("user@")
// @Security CookieAuth
// @Success     200  {array}   dto.User
// @Failure     400  {object}  dto.HTTPStatus "Invalid query parameters"
// @Failure     500  {object}  dto.HTTPStatus "Internal server error"
// @Router      /api/v1/user/all [get]
func (h *handler) GetAllByFilter(c echo.Context) error {
	var req dto.GetAllByFilterUsersRequest

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

	resp, err := h.userService.GetAllByFilter(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
