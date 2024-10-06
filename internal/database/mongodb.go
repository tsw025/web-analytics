package database

import (
	"context"
	"fmt"
	"github.com/tsw025/web_analytics/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// The use of context is to manage the lifecycle of the connection
// when creating the connection to the database, if the connection is not established within 10 seconds, it will return an error
// The cancel function is called to release resources and close the connection
func connectToMongoDB(cfg *config.Config) (*mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURL))
	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}
	return client, ctx, cancel
}

// This function should be called in the main function to close the connection
// No need intermediate function to close the connection, where MongoDB manages the connection pool
func disconnectMongoDB(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
	defer func() {
		cancel()
		if err := client.Disconnect(context); err != nil {
			panic(err)
		}
		fmt.Println("Close connection is called")
	}()
}
