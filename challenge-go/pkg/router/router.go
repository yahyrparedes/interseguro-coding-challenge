package router

import (
	"challenge-go/api"
	"github.com/gofiber/fiber/v2"
)

func MainRoutes(app *fiber.App) {
	// Swagger Route
	SwaggerRoute(app)

	// Grouping Api Routes
	api.Route(app.Group("/api"))
	// Grouping Static Routes
	//StaticRoute(app)
}
