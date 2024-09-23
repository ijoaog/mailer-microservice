package services

import (
	// "fmt"
	"mailer-microservice/config"
	"mailer-microservice/models"
	"net/smtp"
)

// SendMail envia um email formatado em HTML com uma apresentação visual melhor
func SendMail(mail models.Mail) error {
	cfg := config.LoadConfig()

	// Configuração de autenticação
	auth := smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost)
	// appSender := "relatify"
	// Destinatário
	to := []string{mail.To}

	// Corpo do e-mail em HTML com formatação melhorada

	// Monta a mensagem do e-mail
	msg := []byte(
		"From: " + cfg.SMTPUser + "\r\n" +
			"To: " + mail.To + "\r\n" +
			"Subject: " + mail.Subject + "\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n" +
			"\r\n" + mail.Body,
	)

	// Envia o e-mail
	err := smtp.SendMail(cfg.SMTPHost+":"+cfg.SMTPPort, auth, cfg.SMTPUser, to, msg)
	if err != nil {
		return err
	}
	return nil
}
