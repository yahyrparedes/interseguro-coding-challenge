package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a *fiber.App) {
	route := a.Group("/swagger")
	route.Get("*", swagger.HandlerDefault) // default
}

// Deprecated: SwaggerHandler func for describe group of API Docs routes.
//func SwaggerHandler() fiber.Handler {
//	cfg := swagger.Config{
//		BasePath: "/",
//		FilePath: "./docs/swagger.json",
//		Path:     "swagger",
//		Title:    "Swagger API Docs",
//	}
//
//	return swagger.New(cfg)
//}
