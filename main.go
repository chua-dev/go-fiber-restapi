package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// c = context = come from request
func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func main() {
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}
