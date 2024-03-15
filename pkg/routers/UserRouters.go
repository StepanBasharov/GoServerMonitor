package routers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"servermonitor/pkg/handlers"
)

func RegisterUserRouters(app *echo.Group, db *gorm.DB) {
	authGroup := app.Group("/auth")
	userGroup := app.Group("/user")

	// auth routers
	authGroup.POST("/login", handlers.Login)
	authGroup.POST("/logout", handlers.Logout)
	authGroup.POST("/register", handlers.Register)

	// user routers
	userGroup.GET("/me", handlers.UserMe)
}
