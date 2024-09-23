package main

import (
	"log"
	"mailer-microservice/controllers"

	"github.com/gofiber/fiber/v2"
)

// Handler é a função que será chamada pelo Vercel
func Handler(c *fiber.Ctx) error {
	// Chama o controlador para enviar o email
	return controllers.SendMailHandler(c)
}

// Função principal para configurar e rodar o Fiber
func main() {
	app := fiber.New()

	// Define a rota para enviar email
	app.Post("/send-mail", Handler)

	// O Vercel espera que você retorne um manipulador HTTP
	log.Fatal(app.Listen(":3000"))
}
