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

func CreateFactory(c *fiber.Ctx) error {
	type FactoryRequest struct {
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

	var req FactoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	if err := database.DB.Where("userid = ?", req.UserID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User ID not found in users table"})
	}

	var existingFactory models.Factory
	err := database.DB.Where("userid = ?", req.UserID).First(&existingFactory).Error
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User is already registered as a factory"})
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	if err := database.DB.Model(&models.User{}).Where("userid = ?", req.UserID).Update("role", "factory").Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user role"})
	}

	var sequence int64
	if err := database.DB.Raw("SELECT nextval('dairyfactory_id_seq')").Scan(&sequence).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate factory ID"})
	}
	yearPrefix := time.Now().Format("06")
	factoryID := fmt.Sprintf("%s%05d", yearPrefix, sequence)

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

	factory := models.Factory{
		FactoryID:    factoryID,
		UserID:       req.UserID,
		FactoryName:  strings.TrimSpace(req.FirstName) + " " + strings.TrimSpace(req.LastName),
		CompanyName:  companyName,
		Address:      fullAddress,
		City:         req.City,
		Province:     province,
		Country:      req.Country,
		PostCode:     req.PostCode,
		Telephone:    fullPhone,
		LineID:       lineID,
		Facebook:     facebook,
		LocationLink: locationLink,
		CreatedOn:    time.Now(),
		Email:        email.String,
	}

	if err := database.DB.Create(&factory).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save factory data"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Factory registered successfully", "factory_id": factory.FactoryID})
}

func GetFactoryByID(c *fiber.Ctx) error {
	factoryID := c.Params("id")

	var factory models.Factory
	if err := database.DB.Where("factoryid = ?", factoryID).First(&factory).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Factory not found"})
	}

	return c.JSON(factory)
}

func DeleteFactory(c *fiber.Ctx) error {
	factoryID := c.Params("id")
	if err := database.DB.Where("factoryid = ?", factoryID).Delete(&models.Factory{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete factory"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Factory deleted successfully"})
}

func UpdateFactory(c *fiber.Ctx) error {
	factoryID := c.Params("id")
	updateData := make(map[string]interface{})
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := database.DB.Model(&models.Factory{}).Where("factoryid = ?", factoryID).Updates(updateData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update factory"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Factory updated successfully"})
}
