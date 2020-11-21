package handler

import (
	"net/http"
	"strconv"

	"github.com/cecepsprd/crowfu-api/internal/helpers"
	"github.com/cecepsprd/crowfu-api/internal/model"
	"github.com/cecepsprd/crowfu-api/internal/service"
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
		return c.JSON(http.StatusInternalServerError, helpers.GetResponse(http.StatusInternalServerError, err))
	}

	return c.JSON(http.StatusOK, helpers.GetResponse(http.StatusOK, listUser))
}

func (u *UserHandler) CreateUser(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helpers.GetResponse(http.StatusUnprocessableEntity, err))
	}

	// Validate user
	if err = user.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.GetResponse(http.StatusBadRequest, err))
	}

	_, err = u.userService.Save(ctx, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.GetResponse(http.StatusInternalServerError, err))
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.GetResponse(http.StatusNotFound))
	}

	var user model.User
	if err = c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helpers.GetResponse(http.StatusUnprocessableEntity, err))
	}

	// Validate user
	if err = user.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.GetResponse(http.StatusBadRequest, err))
	}

	rowsAffected, err := u.userService.Update(ctx, int64(id), &user)
	if rowsAffected != 1 {
		return c.JSON(http.StatusNotFound, helpers.GetResponse(http.StatusNotFound))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.GetResponse(http.StatusInternalServerError, err))
	}

	return c.JSON(http.StatusOK, helpers.GetResponse(http.StatusOK, user))
}

func (u *UserHandler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}

	rowsAffected, err := u.userService.Delete(ctx, int64(id))
	if rowsAffected != 1 {
		return c.JSON(http.StatusNotFound, helpers.GetResponse(http.StatusNotFound))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.GetResponse(http.StatusInternalServerError, err))
	}

	return c.JSON(http.StatusOK, helpers.GetResponse(http.StatusOK))
}
