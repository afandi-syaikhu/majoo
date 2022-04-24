package rest

import (
	"net/http"

	"github.com/afandi-syaikhu/majoo/model"
	"github.com/afandi-syaikhu/majoo/usecase"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthUseCase usecase.AuthUseCase
}

func NewAuthHandler(e *echo.Echo, authUseCase usecase.AuthUseCase) {
	handler := &AuthHandler{
		AuthUseCase: authUseCase,
	}

	// register route
	e.POST("/login", handler.Login)
}

func (_a *AuthHandler) Login(c echo.Context) error {
	response := model.Response{}
	ctx := c.Request().Context()
	body := model.Auth{}
	if err := c.Bind(&body); err != nil {
		response.Message = "invalid data"
		c.JSON(http.StatusBadRequest, response)

		return echo.ErrBadRequest
	}

	res, err := _a.AuthUseCase.Login(ctx, body)
	if err != nil {
		response.Message = "internal error"
		c.JSON(http.StatusInternalServerError, response)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}
