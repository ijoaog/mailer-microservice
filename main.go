package main

import (
	"log"
	"mailer-microservice/config"
	"mailer-microservice/controllers"
	"mailer-microservice/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Cria uma nova instância do Fiber
	app := fiber.New()

	// Inicializa a configuração
	config.Init()

	// Define a rota com o middleware de autenticação
	app.Use(middleware.AuthMiddleware) // Use o middleware para todas as rotas
	app.Post("/send-mail", controllers.SendMailHandler)

	log.Println("Server running")

	// Inicie o servidor (remova isso ao implantar na Vercel)
	if err := app.Listen(":3010"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
