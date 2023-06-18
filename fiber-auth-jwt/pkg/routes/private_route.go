// ./private_route.go
package routes

import (
	"github.com/afrianska/go-try/fiber-auth-jwt/app/handlers"
	"github.com/afrianska/go-try/fiber-auth-jwt/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func PrivateRoute(r *fiber.App)  {
	// Middleware
	api := r.Group("/api", logger.New())
	api.Get("/", handlers.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)

	// User
	user := api.Group("/user")
	user.Get("/:id", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
	user.Patch("/:id", middlewares.Protected(), handlers.UpdateUser) //
	user.Delete("/:id", middlewares.Protected(), handlers.DeleteUser) //

	// User
	product := api.Group("/product")
	product.Get("/", handlers.GetProducts)
	product.Get("/:id", handlers.GetProduct)
	product.Post("/", handlers.CreateProduct)
	product.Patch("/:id", middlewares.Protected(), handlers.UpdateProduct)
	product.Delete("/:id", middlewares.Protected(), handlers.DeleteProduct)
}