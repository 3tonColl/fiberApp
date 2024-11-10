package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	htmlengine := html.New("../../web/html", ".html")
	app := fiber.New(fiber.Config{AppName: "fiberApp v0.1",
		Views: htmlengine})
	app.Static("/", "../../web")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		return c.SendString("recieved registration request!")
	})
	log.Fatal(app.Listen(":3000"))
}
