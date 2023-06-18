package main

import (
	"log"

	"github.com/afrianska/go-try/fiber-auth-jwt/pkg/configs"
	"github.com/afrianska/go-try/fiber-auth-jwt/pkg/routes"
	"github.com/afrianska/go-try/fiber-auth-jwt/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	routes.PrivateRoute(app)
	routes.PublicRoute(app)
	routes.NotfoundRoute(app) // always on last line from other routes

	log.Fatal(app.Listen(configs.EnvConf("URL_PORT")))
}