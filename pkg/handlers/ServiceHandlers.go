package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"servermonitor/pkg/types"
)

func HealthCheckHandler(c echo.Context) error {
	resp := types.HealthCheck{Status: true}
	return c.JSON(http.StatusOK, resp)
}
