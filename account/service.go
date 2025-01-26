package account

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/hasbyadam/account-service/domain"
	"github.com/sirupsen/logrus"
)

// AccountRepository represent the Account's repository contract
type AccountRepository interface {
	InsertAccount(account *domain.Account) error
	UpdateSaldoByNoRekening(tx *sql.Tx, noRekening string, nominal float64) (float64, error)
	GetSaldoByNoRekening(noRekening string) (float64, error)
}

// TransactionRepository represent the Transaction's repository contract
type TransactionRepository interface {
	InsertTransaction(tx *sql.Tx, transaction *domain.Transaction) error
}

type Service struct {
	AccountRepo     AccountRepository
	TransactionRepo TransactionRepository
	DB              *sql.DB
}

// NewService will create a new Account service object
func NewService(a AccountRepository, ar TransactionRepository, Db *sql.DB) *Service {
	return &Service{
		AccountRepo:     a,
		TransactionRepo: ar,
		DB:              Db,
	}
}

// CreateAccount receives account details and inserts it into the database
func (s *Service) CreateAccount(nama, nik, noHP, noRekening string, saldo float64) error {
	account := &domain.Account{
		ID:         uuid.New(),
		Nama:       nama,
		NIK:        nik,
		NoHP:       noHP,
		NoRekening: noRekening,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Saldo:      saldo,
	}

	err := s.AccountRepo.InsertAccount(account)
	if err != nil {
		logrus.Errorf("Failed to create account: %v", err)
		return err
	}
	logrus.Infof("Successfully created account with ID: %s", account.ID)
	return nil
}

// Transaction updates the saldo of the owner account_no based on nominal and inserts the transaction into the database
func (s *Service) Transaction(accountNo string, nominal float64) error {
    tx, err := s.DB.Begin()
    if err != nil {
        logrus.Errorf("Failed to begin transaction: %v", err)
        return err
    }

    // Update saldo
    updatedSaldo, err := s.AccountRepo.UpdateSaldoByNoRekening(tx, accountNo, nominal)
    if err != nil {
        tx.Rollback()
        return err
    }

    // Determine transaction type
    transactionType := "credit"
    if nominal < 0 {
        transactionType = "debit"
    }

    // Insert transaction
    transaction := &domain.Transaction{
        ID:        uuid.New(),
        Nominal:   nominal,
        Type:      transactionType,
        CreatedAt: time.Now(),
        AccountID: uuid.MustParse(accountNo),
    }

    err = s.TransactionRepo.InsertTransaction(tx, transaction)
    if err != nil {
        tx.Rollback()
        logrus.Errorf("Failed to insert transaction for account_no %s: %v", accountNo, err)
        return err
    }

    err = tx.Commit()
    if err != nil {
        logrus.Errorf("Failed to commit transaction: %v", err)
        return err
    }

    logrus.Infof("Successfully updated saldo and inserted transaction for account_no: %s, new saldo: %f", accountNo, updatedSaldo)
    return nil
}

// GetSaldo retrieves the saldo of an account by account_no
func (s *Service) GetSaldo(accountNo string) (float64, error) {
    saldo, err := s.AccountRepo.GetSaldoByNoRekening(accountNo)
    if err != nil {
        logrus.Errorf("Failed to retrieve saldo for account_no %s: %v", accountNo, err)
        return 0, err
    }
    logrus.Infof("Successfully retrieved saldo for account_no: %s, saldo: %f", accountNo, saldo)
    return saldo, nil
}
