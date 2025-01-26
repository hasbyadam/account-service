package rest

import (
	"github.com/labstack/echo/v4"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// AccountService represent the Account's usecases
type AccountService interface {
}

// AccountHandler  represent the httphandler for Account
type AccountHandler struct {
	Service AccountService
}

// NewAccountHandler will initialize the Accounts/ resources endpoint
func NewAccountHandler(e *echo.Echo, svc AccountService) {
	handler := &AccountHandler{
		Service: svc,
	}
	e.GET("/accounts", handler.Test)
}

func (a *AccountHandler) Test(c echo.Context) error {
	return c.JSON(200, "test ok")
}
