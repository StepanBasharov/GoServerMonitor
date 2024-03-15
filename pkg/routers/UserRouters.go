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

	authGroup.POST("/logout", handlers.Logout)

	registrationHandler := func(c echo.Context) error {
		return handlers.Register(c, db)
	}
	loginHandler := func(c echo.Context) error {
		return handlers.Login(c, db)
	}

	authGroup.POST("/register", registrationHandler)
	authGroup.POST("/login", loginHandler)

	// user routers
	userGroup.GET("/me", handlers.UserMe)
}
