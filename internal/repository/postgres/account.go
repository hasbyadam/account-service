package postgres

import (
	"database/sql"

	"github.com/hasbyadam/account-service/domain"
	"github.com/sirupsen/logrus"
)

type AccountRepository struct {
	DB *sql.DB
}

// NewAccountRepository will create an implementation of Account.Repository
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

// InsertAccount inserts a new account into the database
func (r *AccountRepository) InsertAccount(account *domain.Account) error {
	query := `
        INSERT INTO satrio.accounts (id, nama, nik, no_hp, no_rekening, created_at, updated_at, saldo)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `
	_, err := r.DB.Exec(query, account.ID, account.Nama, account.NIK, account.NoHP, account.NoRekening, account.CreatedAt, account.UpdatedAt, account.Saldo)
	if err != nil {
		logrus.Errorf("Failed to insert account: %v", err)
		return err
	}
	logrus.Infof("Successfully inserted account with ID: %s", account.ID)
	return nil
}

// UpdateSaldoByNoRekening updates the saldo of an account by no_rekening and returns the updated saldo
func (r *AccountRepository) UpdateSaldoByNoRekening(tx *sql.Tx, noRekening string, nominal float64) (float64, error) {
	query := `
        UPDATE satrio.accounts
        SET saldo = saldo + $1
        WHERE no_rekening = $2
        RETURNING saldo
    `
	var updatedSaldo float64
	err := tx.QueryRow(query, nominal, noRekening).Scan(&updatedSaldo)
	if err != nil {
		if err == sql.ErrNoRows {
			logrus.Errorf("Account number not found: %s", noRekening)
			return 0, domain.ErrAccountNoNotFound
		}
		logrus.Errorf("Failed to update saldo for no_rekening %s: %v", noRekening, err)
		return 0, err
	}
	if updatedSaldo < 0 {
		logrus.Errorf("Insufficient saldo for no_rekening %s", noRekening)
		return 0, domain.ErrInsufficientSaldo
	}
	logrus.Infof("Successfully updated saldo for no_rekening: %s", noRekening)
	return updatedSaldo, nil
}

// GetSaldoByNoRekening retrieves the saldo of an account by no_rekening
func (r *AccountRepository) GetSaldoByNoRekening(noRekening string) (float64, error) {
    query := `
        SELECT saldo
        FROM satrio.accounts
        WHERE no_rekening = $1
    `
    var saldo float64
    err := r.DB.QueryRow(query, noRekening).Scan(&saldo)
    if err != nil {
        if err == sql.ErrNoRows {
            logrus.Errorf("Account number not found: %s", noRekening)
            return 0, domain.ErrAccountNoNotFound
        }
        logrus.Errorf("Failed to retrieve saldo for no_rekening %s: %v", noRekening, err)
        return 0, err
    }
    logrus.Infof("Successfully retrieved saldo for no_rekening: %s", noRekening)
    return saldo, nil
}


