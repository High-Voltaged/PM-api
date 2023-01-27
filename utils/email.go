package utils

import (
	"api/config"
	"fmt"
	"net/smtp"
)

type EmailBody struct {
	to      string
	from    string
	subject string
	data    string
}

func SendEmail(cfg *config.Email, to string, subject string, data string) error {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	from := cfg.From

	body := EmailBody{
		to:      to,
		from:    from,
		subject: subject,
		data:    data,
	}

	message := BuildMessage(cfg, &body)

	err := smtp.SendMail(addr, auth, from, []string{to}, []byte(message))

	return err
}

func BuildMessage(cfg *config.Email, body *EmailBody) string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", body.from)
	message += fmt.Sprintf("To: %s\r\n", body.to)
	message += fmt.Sprintf("Subject: %s\r\n", body.subject)
	message += fmt.Sprintf("\r\n%s\r\n", body.data)
	return message
}
