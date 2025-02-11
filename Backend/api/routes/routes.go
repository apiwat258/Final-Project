package routes

import (
	"finalyearproject/Backend/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1") // ✅ Ensure all routes use /api/v1

	// ✅ Health Check and Root Route
	api.Get("/", controllers.Welcome)
	api.Get("/health", controllers.HealthCheck)

	// ✅ User and Authentication Routes
	api.Post("/register", controllers.RegisterUser) // ✅ Register a new user
	api.Post("/login", controllers.LoginUser)       // ✅ User login
	api.Post("/update-role", controllers.UpdateUserRole)

	// ✅ Farmer Routes
	api.Post("/farmer", controllers.CreateFarmer)
	api.Get("/farmer/:id", controllers.GetFarmerByID)
	api.Put("/farmer/:id", controllers.UpdateFarmer)

	// ✅ Certification Routes
	api.Post("/createCertification", controllers.CreateCertification) // ✅ Save Certification
	api.Post("/uploadCertificate", controllers.UploadCertificate)     // ✅ Upload Certification to IPFS
	api.Get("/certification/:id", controllers.GetCertification)

	// ✅ Blockchain Integration Routes
	api.Post("/blockchain/rawmilk", controllers.RecordRawMilk)
	api.Get("/blockchain/rawmilk/:id", controllers.GetRawMilkDetails)
	api.Post("/blockchain/product", controllers.RecordProduct)
	api.Get("/blockchain/product/:id", controllers.GetProductDetails)
}
