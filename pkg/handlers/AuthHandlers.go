package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"servermonitor/pkg/types"
)

func Login(c echo.Context, db *gorm.DB) error {
	request := &types.LoginRequest{}
	err := c.Bind(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	return request.LoginOrNotFound(db, c)
}

func Register(c echo.Context, db *gorm.DB) error {
	request := &types.RegistrationRequest{}
	err := c.Bind(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	return request.RegisterOrError(db, c)
}
