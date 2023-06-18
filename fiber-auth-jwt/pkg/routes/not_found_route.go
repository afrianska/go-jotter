package routes

import (
	"github.com/afrianska/go-try/fiber-auth-jwt/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func NotfoundRoute(r *fiber.App) {
	r.Use(handlers.NotFoundHandler)
}