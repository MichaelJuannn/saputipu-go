package predictionService

import (
	"fiber-saputipu/database"
	models "fiber-saputipu/internals/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreatePrediction(c *fiber.Ctx) error {
	db := database.DB
	prediction := new(models.PredictionText)
	err := c.BodyParser(prediction)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err})
	}
	prediction.CreatedAt = time.Now()
	err = db.Create(&prediction).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err})
	}
	return c.SendString(prediction.Text)
}
