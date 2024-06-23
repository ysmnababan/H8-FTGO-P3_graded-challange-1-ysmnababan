package repository

import (
	"context"
	"graded-challange-1-ysmnababan/helper"
	"graded-challange-1-ysmnababan/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FieldRepo interface {
	GetAllFields() ([]models.Field, error)
}

func (r *Repo) GetAllFields() ([]models.Field, error) {
	var fields []models.Field
	cursor, err := r.DB.Collection("fields").Find(context.TODO(), bson.D{{}})
	if err != nil {
		helper.Logging(nil).Error(err)
		return nil, helper.ErrQuery
	}

	for cursor.Next(context.TODO()) {
		var field models.Field
		if err := cursor.Decode(&field); err != nil {
			helper.Logging(nil).Error(err)
			return nil, helper.ErrQuery
		}
		fields = append(fields, field)
	}
	return fields, nil
}

func (r *Repo) IsFieldExist(field_id primitive.ObjectID) (bool, error) {
	var res bson.M
	err := r.DB.Collection("fields").FindOne(context.TODO(), bson.M{"_id": field_id}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, helper.ErrNoData
		}
		helper.Logging(nil).Error(err)
		return false, helper.ErrQuery
	}
	return true, nil
}

func (r *Repo) GetField(field_id primitive.ObjectID) (models.Field, error) {

	return models.Field{}, nil
}
