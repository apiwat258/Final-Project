package controllers

import (
	"database/sql"
	"finalyearproject/Backend/database"
	"finalyearproject/Backend/models"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateRetailer(c *fiber.Ctx) error {
	type RetailerRequest struct {
		UserID       string  `json:"userid"`
		CompanyName  string  `json:"company_name"`
		FirstName    string  `json:"firstname"`
		LastName     string  `json:"lastname"`
		Email        string  `json:"email"`
		Address      string  `json:"address"`
		Address2     *string `json:"address2"`
		AreaCode     *string `json:"areacode"`
		Phone        string  `json:"phone"`
		PostCode     string  `json:"post"`
		City         string  `json:"city"`
		Province     string  `json:"province"`
		Country      string  `json:"country"`
		LineID       *string `json:"lineid"`
		Facebook     *string `json:"facebook"`
		LocationLink *string `json:"location_link"`
	}

	var req RetailerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	if err := database.DB.Where("userid = ?", req.UserID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User ID not found in users table"})
	}

	var existingRetailer models.Retailer
	err := database.DB.Where("userid = ?", req.UserID).First(&existingRetailer).Error
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User is already registered as a retailer"})
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	if err := database.DB.Model(&models.User{}).Where("userid = ?", req.UserID).Update("role", "retailer").Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user role"})
	}

	var sequence int64
	if err := database.DB.Raw("SELECT nextval('retailer_id_seq')").Scan(&sequence).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate retailer ID"})
	}
	yearPrefix := time.Now().Format("06")
	retailerID := fmt.Sprintf("%s%05d", yearPrefix, sequence)

	fullAddress := strings.TrimSpace(req.Address)
	if req.Address2 != nil && strings.TrimSpace(*req.Address2) != "" {
		fullAddress = fullAddress + ", " + strings.TrimSpace(*req.Address2)
	}

	fullPhone := strings.TrimSpace(req.Phone)
	if req.AreaCode != nil && strings.TrimSpace(*req.AreaCode) != "" {
		areaCode := strings.TrimSpace(*req.AreaCode)
		if !strings.HasPrefix(areaCode, "+") {
			areaCode = "+" + areaCode
		}
		fullPhone = areaCode + " " + fullPhone
	}

	companyName := strings.TrimSpace(req.CompanyName)
	if companyName == "" {
		companyName = "N/A"
	}

	province := strings.TrimSpace(req.Province)
	if province == "" {
		province = req.City
	}

	email := sql.NullString{}
	if strings.TrimSpace(req.Email) != "" {
		email = sql.NullString{String: strings.TrimSpace(req.Email), Valid: true}
	}

	lineID := sql.NullString{}
	if req.LineID != nil && strings.TrimSpace(*req.LineID) != "" {
		lineID = sql.NullString{String: strings.TrimSpace(*req.LineID), Valid: true}
	}

	facebook := sql.NullString{}
	if req.Facebook != nil && strings.TrimSpace(*req.Facebook) != "" {
		facebook = sql.NullString{String: strings.TrimSpace(*req.Facebook), Valid: true}
	}

	locationLink := sql.NullString{}
	if req.LocationLink != nil && strings.TrimSpace(*req.LocationLink) != "" {
		locationLink = sql.NullString{String: strings.TrimSpace(*req.LocationLink), Valid: true}
	}

	retailer := models.Retailer{
		RetailerID:   retailerID,
		UserID:       req.UserID,
		CompanyName:  companyName,
		Address:      fullAddress,
		City:         req.City,
		Province:     province,
		Country:      req.Country,
		PostCode:     req.PostCode,
		Email:        email.String,
		Telephone:    fullPhone,
		LineID:       lineID,
		Facebook:     facebook,
		LocationLink: locationLink,
		CreatedOn:    time.Now(),
	}

	if err := database.DB.Create(&retailer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save retailer data"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Retailer registered successfully", "retailer_id": retailer.RetailerID})
}

func GetRetailerByID(c *fiber.Ctx) error {
	retailerID := c.Params("id")
	var retailer models.Retailer
	if err := database.DB.Where("retailerid = ?", retailerID).First(&retailer).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Retailer not found"})
	}
	return c.JSON(retailer)
}

func UpdateRetailer(c *fiber.Ctx) error {
	retailerID := c.Params("id")
	updateData := make(map[string]interface{})
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := database.DB.Model(&models.Retailer{}).Where("retailerid = ?", retailerID).Updates(updateData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update retailer"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Retailer updated successfully"})
}

func DeleteRetailer(c *fiber.Ctx) error {
	retailerID := c.Params("id")
	if err := database.DB.Where("retailerid = ?", retailerID).Delete(&models.Retailer{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete retailer"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Retailer deleted successfully"})
}
