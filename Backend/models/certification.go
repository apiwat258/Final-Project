package models

import (
	"time"
)

type Certification struct {
	CertificationID   string    `gorm:"primaryKey;column:certificationid"`
	FarmerID          string    `gorm:"column:farmerid"`
	CertificationType string    `gorm:"column:certificationtype"`
	CertificationCID  string    `gorm:"column:certificationcid;unique"`
	EffectiveDate     time.Time `gorm:"column:effective_date"`
	IssuedDate        time.Time `gorm:"column:issued_date"`
	CreatedOn         time.Time `gorm:"column:createdon;autoCreateTime"`
}

// ✅ บังคับ GORM ให้ใช้ตาราง `organiccertification` ที่มีอยู่
func (Certification) TableName() string {
	return "organiccertification"
}
