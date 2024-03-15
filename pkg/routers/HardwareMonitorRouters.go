package routers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"servermonitor/pkg/handlers"
)

func RegisterHardwareMonitorRouters(app *echo.Group, db *gorm.DB) {
	hardwareRouters := app.Group("/hardware")
	hardwareRouters.GET("/process_list", handlers.ProcessListHandler)
}
