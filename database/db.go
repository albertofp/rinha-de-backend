package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Database

func InitDB() error {
	uri := "mongodb://localhost:5432"
	client, err := mongo.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal("Error connecting to db: ", err)
	}

	db = client.Database("rinhadb")
	fmt.Printf("Connected to Database at %s", uri)
	return nil
}

func CloseDB() error {
	return db.Client().Disconnect(context.Background())
}

func GetCollection(col string) *mongo.Collection {
	return db.Collection(col)
}
