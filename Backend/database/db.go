package database

import (
	"finalyearproject/Backend/config"
	"finalyearproject/Backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := config.GetEnv("DB_DSN", "host=localhost user=postgres password=password dbname=supplychain_db port=5432 sslmode=disable")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connected successfully")

	// Run AutoMigrate for necessary models
	DB.AutoMigrate(
		&models.User{},
		&models.Farmer{},
		&models.Certification{},
		&models.Factory{},
		&models.Logistics{},
		&models.Retailer{},
	)
}
