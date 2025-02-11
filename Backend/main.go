package main

import (
	"finalyearproject/Backend/api/controllers"
	"finalyearproject/Backend/api/routes"
	"finalyearproject/Backend/config"
	"finalyearproject/Backend/database"
	"finalyearproject/Backend/middleware"
	"finalyearproject/Backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// ✅ Load Environment Variables
	config.LoadEnv()

	// ✅ Enable Logging Middleware
	app.Use(middleware.LoggingMiddleware)

	// ✅ Enable CORS for Frontend connection
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://0.0.0.0:8081",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH,OPTIONS",
	}))

	// ✅ Connect to Database
	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Farmer{}, &models.Factory{}, &models.Logistic{}, &models.Retailer{}, &models.Certification{})

	// ✅ Serve Static Files (For Frontend)
	app.Static("/", "./frontend")

	// ✅ Define API Routes
	app.Post("/api/register", controllers.RegisterUser)
	app.Post("/api/login", controllers.LoginUser)
	app.Post("/api/uploadCertificate", controllers.UploadCertificate)

	// ✅ Setup API Routes
	routes.SetupRoutes(app)

	// ✅ Start Server
	app.Listen(":8080")
}
