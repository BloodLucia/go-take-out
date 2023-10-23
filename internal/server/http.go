package server

import "github.com/gofiber/fiber/v2"

func NewHTTPServer() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}
