package handlers

import "github.com/labstack/echo/v4"

func UserMe(c echo.Context) error {
	return echo.ErrUnauthorized
}
