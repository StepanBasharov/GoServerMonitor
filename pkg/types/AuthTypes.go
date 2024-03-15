package types

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"servermonitor/pkg/db/models"
	"servermonitor/pkg/utils"
)

type RegistrationRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (loginRequest LoginRequest) LoginOrNotFound(db *gorm.DB, c echo.Context) error {
	var user models.Users
	err := db.Where("user_name = ? OR email = ? AND hashed_password = ?",
		loginRequest.Login,
		loginRequest.Login,
		utils.CreateHashPassword(loginRequest.Password),
	).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect login or password")
		}
	}
	return c.String(http.StatusOK, fmt.Sprintf("Hello, %s!", user.UserName))
}

func (registrationRequest RegistrationRequest) RegisterOrError(db *gorm.DB, c echo.Context) error {
	if len(registrationRequest.Password) < 8 {
		return echo.NewHTTPError(http.StatusOK, "passwords too short")
	}
	if registrationRequest.Password != registrationRequest.PasswordConfirm {
		return echo.NewHTTPError(http.StatusOK, "passwords don't match")
	}
	var usersWithEmail []models.Users
	db.Where("email = ?", registrationRequest.Email).Find(&usersWithEmail)
	if len(usersWithEmail) != 0 {
		return echo.NewHTTPError(http.StatusConflict, "User with this email already exists")
	}
	newUser := models.Users{
		UserId:         uuid.New(),
		UserName:       registrationRequest.Username,
		HashedPassword: utils.CreateHashPassword(registrationRequest.Password),
		Email:          registrationRequest.Email,
	}
	db.Create(&newUser)

	return c.String(http.StatusCreated, fmt.Sprintf("UserId: %s", &newUser.UserId))
}
