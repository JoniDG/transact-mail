package domain

import "strconv"

type UserTransactions struct {
	ID          uint64 `db:"id"`
	Date        string `db:"name"`
	Transaction string
}

func RowFileToUserTransactions(row []string) (*UserTransactions, error) {
	id, err := strconv.ParseUint(row[0], 10, 64)
	if err != nil {
		return nil, err
	}
	return &UserTransactions{
		ID:          id,
		Date:        row[1],
		Transaction: row[2],
	}, nil
}
