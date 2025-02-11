package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    string         `gorm:"primaryKey;column:userid;default:gen_random_uuid()"`
	Email     string         `gorm:"unique;column:email;not null"`
	Password  string         `gorm:"column:password;not null"`
	Role      string         `gorm:"column:role;not null"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
