package rest

import (
	"net/http"

	"github.com/hasbyadam/account-service/domain"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Remark string `json:"remark"`
}

// AccountService represent the Account's usecases
type AccountService interface {
	CreateAccount(nama, nik, noHP string) (string, error)
	Transaction(accountNo string, nominal float64) (float64, error)
	GetSaldo(accountNo string) (float64, error)
}

// AccountHandler represent the httphandler for Account
type AccountHandler struct {
	Service AccountService
}

// NewAccountHandler will initialize the Accounts/ resources endpoint
func NewAccountHandler(e *echo.Echo, svc AccountService) {
	handler := &AccountHandler{
		Service: svc,
	}
	e.POST("/daftar", handler.CreateAccount)
	e.POST("/tabung", handler.Deposit)
	e.POST("/tarik", handler.Withdraw)
	e.GET("/saldo/:no_rekening", handler.GetSaldo)
}

// CreateAccount handles the creation of a new account
func (h *AccountHandler) CreateAccount(c echo.Context) error {
	var account domain.Account
	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Remark: err.Error()})
	}

	noRekening, err := h.Service.CreateAccount(account.Nama, account.NIK, account.NoHP)
	if err != nil {
		logrus.Errorf("Failed to create account: %v", err)
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"no_rekening": noRekening})
}

// Withdraw handles the withdrawal transaction for an account
func (h *AccountHandler) Withdraw(c echo.Context) error {
	type TransactionRequest struct {
		NoRekening string  `json:"no_rekening"`
		Nominal    float64 `json:"nominal"`
	}

	var req TransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Remark: err.Error()})
	}

	saldo, err := h.Service.Transaction(req.NoRekening, req.Nominal*-1)
	if err != nil {
		logrus.Errorf("Failed to process transaction: %v", err)
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, map[string]float64{"saldo": saldo})
}

// Withdraw handles the withdrawal transaction for an account
func (h *AccountHandler) Deposit(c echo.Context) error {
	type TransactionRequest struct {
		NoRekening string  `json:"no_rekening"`
		Nominal    float64 `json:"nominal"`
	}

	var req TransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Remark: err.Error()})
	}

	saldo, err := h.Service.Transaction(req.NoRekening, req.Nominal)
	if err != nil {
		logrus.Errorf("Failed to process transaction: %v", err)
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, map[string]float64{"saldo": saldo})
}

// GetSaldo handles retrieving the saldo for an account
func (h *AccountHandler) GetSaldo(c echo.Context) error {
	noRekening := c.Param("no_rekening")

	saldo, err := h.Service.GetSaldo(noRekening)
	if err != nil {
		logrus.Errorf("Failed to retrieve saldo for no_rekening %s: %v", noRekening, err)
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, map[string]float64{"saldo": saldo})
}

// handleServiceError handles service errors and maps them to appropriate HTTP responses
func handleServiceError(c echo.Context, err error) error {
    statusCode, ok := domain.ErrorStatusCodeMap[err]
    if !ok {
        statusCode = http.StatusInternalServerError
    }
    return c.JSON(statusCode, ResponseError{Remark: err.Error()})
}
