// ./404_handler.go
package handlers

import "github.com/gofiber/fiber/v2"

func NotFoundHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotFound)
}