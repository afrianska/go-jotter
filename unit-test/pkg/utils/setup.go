package utils

import (
	"github.com/afrianska/playgo/unit-test/pkg/routers"
	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()

	routers.Home(app)

	return app
}