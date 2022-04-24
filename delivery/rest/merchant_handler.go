package rest

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/afandi-syaikhu/majoo/constant"
	"github.com/afandi-syaikhu/majoo/model"
	"github.com/afandi-syaikhu/majoo/usecase"
	"github.com/labstack/echo/v4"
)

type MerchantHandler struct {
	AuthUseCase     usecase.AuthUseCase
	MerchantUseCase usecase.MerchantUseCase
	Config          *model.Config
}

func NewMerchantHandler(e *echo.Echo, authUC usecase.AuthUseCase, merchantUC usecase.MerchantUseCase, cfg *model.Config) {
	handler := &MerchantHandler{
		AuthUseCase:     authUC,
		MerchantUseCase: merchantUC,
		Config:          cfg,
	}

	// register route
	e.GET("/v1/merchants/:id/report", handler.GetReportByMerchantID)
}

func (_m *MerchantHandler) GetReportByMerchantID(c echo.Context) error {
	response := model.Response{}
	ctx := c.Request().Context()
	pathID := c.Param("id")

	// validate data request
	if len(strings.TrimSpace(pathID)) == 0 {
		response.Message = constant.NotFound
		c.JSON(http.StatusNotFound, response)

		return echo.ErrNotFound
	}

	id, err := strconv.Atoi(pathID)
	if err != nil {
		response.Message = constant.NotFound
		c.JSON(http.StatusNotFound, response)

		return echo.ErrNotFound
	}

	reqHeaderAuth := c.Request().Header.Get("Authorization")
	if len(strings.TrimSpace(reqHeaderAuth)) == 0 {
		response.Message = constant.Unauthorized
		c.JSON(http.StatusUnauthorized, response)

		return echo.ErrUnauthorized
	}

	tokenData, err := _m.AuthUseCase.ValidateToken(ctx, reqHeaderAuth)
	if err != nil {
		response.Message = err.Error()
		c.JSON(http.StatusUnauthorized, response)

		return echo.ErrUnauthorized
	}

	if tokenData == nil {
		response.Message = constant.InvalidToken
		c.JSON(http.StatusUnauthorized, response)

		return echo.ErrUnauthorized
	}

	isValid, err := _m.MerchantUseCase.IsValidMerchantForUser(ctx, int64(id), tokenData.Username)
	if err != nil {
		response.Message = constant.InternalServerError
		c.JSON(http.StatusInternalServerError, response)

		return echo.ErrInternalServerError
	}

	if !isValid {
		response.Message = constant.Forbidden
		c.JSON(http.StatusForbidden, response)

		return echo.ErrForbidden
	}

	paramLimit := c.QueryParam("limit")
	if len(strings.TrimSpace(paramLimit)) == 0 {
		paramLimit = strconv.Itoa(constant.ParamLimit)
	}

	limit, err := strconv.Atoi(paramLimit)
	if err != nil {
		response.Message = constant.BadRequest
		c.JSON(http.StatusBadRequest, response)

		return echo.ErrBadRequest
	}

	paramPage := c.QueryParam("page")
	if len(strings.TrimSpace(paramPage)) == 0 {
		paramPage = strconv.Itoa(constant.ParamPage)
	}

	page, err := strconv.Atoi(paramPage)
	if err != nil {
		response.Message = constant.BadRequest
		c.JSON(http.StatusBadRequest, response)

		return echo.ErrBadRequest
	}

	pagination := model.Pagination{
		Limit: limit,
		Page:  page,
	}

	res, err := _m.MerchantUseCase.GetReportByMerchantID(ctx, int64(id), pagination)
	if err != nil {
		response.Message = constant.InternalServerError
		c.JSON(http.StatusInternalServerError, response)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}
