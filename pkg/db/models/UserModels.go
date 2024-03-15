package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserId         uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserName       string
	Email          string `gorm:"unique"`
	HashedPassword string
	IsSuperUser    bool `gorm:"default:false"`
	IsVerify       bool `gorm:"default:false"`
}
