package middleware

import (
	"finalyearproject/Backend/database"
	"finalyearproject/Backend/models"
	"finalyearproject/Backend/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// ตรวจสอบว่าอีเมลนี้มีอยู่ในระบบแล้วหรือไม่
	var existingUser models.User
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already exists"})
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// ดึงค่าลำดับถัดไปจากลำดับในฐานข้อมูล
	var sequence int64
	if err := database.DB.Raw("SELECT nextval('user_id_seq')").Scan(&sequence).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate user ID"})
	}

	// สร้าง userid ในรูปแบบ YYNNNNN
	yearPrefix := time.Now().Format("06") // ได้เลขปีสองหลัก เช่น "25" สำหรับปี 2025
	userID := fmt.Sprintf("%s%05d", yearPrefix, sequence)

	user := models.User{
		UserID:   userID,
		Email:    req.Email,
		Password: hashedPassword,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully", "user_id": user.UserID})
}
