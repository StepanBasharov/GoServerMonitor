package db

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"servermonitor/pkg/db/models"
	"servermonitor/pkg/schemas"
	"servermonitor/pkg/tools"
)

func CreateConnection(config schemas.DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Users{})
	return db, nil
}

func CreateSuperUser(db *gorm.DB, config schemas.SuperUserConfig) {
	var superuser []models.Users

	db.Where("is_super_user = ?", true).Find(&superuser)
	passwordHash := tools.CreateHashPassword(config.Password)
	if len(superuser) == 0 {
		db.Create(&models.Users{
			UserId:         uuid.New(),
			UserName:       config.Username,
			Email:          config.Email,
			HashedPassword: passwordHash,
			IsSuperUser:    true,
			IsVerify:       true,
		})
	}
}
