package main

import (
	"log"
	// "mailer-microservice/config"
	"mailer-microservice/controllers"
	"mailer-microservice/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Cria uma nova instância do Fiber
	app := fiber.New()

	// Inicializa a configuração
	// config.Init()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to mailer!")
	})
	app.Use(middleware.AuthMiddleware)
	app.Post("/send-mail", controllers.SendMailHandler)

	log.Println("Server running")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3010"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
