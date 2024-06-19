package api

import (
	"challenge-go/api/v1/challenge"
	"github.com/gofiber/fiber/v2"
)

func Route(router fiber.Router) {
	v1 := router.Group("/v1")

	// Challenge Route
	challenge.Route(v1)

}
