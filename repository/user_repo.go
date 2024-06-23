package repository

import (
	"context"
	"graded-challange-1-ysmnababan/helper"
	"graded-challange-1-ysmnababan/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	GetAllUser() ([]models.User, error)
}

func (r *Repo) GetAllUser() ([]models.User, error) {
	var users []models.User
	cursor, err := r.DB.Collection("users").Find(context.TODO(), bson.D{{}})
	if err != nil {
		helper.Logging(nil).Error(err)
		return nil, helper.ErrQuery
	}

	for cursor.Next(context.TODO()) {
		var u models.User
		if err := cursor.Decode(&u); err != nil {
			helper.Logging(nil).Error(err)
			return nil, helper.ErrQuery
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *Repo) IsUserExist(user_id primitive.ObjectID) (bool, error) {
	var result bson.M
	err := r.DB.Collection("users").FindOne(context.TODO(), bson.M{"_id": user_id}).Decode(&result)
	if err != nil {
		// user not found
		if err == mongo.ErrNoDocuments {
			return false, helper.ErrNoUser
		}
		helper.Logging(nil).Error(err)
		return false, err
	}
	return true, nil
}

func (r *Repo) GetUser(user_id primitive.ObjectID) (models.User, error) {
	return models.User{}, nil
}
