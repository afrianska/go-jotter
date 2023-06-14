package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload" // aouto load godotenv
)

type User struct {
	ID 			uint	`validate:"required,omitempty"`
	Username	string	`validate:"required"`
	Password	string	`validate:"gte=6"` // gte = greater than or equal (6 is value)

}

func getTest(ctx *fiber.Ctx) error {
	// // sample data input to match with struct
		user := User{
			ID : 1,
			Username : "afrian",
			Password : "F!b#r890",
		}

		validate := validator.New() //Create validate for using
		if err := validate.Struct(user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

	return ctx.Status(fiber.StatusOK).JSON("Success")
}

func main() {
	app := fiber.New()
	

	// use cors
	app.Use(cors.New())
	app.Get("/tests", getTest)
	log.Fatal(app.Listen(os.Getenv("PORT")))
}