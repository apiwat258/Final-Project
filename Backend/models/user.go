package models

import (
	"time"
)

type User struct {
	UserID    string    `gorm:"primaryKey;column:userid"`
	Email     string    `gorm:"unique;column:email"`
	Password  string    `gorm:"column:password"`
	Role      string    `gorm:"column:role"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
