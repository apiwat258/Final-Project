package models

import (
	"time"
)

type Farmer struct {
	FarmerID    string `gorm:"primaryKey"`
	UserID      string `gorm:"foreignKey:UserID"`
	FarmerName  string
	CompanyName string
	LocationID  int
	Telephone   string
	LineID      string
	Facebook    string
	CreatedOn   time.Time `gorm:"autoCreateTime"`
}
