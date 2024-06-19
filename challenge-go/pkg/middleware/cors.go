package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsMiddleware() fiber.Handler {
	return cors.New(
		cors.Config{
			AllowCredentials: false,
			AllowOrigins:     "*",
			AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
		})
}
