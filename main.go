package main

import (
	"context"
	"graded-challange-1-ysmnababan/config"
	"graded-challange-1-ysmnababan/helper"
	"graded-challange-1-ysmnababan/models"
	"graded-challange-1-ysmnababan/router"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client, db := config.Connect(context.TODO(), "soccer_field_rental_db")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	e := echo.New()

	//initialize endpoint
	router.InitRoutes(e, db)

	// initialize seed data for user and field
	dataSeeding(db)

	// get port from .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func dataSeeding(db *mongo.Database) {
	users := []interface{}{
		models.User{Name: "yoland", Email: "yoland@gmail.com"},
		models.User{Name: "admin", Email: "admin@gmail.com"},
	}

	fields := []interface{}{
		models.Field{Name: "Field 1", Price: 50.0, Status: "Available"},
		models.Field{Name: "Field 2", Price: 60.0, Status: "Available"},
		models.Field{Name: "Field 3", Price: 55.0, Status: "Unavailable"},
	}

	res, err := db.Collection("users").InsertMany(context.TODO(), users)
	if err != nil {
		helper.Logging(nil).Error(err)
	}
	helper.Logging(nil).Info("USER: ", res)

	res, err = db.Collection("fields").InsertMany(context.TODO(), fields)
	if err != nil {
		helper.Logging(nil).Error(err)
	}
	helper.Logging(nil).Info("FIELDS: ", res)
}
