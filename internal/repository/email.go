package repository

import (
	"fmt"
	"github.com/JoniDG/transact-mail/internal/domain"
	"net/smtp"
)

type EmailRepository interface {
	Send(transactions []*domain.UserTransaction, totalCredit float64, totalDebit float64, cantCredit int, cantDebit int) error
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

func (r *emailRepository) Send(transactions []*domain.UserTransaction, totalCredit float64, totalDebit float64, cantCredit int, cantDebit int) error {
	fmt.Printf("Total Balance: %.2f\n", totalCredit+totalDebit)
	averageCredit := totalCredit / float64(cantCredit)
	averageDebit := totalDebit / float64(cantDebit)
	fmt.Printf("Average Debit Amount: %.2f\n", averageDebit)
	fmt.Printf("Average Credit Amount: %.2f\n", averageCredit)
	return nil
}
