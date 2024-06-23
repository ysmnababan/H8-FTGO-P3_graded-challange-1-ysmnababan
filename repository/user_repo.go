package repository

import (
	"graded-challange-1-ysmnababan/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repo) IsUserUnique(user_id primitive.ObjectID) (bool, error) {
	return true, nil
}

func (r *Repo) GetUser(user_id primitive.ObjectID) (models.User, error) {

	return models.User{}, nil
}
