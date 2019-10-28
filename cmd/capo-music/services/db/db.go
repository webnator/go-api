package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBName = "test"
	URI    = "mongodb://localhost:27017/test"
)

var db *mongo.Database
var ctx context.Context

func Connect() {
	// Base context.
	ctx = context.Background() // Options to the database.
	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		fmt.Println(err)
		return
	}
	db = client.Database(DBName)
	fmt.Println("Conneceted to DB: ", db.Name()) // output: glottery
}

func GetAll(collection string) ([]bson.M, error) {
	if db == nil {
		panic("DB connection not initialized")
	}
	coll := db.Collection(collection)

	results := []bson.M{}
	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Iterate through the returned cursor.
	for cursor.Next(ctx) {
		var result bson.M
		cursor.Decode(&result)
		results = append(results, result)
	}
	return results, nil
}
