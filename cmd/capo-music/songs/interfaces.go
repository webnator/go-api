package songs

import (
	"go.mongodb.org/mongo-driver/bson"
)

type DBLibrary interface {
	GetAll(string) ([]bson.M, error)
}
