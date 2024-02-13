package reportService

import (
	models "fiber-saputipu/internals/model"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	report := new(models.Laporan)
	err := c.BodyParser(&report)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err})
	}
	return c.Status(200).JSON(report)
}
