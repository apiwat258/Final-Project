package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv โหลดไฟล์ .env เพื่อใช้ตั้งค่าระบบ
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}
}

// GetEnv คืนค่าตัวแปรสิ่งแวดล้อมหรือค่าเริ่มต้นหากไม่มี
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
