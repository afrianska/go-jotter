// ./main.go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := fiber.New()

	// create available route: "/hello"
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello everyone!")
	})

	// Not found handler
	app.Use(func (f *fiber.Ctx) error {
		return f.SendStatus(fiber.StatusNotFound) // not found status 404
	})

	log.Fatal(app.Listen(":5000"))
}