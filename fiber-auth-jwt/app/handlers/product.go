// ./product.go
package handlers

import (
	"github.com/afrianska/go-try/fiber-auth-jwt/app/models"
	"github.com/afrianska/go-try/fiber-auth-jwt/platform/database"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []models.Product
	db.Find(&products)
	if count := len(products); count == 0 {
		return c.Status(404).JSON(fiber.Map{"status" : "error", "msg" : "no product recorded", "data" : nil})
	}
	return c.JSON(fiber.Map{"status" : "success", "msg" : "all products", "data" : products})
}

func GetProduct(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var product models.Product
	db.Find(&product, id)
	if product.Tittle == "" {
		return c.Status(404).JSON(fiber.Map{"status" : "error", "msg" : "product not found", "data": nil})
	}
	return c.JSON(fiber.Map{"status" : "success", "msg" : "product found", "data" : product})
}

func CreateProduct(c *fiber.Ctx) error {
	db := database.DB
	product := new(models.Product)
	if err := c.BodyParser(&product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status" : "error", "msg" : "check your input data", "data" : nil})
	}
	db.Create(&product)
	return c.JSON(fiber.Map{"status" : "success", "msg" : "product created", "data" : product})
}

func UpdateProduct(c *fiber.Ctx) error {
	db := database.DB

	var upi models.UpdateProductInput
	if err := c.BodyParser(&upi); err != nil {
		return c.Status(500).JSON(fiber.Map{"status" : "error", "msg" : "check product input", "data" : nil})
	}
	id := c.Params("id")
	var product models.Product
	if err := db.First(&product, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status" : "error", "msg" : "no product found with id", "data" : nil})
	}

	if upi.Tittle != "" {  product.Tittle =  upi.Tittle}
	if upi.Description != "" {  product.Description = upi.Description }
	if upi.Amount != 0 { product.Amount = upi.Amount }
	db.Save(&product)
	return c.JSON(fiber.Map{"status" : "success", "msg" : "product updated", "data" : product})
}

func DeleteProduct(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var product models.Product
	if err := db.First(&product, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status" : "error", "msg" : "product not found", "data" : nil})
	}
	db.Delete(&product)
	return c.Status(200).JSON(fiber.Map{"status" : "success", "msg" : "product deleted", "data" : nil})
}