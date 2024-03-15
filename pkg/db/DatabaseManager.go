package db

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"servermonitor/pkg/db/models"
	"servermonitor/pkg/types"
	"servermonitor/pkg/utils"
)

func CreateConnection(config types.DatabaseConfig) (*gorm.DB, error) {
	dsn := config.GetDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Users{})
	return db, nil
}

func CreateSuperUser(db *gorm.DB, config types.SuperUserConfig) {
	var superuser []models.Users

	db.Where("is_super_user = ?", true).Find(&superuser)
	passwordHash := utils.CreateHashPassword(config.Password)
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
