package domain

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type UserTransaction struct {
	ID          int64
	Date        time.Time
	Transaction float64
	IsCredit    bool
	IsDebit     bool
}

type UserTransactionEntity struct {
	ID          int64   `db:"id"`
	Date        string  `db:"date"`
	Transaction float64 `db:"transaction"`
}

func RowFileToUserTransactions(row []string) (*UserTransaction, error) {
	var isCredit, isDebit bool
	var amountStr string
	id, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		return nil, err
	}
	dateStr := row[1]
	date, err := time.Parse("02/01", dateStr)
	if err != nil {
		log.Println(err)
	}

	if strings.Contains(row[2], "+") {
		amountStr = strings.TrimPrefix(row[2], "+")
		isCredit = true
		isDebit = false
	}
	if strings.Contains(row[2], "-") {
		amountStr = strings.TrimPrefix(row[2], "-")
		isCredit = false
		isDebit = true
	}
	transaction, err := strconv.ParseFloat(amountStr, 64)
	return &UserTransaction{
		ID:          id,
		Date:        date,
		IsCredit:    isCredit,
		IsDebit:     isDebit,
		Transaction: transaction,
	}, nil
}

func (ut *UserTransaction) ToEntity() *UserTransactionEntity {
	return &UserTransactionEntity{
		ID:          ut.ID,
		Date:        ut.Date.Format("02/01"),
		Transaction: ut.Transaction,
	}
}
