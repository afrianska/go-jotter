// ./auth.go
package handlers

import (
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/afrianska/go-try/fiber-auth-jwt/app/models"
	"github.com/afrianska/go-try/fiber-auth-jwt/pkg/configs"
	"github.com/afrianska/go-try/fiber-auth-jwt/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*models.User, error) {
		
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Email : e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func getUserByUsername(u string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Username: u}).Find(&user).Error;err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	deleted := user.DeletedAt
	fmt.Println("this", deleted)
	return &user, nil
}

func isMail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity	string `json:"identity"`
		Password	string `json:"password"`
	}

	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	input := new(LoginInput)
	var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status" : "error", "message" : "Error on login request", "data" : err})
	}

	identity := input.Identity
	pass := input.Password
	userModel, err := new(models.User), *new(error)

	if isMail(identity) {
		userModel, err = getUserByEmail(identity)
	} else {
		userModel, err = getUserByUsername(identity)
	}
	
	if userModel == nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status" : "error", "message" : "User not found", "data" : err})
	} else {
		userData = UserData{
			ID:			userModel.ID,
			Username:	userModel.Username,
			Email:		userModel.Email,
			Password:	userModel.Password,
		}
	}

	if !CheckPasswordHash(pass, userData.Password) {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status" : "error", "message" : "Invalid password", "data" : err})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(configs.EnvConf("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status" : "success", "message" : "Succes login", "data": t})
}