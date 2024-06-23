package main

import (
	"context"
	"graded-challange-1-ysmnababan/config"
	"graded-challange-1-ysmnababan/router"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	// get port from .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
