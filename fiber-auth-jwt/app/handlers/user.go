package handlers

import (
	"strconv"

	"github.com/afrianska/go-try/fiber-auth-jwt/app/models"
	"github.com/afrianska/go-try/fiber-auth-jwt/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(pass string) (string, error) {
	bytes , err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func validToken(token *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))
	return n == uid
}

func validUser(id, pass string) bool {
	db := database.DB
	var user models.User
	db.First(&user, id)
	if user.Username == "" {
		return false
	}
	if !CheckPasswordHash(pass, user.Password) {
		return false
	}
	return true
}

// GetUser func to get a user
func GetUser(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var user models.User
	db.Find(&user, id)
	if user.Username == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status" : "error", "msg" : "no user found with id", "data" : nil})
	}
	return c.JSON(fiber.Map{"status" : "success", "msg" : "user found", "data" : user })
}

// CreateUser func to create new user
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status" : "error", "msg" : "check your input", "data" : err})
	}
	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status" : "error", "msg" : "can't hash password", "data" : err})
	}
	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status" : "error", "msg" : "can't create user", "data" : err})
	}
	newUser := models.NewUser{
		Email: user.Email,
		Username: user.Username,
	}
	return c.JSON(fiber.Map{"status" : "success", "msg" : "user created", "data" : newUser })
}

// UpdateUser func to update data user
func UpdateUser(c *fiber.Ctx) error {
	var updateUserInput models.UpdateUserInput
	if err := c.BodyParser(&updateUserInput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status" : "error", "msg" : "check your input", "data" : err})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)
	
	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status" : "error", "msg" : "invalid token", "data" : nil})
	}
	db := database.DB
	var user models.User
	db.First(&user, id)
	user.Names = updateUserInput.Names
	db.Save(&user)
	return c.JSON(fiber.Map{"status" : "success", "msg" : "user update", "data" : user })
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB
	var userPass models.UserPassInput
	if err := c.BodyParser(&userPass); err != nil {
		return c.Status(500).JSON(fiber.Map{"status" : "error", "msg" : "check your input pass", "data" : err})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status" : "error", "msg" : "invalid token", "data" : nil})
	}

	if !validUser(id, userPass.Password) {
		return c.Status(500).JSON(fiber.Map{"status" : "error", "msg" : "invalid user", "data" : nil})
	}

	var user models.User
	db.First(&user, id)
	db.Delete(&user)
	return c.JSON(fiber.Map{"status" : "success", "msg" : "user deleted", "data" : user})
}