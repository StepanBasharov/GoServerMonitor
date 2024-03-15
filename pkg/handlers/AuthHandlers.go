package handlers

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"servermonitor/pkg/db/models"
	"servermonitor/pkg/schemas"
	"servermonitor/pkg/tools"
)

func Login(c echo.Context, db *gorm.DB) error {
	request := &schemas.LoginRequestSchema{}
	err := c.Bind(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	var user models.Users
	err = db.Where("user_name = ? OR email = ? AND hashed_password = ?",
		request.Login,
		request.Login,
		tools.CreateHashPassword(request.Password),
	).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect login or password")
		}
	}
	return echo.NewHTTPError(http.StatusOK, fmt.Sprintf("Hello, %s!", user.UserName))
}

func Logout(c echo.Context) error {
	return echo.ErrUnauthorized
}

func Register(c echo.Context, db *gorm.DB) error {
	request := &schemas.RegistrationRequestSchema{}
	err := c.Bind(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	if len(request.Password) < 8 {
		return echo.NewHTTPError(http.StatusOK, "passwords too short")
	}
	if request.Password != request.PasswordConfirm {
		return echo.NewHTTPError(http.StatusOK, "passwords don't match")
	}
	var usersWithEmail []models.Users
	db.Where("email = ?", request.Email).Find(&usersWithEmail)
	if len(usersWithEmail) != 0 {
		return echo.NewHTTPError(http.StatusConflict, "User with this email already exists")
	}
	newUser := models.Users{
		UserId:         uuid.New(),
		UserName:       request.Username,
		HashedPassword: tools.CreateHashPassword(request.Password),
		Email:          request.Email,
	}
	db.Create(&newUser)

	return c.String(http.StatusCreated, fmt.Sprintf("UserId: %s", &newUser.UserId))
}
