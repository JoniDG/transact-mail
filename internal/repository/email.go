package repository

import (
	"github.com/JoniDG/transact-mail/internal/defines"
	"github.com/JoniDG/transact-mail/internal/domain"
	"log"
	"net/smtp"
	"os"
)

type EmailRepository interface {
	Send(email domain.Email) error
}

type emailRepository struct {
	auth smtp.Auth
	host string
	from string
}

func NewEmailRepository(auth smtp.Auth) EmailRepository {
	return &emailRepository{
		auth: auth,
	}
}

func (r *emailRepository) Send(email domain.Email) error {
	addr := os.Getenv(defines.EnvEmailHost) + ":" + os.Getenv(defines.EnvEmailPort)
	body := []byte(email.Body.Headers + email.Body.Message)
	err := smtp.SendMail(addr, r.auth, os.Getenv(defines.EnvSenderUser), email.To, body)
	if err != nil {
		return err
	}
	log.Println("Email Enviado")
	return nil
}
