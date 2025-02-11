package models

import (
	"database/sql"
	"time"
)

type Retailer struct {
	RetailerID   string         `gorm:"primaryKey;column:retailerid"`
	UserID       string         `gorm:"column:userid;unique"`
	CompanyName  string         `gorm:"column:companyname;not null"`
	Address      string         `gorm:"column:address;not null"`
	City         string         `gorm:"column:city;not null"`
	Province     string         `gorm:"column:province;not null"`
	Country      string         `gorm:"column:country;default:Thailand;not null"`
	PostCode     string         `gorm:"column:postcode;not null"`
	Email        string         `gorm:"column:email;unique;not null"`
	Telephone    sql.NullString `gorm:"column:telephone"`
	LineID       sql.NullString `gorm:"column:lineid"`
	Facebook     sql.NullString `gorm:"column:facebook"`
	LocationLink sql.NullString `gorm:"column:location_link"`
	CreatedOn    time.Time      `gorm:"column:createdon;autoCreateTime"`
}

// ✅ บังคับ GORM ให้ใช้ตาราง `retailer`
func (Retailer) TableName() string {
	return "retailer"
}
