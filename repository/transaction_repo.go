package repository

import (
	"graded-challange-1-ysmnababan/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepo interface {
	GetTransactions() ([]models.Transaction, error)
	GetTransactionID(t_ID string) (models.Transaction, error)
	CreateTransaction(t models.Transaction) (interface{}, error)
	UpdateTransaction(t_id string, t models.Transaction) (interface{}, error)
	DeleteTransaction(t_id string) (interface{}, error)
}

type Repo struct {
	DB *mongo.Database
}

func (r *Repo) GetTransactions() ([]models.Transaction, error) {

	return nil, nil
}

func (r *Repo) GetTransactionID(t_ID string) (models.Transaction, error) {

	return models.Transaction{}, nil
}

func (r *Repo) CreateTransaction(t models.Transaction) (interface{}, error) {

	return nil, nil
}

func (r *Repo) UpdateTransaction(t_id string, t models.Transaction) (interface{}, error) {

	return nil, nil
}

func (r *Repo) DeleteTransaction(t_id string) (interface{}, error) {

	return nil, nil
}
