package handler

import (
	"net/http"
	"strconv"

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
	e.POST("/v1/users", handler.CreateUser)
	e.PUT("/v1/users/:id", handler.UpdateUser)
	e.DELETE("/v1/users/:id", handler.DeleteUser)
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

func (u *UserHandler) CreateUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// TODO:VALIDA HERE LATER
	// ......

	ctx := c.Request().Context()
	_, err = u.userService.Save(ctx, &user)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}

	var user model.User
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// TODO:VALIDA HERE LATER
	// ......

	ctx := c.Request().Context()
	err = u.userService.Update(ctx, int64(id), &user)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}

	ctx := c.Request().Context()
	ra, err := u.userService.Delete(ctx, int64(id))
	if ra != 1 {
		return c.JSON(http.StatusNotFound, model.ResponseError{Message: "user not found :( "})
	}

	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusOK)
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
