package service

import (
	"bytes"
	"encoding/csv"
	"github.com/JoniDG/transact-mail/internal/defines"
	"github.com/JoniDG/transact-mail/internal/domain"
	"github.com/JoniDG/transact-mail/internal/repository"
	_ "github.com/joho/godotenv/autoload"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type FileService interface {
	HandlerFile(fileName string) error
}

type fileService struct {
	repoEmail       repository.EmailRepository
	repoTransaction repository.TransactionRepository
}

func NewFileService(repoEmail repository.EmailRepository, transactionRepo repository.TransactionRepository) FileService {
	return &fileService{
		repoEmail:       repoEmail,
		repoTransaction: transactionRepo,
	}
}

func (c *fileService) HandlerFile(fileName string) error {
	totalCredit := 0.0
	totalDebit := 0.0
	cantCredit := 0
	cantDebit := 0
	transCountByMonth := make(map[string]int)
	file, err := os.Open(fileName)

	if err != nil {
		return err
	}
	defer file.Close()
	rows, err := ReadFile(file)
	if err != nil {
		return err
	}
	for _, row := range *rows {
		transaction, err := domain.RowFileToUserTransactions(row)
		if err != nil {
			log.Println(err)
			continue
		}
		err = c.repoTransaction.Create(transaction.ToEntity())
		if err != nil {
			log.Println(err)
		}
		month := transaction.Date.Month().String()
		transCountByMonth[month]++
		if transaction.IsCredit {
			totalCredit += transaction.Transaction
			cantCredit += 1
		}
		if transaction.IsDebit {
			totalDebit -= transaction.Transaction
			cantDebit += 1
		}
	}
	emailData := domain.UserTransactionToEmail(totalCredit, totalDebit, cantCredit, cantDebit, transCountByMonth)
	fp := filepath.Join("./templates/body.html")
	t, err := template.ParseFiles(fp)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, emailData)
	if err != nil {
		return err
	}
	email := domain.Email{
		To: []string{os.Getenv(defines.EnvEmailTo)},
		Body: domain.BodyMail{
			Headers: "From: " + os.Getenv(defines.EnvEmailFrom) + "\n" +
				"To: " + os.Getenv(defines.EnvEmailTo) + "\n" +
				defines.SubjectEmail + "\n" +
				defines.Mime + "\r\n",
			Message: buf.String(),
		},
	}
	err = c.repoEmail.Send(email)
	return err
}

func ReadFile(file *os.File) (*[][]string, error) {
	reader := csv.NewReader(file)
	reader.Comma = ','
	rows, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}
	return &rows, nil
}
