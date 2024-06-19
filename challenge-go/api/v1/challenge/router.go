package challenge

import "github.com/gofiber/fiber/v2"

func Route(router fiber.Router) {
	controller := NewChallengeController()
	router.Post("/challenge", ValidateChallenge, controller.Challenge)
}
