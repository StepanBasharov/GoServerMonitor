package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"servermonitor/pkg/schemas"
)

func HealthCheckHandler(c echo.Context) error {
	resp := schemas.HealthCheck{Status: true}
	return c.JSON(http.StatusOK, resp)
}
