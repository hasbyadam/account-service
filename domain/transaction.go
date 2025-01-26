package domain

import (
	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Nominal   float64   `json:"nominal" db:"nominal"`
	Type      string    `json:"type" db:"type"`
	CreatedAt int64     `json:"created_at" db:"created_at"`
	AccountID uuid.UUID `json:"account_id" db:"account_id"`
}
