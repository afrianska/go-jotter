package controllers

import "github.com/gofiber/fiber/v2"

func GetHome(c *fiber.Ctx) error {
	return c.SendString("Welcome, You are at home page!")
}