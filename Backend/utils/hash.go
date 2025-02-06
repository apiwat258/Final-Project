package utils

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateUserID สร้าง userid โดยใช้ปีปัจจุบันและหมายเลขลำดับ
func GenerateUserID(sequence int) string {
	currentYear := time.Now().Year() % 100 // รับปีปัจจุบันและใช้เฉพาะสองหลักสุดท้าย
	return fmt.Sprintf("%02d%04d", currentYear, sequence)
}
