package controller

import (
	"graded-challange-1-ysmnababan/helper"
	"graded-challange-1-ysmnababan/models"
	"graded-challange-1-ysmnababan/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionController struct {
	TR repository.TransactionRepo
}

func (c *TransactionController) GetAllTransactions(e echo.Context) error {
	res, err := c.TR.GetTransactions()
	if err != nil {
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{"Message:": "All transaction list", "Data": res})
}

func (c *TransactionController) GetTransactionsByID(e echo.Context) error {
	// get id
	t_id := e.Param("id")
	res, err := c.TR.GetTransactionID(t_id)
	if err != nil {
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{"Message:": "Transaction by a user", "Data": res})
}

func (c *TransactionController) CreateTransaction(e echo.Context) error {
	// get body request
	var req models.Transaction
	err := e.Bind(&req)
	if err != nil {
		helper.Logging(e).Error(err)
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	helper.Logging(e).Info("INPUT DATA :", req)

	// validate request
	if req.FieldID == primitive.NilObjectID || req.UserID == primitive.NilObjectID || req.EndTime == "" || req.StartTime == "" {
		return e.JSON(http.StatusBadRequest, "error or missing param")
	}

	// validate time
	if !helper.IsTimeValid(req.StartTime, req.EndTime) {
		return e.JSON(http.StatusBadRequest, "time is invalid")
	}

	res, err := c.TR.CreateTransaction(&req)
	if err != nil {
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusCreated, res)
}

func (c *TransactionController) EditTransaction(e echo.Context) error {
	// get id
	t_id := e.Param("id")

	// get body request
	var req models.Transaction
	err := e.Bind(&req)
	if err != nil {
		helper.Logging(e).Error(err)
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	helper.Logging(e).Info("INPUT DATA :", req)

	// validate request
	if req.FieldID == primitive.NilObjectID || req.UserID == primitive.NilObjectID || req.EndTime == "" || req.StartTime == "" {
		return e.JSON(http.StatusBadRequest, "error or missing param")
	}

	// validate time
	if !helper.IsTimeValid(req.StartTime, req.EndTime) {
		return e.JSON(http.StatusBadRequest, "time is invalid")
	}

	res, err := c.TR.UpdateTransaction(t_id, &req)
	if err != nil {
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, res)
}

func (c *TransactionController) DeleteTransaction(e echo.Context) error {
	// get id
	t_id := e.Param("id")

	res, err := c.TR.DeleteTransaction(t_id)
	if err != nil {
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, res)
}
