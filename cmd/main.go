package main

import (
	"fmt"
	"github.com/JoniDG/transact-mail/internal/controller"
	"github.com/JoniDG/transact-mail/internal/defines"
	"github.com/JoniDG/transact-mail/internal/repository"
	"github.com/JoniDG/transact-mail/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
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

	postgresURI := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		os.Getenv(defines.EnvPostgresUser),
		os.Getenv(defines.EnvPostgresPassword),
		os.Getenv(defines.EnvPostgresHost),
		os.Getenv(defines.EnvPostgresPort),
	)
	db, err := sqlx.Open("postgres", postgresURI)
	if err != nil {
		log.Panic(err)
	}
	// Repositories init
	emailRepo := repository.NewEmailRepository(auth)
	transactionRepo := repository.NewTransactionRepository(db)
	// Services init
	svc := service.NewFileService(emailRepo, transactionRepo)

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
