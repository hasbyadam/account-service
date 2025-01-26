package domain

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Nama       string    `json:"nama" db:"nama"`
	NIK        string    `json:"nik" db:"nik" unique:"true"`
	NoHP       string    `json:"no_hp" db:"no_hp" unique:"true"`
	NoRekening string    `json:"no_rekening" db:"no_rekening" unique:"true"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at,omitempty"`
	Saldo      float64   `json:"saldo" db:"saldo"`
}
