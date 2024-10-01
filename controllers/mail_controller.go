package controllers

import (
	"encoding/json"
	"mailer-microservice/models"
	"mailer-microservice/services"

	"github.com/gofiber/fiber/v2"
)

func SendMailHandler(c *fiber.Ctx) error {
	var mail models.Mail
	err := json.Unmarshal(c.Body(), &mail)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	err = services.SendMail(mail)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send mail"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Email sent successfully"})
}
