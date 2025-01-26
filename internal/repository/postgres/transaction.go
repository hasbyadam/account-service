package postgres

import (
	"database/sql"

	"github.com/hasbyadam/account-service/domain"
	"github.com/sirupsen/logrus"
)

type TransactionRepository struct {
	DB *sql.DB
}

// NewMysqlTransactionRepository will create an implementation of Transaction.Repository
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

// InsertTransaction inserts a new transaction into the database
func (r *TransactionRepository) InsertTransaction(tx *sql.Tx, transaction *domain.Transaction) error {
	query := `
        INSERT INTO satrio.transactions (id, nominal, type, created_at, account_id)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := r.DB.Exec(query, transaction.ID, transaction.Nominal, transaction.Type, transaction.CreatedAt, transaction.AccountID)
	if err != nil {
		logrus.Errorf("Failed to insert transaction: %v", err)
		return err
	}
	logrus.Infof("Successfully inserted transaction with ID: %s", transaction.ID)
	return nil
}
