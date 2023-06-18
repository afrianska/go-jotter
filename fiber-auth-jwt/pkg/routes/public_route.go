// ./public_route.go
package routes

import "github.com/gofiber/fiber/v2"

func PublicRoute(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status" : "success",
			"message" : "Hello world",
			"data" : nil,
		})
	})
}