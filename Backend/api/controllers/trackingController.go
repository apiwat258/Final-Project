package controllers

import (
	"finalyearproject/Backend/database"
	"finalyearproject/Backend/models"
	"finalyearproject/Backend/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type TrackingController struct {
	QRService         *services.QRCodeService
	BlockchainService *services.BlockchainService
	IPFSService       *services.IPFSService
}

func NewTrackingController(qrService *services.QRCodeService, blockchainService *services.BlockchainService, ipfsService *services.IPFSService) *TrackingController {
	return &TrackingController{
		QRService:         qrService,
		BlockchainService: blockchainService,
		IPFSService:       ipfsService,
	}
}

func (tc *TrackingController) GenerateQRCode(c *fiber.Ctx) error {
	var req struct {
		UserID       string `json:"userID"`
		ProductLotID string `json:"productLotID"`
		ShippingID   string `json:"shippingID"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Check user role before allowing QR Code generation
	var user models.User
	if err := database.DB.Where("userid = ?", req.UserID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found or unauthorized"})
	}

	if user.Role != "farmer" && user.Role != "factory" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User does not have permission to generate QR Code"})
	}

	qrData := "ProductLot: " + req.ProductLotID + ", Shipping: " + req.ShippingID
	filePath, err := tc.QRService.GenerateQRCode(qrData)
	if err != nil {
		log.Println("Error generating QR Code:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate QR Code"})
	}

	// Upload QR Code image to IPFS
	cid, err := tc.IPFSService.UploadFile(filePath)
	if err != nil {
		log.Println("Error uploading QR Code to IPFS:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload QR Code to IPFS"})
	}

	// Store CID on Blockchain under Shipping Event
	err = tc.BlockchainService.AddShippingEventCID(req.ShippingID, cid)
	if err != nil {
		log.Println("Error saving QR Code CID to Blockchain:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save QR Code CID to Blockchain"})
	}

	return c.JSON(fiber.Map{"message": "QR Code generated and stored successfully", "cid": cid})
}

func (tc *TrackingController) ReadQRCode(c *fiber.Ctx) error {
	var req struct {
		FilePath string `json:"file_path"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	data, err := tc.QRService.ReadQRCode(req.FilePath)
	if err != nil {
		log.Println("Error reading QR Code:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read QR Code"})
	}

	return c.JSON(fiber.Map{"message": "QR Code read successfully", "data": data})
}
