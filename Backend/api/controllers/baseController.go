package controllers

import (
	"finalyearproject/Backend/database"
	"finalyearproject/Backend/models"

	"github.com/gofiber/fiber/v2"
)

func UpdateUserRole(c *fiber.Ctx) error {
	type RoleRequest struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	var req RoleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	user.Role = req.Role
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user role"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User role updated successfully"})
}

// Welcome route
func Welcome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Welcome to Dairy Supply Chain API"})
}

// HealthCheck route
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "OK"})
}
