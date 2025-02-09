package routes

import (
	"finalyearproject/Backend/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1") // เปลี่ยนให้ทุก API ใช้ /api/v1

	api.Get("/", controllers.Welcome)
	api.Get("/health", controllers.HealthCheck)
	api.Post("/update-role", controllers.UpdateUserRole)
	api.Post("/farmer", controllers.CreateFarmer)                     // ✅ เพิ่ม Route สำหรับ Farmer
	api.Post("/createCertification", controllers.CreateCertification) // ✅ บันทึกใบรับรอง
	api.Post("/uploadCertificate", controllers.UploadCertificate)     // ✅ อัปโหลดใบรับรองไป IPFS
}
