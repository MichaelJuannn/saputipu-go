package bankaccountservice

import "github.com/gofiber/fiber/v2"

func Account(c *fiber.Ctx) error {

	id := c.Params("id")
	return c.SendString(id)
}
