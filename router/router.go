package router

import (
	bankaccountservice "fiber-saputipu/internals/bank_account"
	predictionService "fiber-saputipu/internals/prediction_text"
	reportService "fiber-saputipu/internals/report"
	userService "fiber-saputipu/internals/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// bank account routes
	bankaccount := api.Group("/bank-account")
	bankaccount.Get("/:id", bankaccountservice.Account)

	// prediction routes
	prediction := api.Group("/prediction-text")
	prediction.Post("/", predictionService.Create)

	// user routes
	user := api.Group("user")
	user.Post("/", userService.Create)

	// report routes
	report := api.Group("report")
	report.Post("/", reportService.Create)
}
