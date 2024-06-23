package repository

import (
	"context"
	"graded-challange-1-ysmnababan/helper"
	"graded-challange-1-ysmnababan/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepo interface {
	GetTransactions() ([]models.Transaction, error)
	GetTransactionID(t_ID string) (models.Transaction, error)
	CreateTransaction(t *models.Transaction) (interface{}, error)
	UpdateTransaction(t_id string, t *models.Transaction) (interface{}, error)
	DeleteTransaction(t_id string) (interface{}, error)
}

type Repo struct {
	DB *mongo.Database
}

func (r *Repo) GetTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction

	cursor, err := r.DB.Collection("transactions").Find(context.TODO(), bson.D{{}})
	if err != nil {
		helper.Logging(nil).Error(err)
		return nil, helper.ErrQuery
	}

	for cursor.Next(context.TODO()) {
		var t models.Transaction
		if err := cursor.Decode(&t); err != nil {
			helper.Logging(nil).Error(err)
			return nil, helper.ErrQuery
		}

		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (r *Repo) GetTransactionID(t_ID string) (models.Transaction, error) {
	id, err := primitive.ObjectIDFromHex(t_ID)
	if err != nil {
		helper.Logging(nil).Error(err)
		return models.Transaction{}, helper.ErrInvalidId
	}

	isExist, err := r.IsTransactionExist(id)
	if err != nil {
		return models.Transaction{}, err
	}

	if !isExist {
		return models.Transaction{}, err
	}

	var t models.Transaction
	err = r.DB.Collection("transactions").FindOne(context.TODO(), bson.D{{}}).Decode(&t)
	if err != nil {
		helper.Logging(nil).Error(err)
		return models.Transaction{}, err
	}
	return t, nil
}

func (r *Repo) CreateTransaction(t *models.Transaction) (interface{}, error) {

	isUserExist, err := r.IsUserExist(t.UserID)
	if err != nil || !isUserExist {
		return nil, err
	}

	isFieldExist, err := r.IsFieldExist(t.FieldID)
	if err != nil || !isFieldExist {
		return nil, err
	}

	res, err := r.DB.Collection("transactions").InsertOne(context.TODO(), *t)
	if err != nil {
		helper.Logging(nil).Error(err)
		return nil, helper.ErrQuery
	}

	return res, nil
}

func (r *Repo) UpdateTransaction(t_id string, t *models.Transaction) (interface{}, error) {
	// check transaction id exist or not
	transaction_id, err := primitive.ObjectIDFromHex(t_id)
	if err != nil {
		helper.Logging(nil).Error(err)
		return nil, helper.ErrInvalidId
	}

	isTransactionExist, err := r.IsTransactionExist(transaction_id)
	if err != nil || !isTransactionExist {
		return nil, err
	}

	isUserExist, err := r.IsUserExist(t.UserID)
	if err != nil || !isUserExist {
		return nil, err
	}

	isFieldExist, err := r.IsFieldExist(t.FieldID)
	if err != nil || !isFieldExist {
		return nil, err
	}

	res, err := r.DB.Collection("transactions").UpdateOne(
		context.TODO(),
		bson.M{"_id": transaction_id},
		bson.M{"$set": t},
	)

	if err != nil {
		helper.Logging(nil).Error(err)
		return nil, helper.ErrQuery
	}

	return res, nil
}

func (r *Repo) DeleteTransaction(t_id string) (interface{}, error) {
	// check transaction id exist or not
	transaction_id, err := primitive.ObjectIDFromHex(t_id)
	if err != nil {
		helper.Logging(nil).Error(err)
		return nil, helper.ErrInvalidId
	}

	isTransactionExist, err := r.IsTransactionExist(transaction_id)
	if err != nil || !isTransactionExist {
		return nil, err
	}
	res, err := r.DB.Collection("transactions").DeleteOne(context.TODO(), bson.M{"_id": transaction_id})
	if err != nil {
		return nil, helper.ErrQuery
	}
	return res, nil
}

func (r *Repo) IsTransactionExist(t_id primitive.ObjectID) (bool, error) {
	var result bson.M
	err := r.DB.Collection("transactions").FindOne(context.TODO(), bson.M{"_id": t_id}).Decode(&result)
	if err != nil {
		// no data found
		if err == mongo.ErrNoDocuments {
			return false, helper.ErrNoData
		}
		helper.Logging(nil).Error(err)
		return false, err
	}
	return true, nil
}
