package main

import (
	"fiber-saputipu/database"
	"fiber-saputipu/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()
	app.Get("/", hello)
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("hello go fiber try me air")
}
