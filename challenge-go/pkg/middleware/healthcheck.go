package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func HealthCheck() fiber.Handler {

	// Healthcheck middleware default
	//return healthcheck.New()

	// Healthcheck middleware custom
	return healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			// Add your own liveness check here
			return true
		},
		LivenessEndpoint: "/liveness",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			// Add your own readiness check here
			return true
		},
		ReadinessEndpoint: "/readiness",
	})
}
