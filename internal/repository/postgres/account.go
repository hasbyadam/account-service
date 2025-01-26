package postgres

import "database/sql"

type AccountRepository struct {
	DB *sql.DB
}

// NewMysqlAccountRepository will create an implementation of Account.Repository
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}
