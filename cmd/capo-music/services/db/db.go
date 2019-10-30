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

var ctx context.Context

type DbService struct {
	db *mongo.Database
}

type ConnOptions struct {
	DBName string
	Host   string
	Port   string
}

var service *DbService

func Connect(opts ConnOptions) {
	fmt.Println("->", opts.DBName, opts.Host, opts.Port)
	ctx = context.Background()
	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		fmt.Println(err)
		return
	}
	db := client.Database(DBName)
	fmt.Println("Connected to DB: ", db.Name())
	service = &DbService{db}
}

func Service() *DbService {
	return service
}

func (dbService *DbService) GetAll(collection string) ([]bson.M, error) {
	if dbService.db == nil {
		panic("DB connection not initialized")
	}
	ctx = context.Background()
	coll := dbService.db.Collection(collection)

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
