package controller

import (
	"zxk/go-fiber-demo/models"

	"github.com/gofiber/fiber"
)

func UserList(ctx *fiber.Ctx) error {
	users, _ := models.GetUsers(&models.User{})
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"list": users,
	})
	return nil
}

func UserProfile(c *fiber.Ctx) error {
	type ID struct {
		Id int `json:"id"`
	}
	var id ID
	c.BodyParser(&id)
	user, err := models.GetUser(id)
	if err != nil {
		c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"message": "content not found",
		})
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
	return err
}

func Login(c *fiber.Ctx) error {

	return nil
}
