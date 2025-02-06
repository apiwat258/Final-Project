package controllers

import (
	"finalyearproject/Backend/database"
	"finalyearproject/Backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateFarmer(c *fiber.Ctx) error {
	type FarmerRequest struct {
		UserID              string `json:"userid"`
		CompanyName         string `json:"company_name"`
		FirstName           string `json:"firstname"`
		LastName            string `json:"lastname"`
		Email               string `json:"email"`
		Address             string `json:"address"`
		Address2            string `json:"address2"`
		AreaCode            string `json:"areacode"`
		Phone               string `json:"phone"`
		PostCode            string `json:"post"`
		City                string `json:"city"`
		UploadCertification string `json:"upload_certification"`
		LocationLink        string `json:"location_link"`
	}

	var req FarmerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// ตรวจสอบว่า User ID มีอยู่ในฐานข้อมูลหรือไม่
	var user models.User
	if err := database.DB.Where("userid = ?", req.UserID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User ID not found"})
	}

	// สร้างข้อมูล Farmer
	farmer := models.Farmer{
		FarmerID:    req.UserID,
		UserID:      req.UserID,
		FarmerName:  req.FirstName + " " + req.LastName,
		CompanyName: req.CompanyName,
		LocationID:  1, // ค่าเริ่มต้น อาจต้องแก้ไขให้เป็น location จริง
		Telephone:   req.Phone,
		LineID:      "",
		Facebook:    "",
	}

	// บันทึกลง Database
	if err := database.DB.Create(&farmer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save farmer data"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Farmer registered successfully"})
}
