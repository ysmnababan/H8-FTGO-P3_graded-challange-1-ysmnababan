package router

import (
	"graded-challange-1-ysmnababan/controller"
	"graded-challange-1-ysmnababan/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(e *echo.Echo, db *mongo.Database) {
	e.Use(middleware.Recover())

	repo := &repository.Repo{DB: db}

	// transaction controller init
	tc := &controller.TransactionController{TR: repo}
	uc := &controller.UserController{UR: repo}
	fc := &controller.FieldController{FR: repo}

	e.GET("/transactions", tc.GetAllTransactions)
	e.GET("/transaction/:id", tc.GetTransactionsByID)
	e.POST("/transaction", tc.CreateTransaction)
	e.PUT("/transaction/:id", tc.EditTransaction)
	e.DELETE("/transaction/:id", tc.DeleteTransaction)

	e.GET("/users", uc.GetAllUsers)
	e.GET("/fields", fc.GetAllFields)
}
