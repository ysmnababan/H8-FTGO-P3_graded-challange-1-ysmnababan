package repository

import (
	"graded-challange-1-ysmnababan/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repo) IsFieldExist(field_id primitive.ObjectID) (bool, error) {
	return true, nil
}

func (r *Repo) GetField(field_id primitive.ObjectID) (models.Field, error) {

	return models.Field{}, nil
}
