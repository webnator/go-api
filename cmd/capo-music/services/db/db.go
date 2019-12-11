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
	URI    string
	DBName string
}

var service *DbService

func Connect(opts ConnOptions) {
	ctx = context.Background()
	clientOpts := options.Client().ApplyURI(opts.URI)
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

func (dbService *DbService) Find(collection string, query bson.M, opts ...*options.FindOptions) ([]bson.M, error) {
	coll := dbService.getColl(collection)

	results := []bson.M{}
	cursor, err := coll.Find(ctx, query, opts[0])

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

func (dbService *DbService) Insert(collection string, record interface{}) error {
	coll := dbService.getColl(collection)

	_, err := coll.InsertOne(ctx, record)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (dbService *DbService) UpdateOne(collection string, record interface{}, newInfo interface{}) error {
	coll := dbService.getColl(collection)

	_, err := coll.UpdateOne(ctx, record, newInfo)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (dbService *DbService) DeleteMany(collection string, filter interface{}) (*mongo.DeleteResult, error) {
	coll := dbService.getColl(collection)

	res, err := coll.DeleteMany(ctx, filter)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}

func (dbService *DbService) Aggregate(collection string, pipeQuery []bson.M) ([]bson.M, error) {
	coll := dbService.getColl(collection)

	results := []bson.M{}
	cursor, err := coll.Aggregate(ctx, pipeQuery)

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
