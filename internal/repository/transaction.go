package repository

import (
	"github.com/JoniDG/transact-mail/internal/domain"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Create(ut *domain.UserTransactionEntity) error
}

type transactionRepository struct {
	db         *sqlx.DB
	sqlBuilder transactionsSQL
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
		sqlBuilder: transactionsSQL{
			table: "core.transactions",
		},
	}
}

type transactionsSQL struct {
	table string
}

func (r *transactionRepository) Create(ut *domain.UserTransactionEntity) error {
	query, args, err := r.sqlBuilder.CreateSQL(ut)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	return err
}

func (s *transactionsSQL) CreateSQL(t *domain.UserTransactionEntity) (string, []interface{}, error) {
	query, args, err := squirrel.Insert(s.table).
		Columns("id", "date", "transaction").
		Values(t.ID, t.Date, t.Transaction).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	return query, args, err
}
