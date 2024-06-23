package controller

import (
	"graded-challange-1-ysmnababan/repository"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TR repository.TransactionRepo
}

func (c *TransactionController) GetAllTransactions(e echo.Context) error {

	return nil
}
 
func (c *TransactionController) GetTransactionsByID(e echo.Context) error {

	return nil
}

func (c *TransactionController) CreateTransaction(e echo.Context) error {

	return nil
}

func (c *TransactionController) EditTransaction(e echo.Context) error {

	return nil
}

func (c *TransactionController) DeleteTransaction(e echo.Context) error {

	return nil
}
