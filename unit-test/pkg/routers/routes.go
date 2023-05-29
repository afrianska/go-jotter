package routers

import (
	"github.com/afrianska/playgo/unit-test/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func Home(r *fiber.App) {

	r.Get("/", controllers.GetHome)
}