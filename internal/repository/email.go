package repository

import (
	"fmt"
	"github.com/JoniDG/transact-mail/internal/domain"
)

type EmailRepository interface {
	Send(transact *domain.UserTransactions) error
}

type emailRepository struct {
}

func NewEmailRepository() EmailRepository {
	return &emailRepository{}
}

func (r *emailRepository) Send(transact *domain.UserTransactions) error {
	fmt.Println(transact)
	return nil
}
