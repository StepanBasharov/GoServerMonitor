package routers

import (
	"github.com/labstack/echo/v4"
	"servermonitor/pkg/handlers"
)

func RegisterServiceRouters(app *echo.Group) {
	app.GET("/healthcheck", handlers.HealthCheckHandler)
}
