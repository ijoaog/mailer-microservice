package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	fmt.Println("Token recebido:", token)

	if token == "" {
		fmt.Println("Token ausente")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		fmt.Println("Formato do token inválido:", token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	actualToken := parts[1]
	fmt.Println("Token extraído:", actualToken)
	fmt.Println("Token esperado:", os.Getenv("TOKEN"))

	if actualToken == "" || actualToken != os.Getenv("TOKEN") {
		fmt.Println("Token inválido ou ausente:", actualToken)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.Next()
}
