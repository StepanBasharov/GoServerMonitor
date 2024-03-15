package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"servermonitor/pkg/tools"
)

func ProcessListHandler(c echo.Context) error {
	response, err := tools.GetProcessList()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot get process list")
	}
	return c.JSON(http.StatusOK, response)
}
