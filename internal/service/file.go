package service

import (
	"encoding/csv"
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
		err = c.repoEmail.Send(transact)
	}
	return nil
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
