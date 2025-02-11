package models

import (
	"database/sql"
	"time"
)

type Factory struct {
	FactoryID    string         `gorm:"primaryKey;column:factoryid"`
	UserID       string         `gorm:"column:userid;unique;not null"`
	FactoryName  sql.NullString `gorm:"column:factory_name"`
	CompanyName  string         `gorm:"column:companyname;not null"`
	Address      string         `gorm:"column:address;not null"`
	City         sql.NullString `gorm:"column:city"`
	Province     sql.NullString `gorm:"column:province"`
	Country      string         `gorm:"column:country;default:Thailand;not null"`
	PostCode     sql.NullString `gorm:"column:postcode"`
	Email        string         `gorm:"column:email;unique;not null"`
	Telephone    sql.NullString `gorm:"column:telephone"`
	LineID       sql.NullString `gorm:"column:lineid"`
	Facebook     sql.NullString `gorm:"column:facebook"`
	LocationLink sql.NullString `gorm:"column:location_link"`
	CreatedOn    time.Time      `gorm:"column:createdon;autoCreateTime"`
}

// ✅ บังคับ GORM ให้ใช้ตาราง `dairyfactory`
func (Factory) TableName() string {
	return "dairyfactory"
}
