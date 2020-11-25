package handler

import (
	"net/http"

	"github.com/cecepsprd/crowfu-api/internal/helpers"
	"github.com/cecepsprd/crowfu-api/internal/model"
	"github.com/cecepsprd/crowfu-api/internal/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(e *echo.Echo, as service.AuthService) {
	handler := &AuthHandler{
		authService: as,
	}
	e.POST("/v1/signin", handler.Login)
}

var Authentication = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(viper.GetString("secret_key")),
})

func (h *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var userReq model.User
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helpers.GetResponse(http.StatusUnprocessableEntity, err))
	}

	loggedinUser, err := h.authService.Login(ctx, userReq.Email, userReq.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.GetResponse(http.StatusBadRequest, "wrong email or password"))
	}

	loggedinUser.Token, err = h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.GetResponse(http.StatusBadRequest, err))
	}

	return c.JSON(http.StatusOK, helpers.GetResponse(http.StatusOK, loggedinUser))
}
