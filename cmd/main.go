package main

import (
	"github.com/JoniDG/transact-mail/internal/controller"
	"github.com/JoniDG/transact-mail/internal/repository"
	"github.com/JoniDG/transact-mail/internal/service"
	"log"
)

func main() {
	// Repositories init
	emailRepo := repository.NewEmailRepository()

	// Services init
	svc := service.NewFileService(emailRepo)

	// Controllers init
	ctrl := controller.NewFileController(svc)

	err := ctrl.HandlerFile("example.csv")
	if err != nil {
		log.Println(err)
	}
}
