package songs

import (
	"go.mongodb.org/mongo-driver/bson"
)

type DBLibrary interface {
	FindAll(string) ([]bson.M, error)
	Find(string, bson.M) ([]bson.M, error)
	FindOne(string, bson.M) (bson.M, error)
}
