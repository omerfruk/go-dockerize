package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Merhaba ben dockerize olacak olan bir uygulama")
	})

	app.Listen(":80")
}
