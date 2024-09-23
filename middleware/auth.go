package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware verifica se o token de autenticação está presente
func AuthMiddleware(c *fiber.Ctx) error {
	// Exemplo de verificação de token
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Se o token estiver presente, continue com a próxima função
	return c.Next()
}
