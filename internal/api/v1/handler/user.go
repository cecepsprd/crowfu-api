package handler

import (
	"net/http"

	"github.com/cecepsprd/crowfu-api/internal/model"
	"github.com/cecepsprd/crowfu-api/internal/service"
	"github.com/cecepsprd/crowfu-api/pkg/log"
	"github.com/labstack/echo"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(e *echo.Echo, us service.UserService) {
	handler := &UserHandler{
		userService: us,
	}
	e.GET("/v1/users", handler.GetListUser)
}

func (u *UserHandler) GetListUser(c echo.Context) error {
	ctx := c.Request().Context()
	listUser, err := u.userService.Get(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	c.Response().Header().Set(`X-Cursor`, "")
	return c.JSON(http.StatusOK, listUser)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Error(err)
	switch err {
	case model.ErrInternalServerError:
		return http.StatusInternalServerError
	case model.ErrNotFound:
		return http.StatusNotFound
	case model.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
