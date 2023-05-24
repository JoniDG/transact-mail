package main

import (
	"github.com/JoniDG/transact-mail/internal/controller"
	"github.com/JoniDG/transact-mail/internal/repository"
	"github.com/JoniDG/transact-mail/internal/service"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
)

func main() {

	from := os.Getenv("EMAIL_SENDER_USER")
	password := os.Getenv("EMAIL_SENDER_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	auth := smtp.PlainAuth("", from, password, host)

	// Repositories init
	emailRepo := repository.NewEmailRepository(auth)

	// Services init
	svc := service.NewFileService(emailRepo)

	// Controllers init
	ctrl := controller.NewFileController(svc)

	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	csvPath := filepath.Join(dir, "files", "example.csv")
	err = ctrl.HandlerFile(csvPath)
	if err != nil {
		log.Println(err)
	}
}
