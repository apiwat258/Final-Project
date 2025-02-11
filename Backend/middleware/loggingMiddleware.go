package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	log.Printf("[REQUEST] %s %s from %s", c.Method(), c.OriginalURL(), c.IP())

	err := c.Next()

	statusCode := c.Response().StatusCode()
	elapsed := time.Since(start)
	log.Printf("[RESPONSE] %s %s - Status: %d - Duration: %v", c.Method(), c.OriginalURL(), statusCode, elapsed)

	return err
}
