package main

import (
	"log"
	"mailer-microservice/controllers"

	"github.com/gofiber/fiber/v2"
)

func SendMailHandler(c *fiber.Ctx) error {
	return controllers.SendMailHandler(c)
}

func main() {
	app := fiber.New()

	// Define a rota "/" para retornar "Hello, World!"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/send-mail", SendMailHandler)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
