package controllers

import (
	"finalyearproject/Backend/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type BlockchainController struct {
	Service *services.BlockchainService
}

func NewBlockchainController(service *services.BlockchainService) *BlockchainController {
	return &BlockchainController{Service: service}
}

func (bc *BlockchainController) GetBlockNumber(c *fiber.Ctx) error {
	blockNumber, err := bc.Service.GetLatestBlockNumber()
	if err != nil {
		log.Println("Error retrieving block number:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve block number",
		})
	}
	return c.JSON(fiber.Map{
		"blockNumber": blockNumber.String(),
	})
}

func (bc *BlockchainController) GetRawMilkBatch(c *fiber.Ctx) error {
	rawMilkID := c.Params("rawMilkID")
	batch, err := bc.Service.GetRawMilkBatch(rawMilkID)
	if err != nil {
		log.Println("Error retrieving raw milk batch:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve raw milk batch",
		})
	}
	return c.JSON(batch)
}

func (bc *BlockchainController) GetProduct(c *fiber.Ctx) error {
	productID := c.Params("productID")
	product, err := bc.Service.GetProduct(productID)
	if err != nil {
		log.Println("Error retrieving product:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve product",
		})
	}
	return c.JSON(product)
}

func (bc *BlockchainController) GetProductLot(c *fiber.Ctx) error {
	productLotID := c.Params("productLotID")
	lot, err := bc.Service.GetProductLot(productLotID)
	if err != nil {
		log.Println("Error retrieving product lot:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve product lot",
		})
	}
	return c.JSON(lot)
}

func (bc *BlockchainController) GetShippingEvent(c *fiber.Ctx) error {
	shippingID := c.Params("shippingID")
	event, err := bc.Service.GetShippingEvent(shippingID)
	if err != nil {
		log.Println("Error retrieving shipping event:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve shipping event",
		})
	}
	return c.JSON(event)
}

func (bc *BlockchainController) AddProductLot(c *fiber.Ctx) error {
	var req struct {
		ProductLotID string   `json:"productLotID"`
		ProductID    string   `json:"productID"`
		QRCode       string   `json:"qrCode"`
		QuantityBox  uint64   `json:"quantityBox"`
		MfgDate      uint64   `json:"mfgDate"`
		ExpDate      uint64   `json:"expDate"`
		RawMilkIDs   []string `json:"rawMilkIDs"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	err := bc.Service.AddProductLot(req.ProductLotID, req.ProductID, req.QRCode, req.QuantityBox, req.MfgDate, req.ExpDate, req.RawMilkIDs)
	if err != nil {
		log.Println("Error adding product lot:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add product lot"})
	}
	return c.JSON(fiber.Map{"message": "Product lot added successfully"})
}

func (bc *BlockchainController) AddShippingEvent(c *fiber.Ctx) error {
	var req struct {
		ShippingID           string `json:"shippingID"`
		ProductLotID         string `json:"productLotID"`
		FromLocation         string `json:"fromLocation"`
		ToLocation           string `json:"toLocation"`
		Transporter          string `json:"transporter"`
		Temperature          uint64 `json:"temperature"`
		Humidity             uint64 `json:"humidity"`
		QualityInspectionCID string `json:"qualityInspectionCID"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	err := bc.Service.AddShippingEvent(req.ShippingID, req.ProductLotID, req.FromLocation, req.ToLocation, req.Transporter, req.Temperature, req.Humidity, req.QualityInspectionCID)
	if err != nil {
		log.Println("Error adding shipping event:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add shipping event"})
	}
	return c.JSON(fiber.Map{"message": "Shipping event added successfully"})
}
