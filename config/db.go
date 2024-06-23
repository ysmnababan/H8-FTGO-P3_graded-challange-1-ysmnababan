package config

import (
	"context"
	"fmt"

	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, dbname string) (*mongo.Client, *mongo.Database) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://ysmnababan:9S3o8HFpBtYKpeRk@devcluster.djcnk8z.mongodb.net/?retryWrites=true&w=majority&appName=DevCluster").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	db := client.Database(dbname)
	if err := db.RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, []any{"Pinged your deployment. You successfully connected to MongoDB!"}...)

	return client, db
}
