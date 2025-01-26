package postgres

import "database/sql"

type TransactionRepository struct {
	DB *sql.DB
}

// NewMysqlTransactionRepository will create an implementation of Transaction.Repository
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}
