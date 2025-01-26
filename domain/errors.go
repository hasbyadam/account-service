package domain

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrDuplicateNIKOrNoHP will throw if there is a duplicate nik or no_hp
	ErrDuplicateNIKOrNoHP = errors.New("duplicate NIK or No HP")
	// ErrAccountNoNotFound will throw if the account_no is not found
	ErrAccountNoNotFound = errors.New("account number not found")
	// ErrInsufficientSaldo will throw if there is insufficient saldo
	ErrInsufficientSaldo = errors.New("insufficient saldo")
)
