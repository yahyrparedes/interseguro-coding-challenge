package server

import (
	"github.com/gofiber/contrib/fibernewrelic"
	"github.com/gofiber/fiber/v2"
)

func NewRelicConfig() fiber.Handler {
	// NewRelic
	cfg := fibernewrelic.Config{
		License: "47630f27d068fc7215af54f614b7c5b1FFFFNRAL",
		AppName: "challenge-go",
		Enabled: true,
	}
	return fibernewrelic.New(cfg)
}
