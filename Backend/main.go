package main

import (
	"finalyearproject/Backend/api/routes"
	"finalyearproject/Backend/database"
	"finalyearproject/Backend/middleware"
	"finalyearproject/Backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// เพิ่มมิดเดิลแวร์ CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://0.0.0.0:8081", // โดเมนของ frontend
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// เชื่อมต่อกับฐานข้อมูลและทำการ Migrate
	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Farmer{})

	// ให้บริการไฟล์สแตติกจากไดเรกทอรี frontend
	app.Static("/", "./frontend")

	// กำหนดเส้นทางสำหรับ API
	app.Post("/api/register", middleware.Register)

	// ✅ เรียกใช้ SetupRoutes() โดยส่ง *fiber.App
	routes.SetupRoutes(app)

	// เริ่มเซิร์ฟเวอร์ที่พอร์ต 8080
	app.Listen(":8080")
}
