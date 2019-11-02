package songs

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Deps struct {
	DB DBLibrary
}

type DBLibrary interface {
	FindAll(string) ([]bson.M, error)
	Find(string, bson.M) ([]bson.M, error)
	FindOne(string, bson.M) (bson.M, error)
	Insert(string, interface{}) error
	UpdateOne(string, interface{}, interface{}) error
	DeleteMany(string, interface{}) (*mongo.DeleteResult, error)
	Aggregate(string, []bson.M) ([]bson.M, error)
}
