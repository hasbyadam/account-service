package account

// AccountRepository represent the Account's repository contract
type AccountRepository interface {
}

// TransactionRepository represent the Transaction's repository contract
type TransactionRepository interface {
}

type Service struct {
	AccountRepo     AccountRepository
	TransactionRepo TransactionRepository
}

// NewService will create a new Account service object
func NewService(a AccountRepository, ar TransactionRepository) *Service {
	return &Service{
		AccountRepo:     a,
		TransactionRepo: ar,
	}
}
