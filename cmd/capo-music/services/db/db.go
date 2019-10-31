package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	ctx = context.Background()
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/%s", opts.Host, opts.Port, opts.DBName))
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		fmt.Println(err)
		return
	}
	db := client.Database(opts.DBName)
	fmt.Println("Connected to DB: ", db.Name())
	service = &DbService{db}
}

func Service() *DbService {
	return service
}

func (dbService *DbService) FindAll(collection string) ([]bson.M, error) {
	return dbService.Find(collection, bson.M{})
}

func (dbService *DbService) Find(collection string, query bson.M) ([]bson.M, error) {
	coll := dbService.getColl(collection)

	results := []bson.M{}
	cursor, err := coll.Find(ctx, query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for cursor.Next(ctx) {
		var result bson.M
		cursor.Decode(&result)
		results = append(results, result)
	}
	return results, nil
}

func (dbService *DbService) getColl(collection string) *mongo.Collection {
	if dbService.db == nil {
		panic("DB connection not initialized")
	}
	ctx = context.Background()
	return dbService.db.Collection(collection)
}

func (dbService *DbService) FindOne(collection string, query bson.M) (bson.M, error) {
	coll := dbService.getColl(collection)

	result := bson.M{}
	err := coll.FindOne(ctx, query).Decode(&result)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}
