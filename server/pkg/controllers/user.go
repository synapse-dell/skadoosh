package controllers

import (
	"gocrud/pkg/config"
	"gocrud/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// var users = []models.User{}
//no need for this cause we are going to use actual DB

func Health(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func GetAllUser(c *fiber.Ctx) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil
	}
	return c.JSON(fiber.Map{
		"message": "Fetched all users",
		"Users":   users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"Message": "User created",
		"user":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(400, "Invalid user ID")
	}
	var user models.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		return fiber.NewError(404, "User not Found")
	}
	var updatedUser models.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return err
	}
	user.Age = updatedUser.Age
	user.Name = updatedUser.Name

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "User updated",
		"user":    user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(400, "Invalid ID param")
	}
	var user models.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		return fiber.NewError(404, "User not found")
	}
	if err := config.DB.Delete(user).Error; err != nil {
		return fiber.NewError(500, "Cannot delete user")
	}
	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
		"user":    user,
	})

}
