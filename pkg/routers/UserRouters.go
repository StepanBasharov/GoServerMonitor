package routers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"servermonitor/pkg/handlers"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (h *AuthHandler) Login(c echo.Context) error {
	return handlers.Login(c, h.DB)
}

func (h *AuthHandler) Registration(c echo.Context) error {
	return handlers.Register(c, h.DB)
}

func RegisterUserRouters(app *echo.Group, db *gorm.DB) {
	authGroup := app.Group("/auth")
	userGroup := app.Group("/user")

	// auth routers
	authHandlers := &AuthHandler{DB: db}

	authGroup.POST("/register", authHandlers.Registration)
	authGroup.POST("/login", authHandlers.Login)

	// user routers
	userGroup.GET("/me", handlers.UserMe)
}
