package songs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SongModel struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title string             `json:"title"`
}
