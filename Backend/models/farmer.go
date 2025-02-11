package models

import (
	"database/sql"
	"time"
)

type Farmer struct {
	FarmerID     string         `gorm:"primaryKey;column:farmerid"`
	UserID       string         `gorm:"column:userid;unique;not null"`
	FarmerName   string         `gorm:"column:farmer_name;not null"`
	CompanyName  string         `gorm:"column:companyname;not null"`
	Address      string         `gorm:"column:address;not null"`
	City         string         `gorm:"column:city;not null"`
	Province     string         `gorm:"column:province;not null"`
	Country      string         `gorm:"column:country;default:Thailand"`
	PostCode     string         `gorm:"column:postcode;not null"`
	Telephone    string         `gorm:"column:telephone;not null"`
	LineID       sql.NullString `gorm:"column:lineid"`
	Facebook     sql.NullString `gorm:"column:facebook"`
	LocationLink sql.NullString `gorm:"column:location_link"`
	CreatedOn    time.Time      `gorm:"column:createdon;autoCreateTime"`
	Email        string         `gorm:"column:email;unique;not null"`
}

// ✅ บังคับ GORM ให้ใช้ตาราง `farmer` ที่มีอยู่
func (Farmer) TableName() string {
	return "farmer"
}
