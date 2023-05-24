package service

import (
	"encoding/csv"
	"fmt"
	"github.com/JoniDG/transact-mail/internal/domain"
	"github.com/JoniDG/transact-mail/internal/repository"
	"log"
	"os"
)

type FileService interface {
	HandlerFile(fileName string) error
}

type fileService struct {
	repoEmail repository.EmailRepository
}

func NewFileService(repoEmail repository.EmailRepository) FileService {
	return &fileService{
		repoEmail: repoEmail,
	}
}

func (c *fileService) HandlerFile(fileName string) error {
	var transactions []*domain.UserTransaction
	//Mapa para el conteo de transacciones por mes
	transCountByMonth := make(map[string]int)

	totalCredit := 0.0
	totalDebit := 0.0
	cantCredit := 0
	cantDebit := 0

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
		transact, err := domain.RowFileToUserTransactions(row)
		if err != nil {
			log.Println(err)
			continue
		}
		month := transact.Date.Month().String()
		transCountByMonth[month]++
		if transact.IsCredit {
			totalCredit += transact.Transaction
			cantCredit += 1
		}
		if transact.IsDebit {
			totalDebit -= transact.Transaction
			cantDebit += 1
		}
		transactions = append(transactions, transact)
	}
	for month, count := range transCountByMonth {
		fmt.Printf("Number of transactions in %s: %d\n", month, count)
	}
	err = c.repoEmail.Send(transactions, totalCredit, totalDebit, cantCredit, cantDebit)
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
