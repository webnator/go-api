package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const (
	// Name of the database.
	DBName = "test"
	URI = "mongodb://localhost:27017/test"
)
var db *mongo.Database
var ctx context.Context


type Test struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title     string             `json:"title"`
}


func Connect() {
	// Base context.
	ctx = context.Background()    // Options to the database.
	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
			fmt.Println(err)
			return
	}
	db = client.Database(DBName)
	fmt.Println(db.Name()) // output: glottery
}

func GetAll(collection string) {
	coll := db.Collection(collection)

	notesResult := []Test{}
	n := Test{}

	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(ctx) {
		cursor.Decode(&n)
		notesResult = append(notesResult, n)
	}

	for _, el := range notesResult {
		fmt.Println(el.Title)
	}

}


