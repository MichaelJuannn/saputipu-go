package userService

import (
	"fiber-saputipu/database"
	models "fiber-saputipu/internals/model"
	"fiber-saputipu/internals/utils"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.User)
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err})
	}
	user.Password, _ = utils.HashPassword(user.Password)
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(200).JSON(fiber.Map{"status": err})
	}
	return c.Status(200).JSON(&user)

}
