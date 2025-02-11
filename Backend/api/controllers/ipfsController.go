package controllers

import (
	"finalyearproject/Backend/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type IPFSController struct {
	Service *services.IPFSService
}

func NewIPFSController(service *services.IPFSService) *IPFSController {
	return &IPFSController{Service: service}
}

func (ic *IPFSController) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid file upload"})
	}

	openFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open file"})
	}
	defer openFile.Close()

	cid, err := ic.Service.UploadFile(openFile, file)
	if err != nil {
		log.Println("Error uploading file to IPFS:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload file to IPFS"})
	}

	return c.JSON(fiber.Map{"message": "File uploaded successfully", "cid": cid})
}

func (ic *IPFSController) RetrieveFile(c *fiber.Ctx) error {
	cid := c.Params("cid")
	if cid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "CID is required"})
	}

	fileURL := ic.Service.IPFSGateway + "/ipfs/" + cid
	return c.JSON(fiber.Map{"file_url": fileURL})
}
