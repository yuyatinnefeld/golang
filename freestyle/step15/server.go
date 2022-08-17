package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Golang World!")
	})

	app.Get("/api/list", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})

	app.Post("/api/register", func(c *fiber.Ctx) error {
		return c.SendString("I'm a POST request!")
	})

	app.Get("/api/param/:param", func(c *fiber.Ctx) error {
		return c.SendString("param: " + c.Params("param"))
	})

	app.Get("/get/json", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
