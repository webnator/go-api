package songs

import (
	"github.com/webnator/capo-music-api/cmd/capo-music/services/db"
	"go.mongodb.org/mongo-driver/bson"
)

// SongDAO persists user data in database
type SongDAO struct{}

// NewSongDAO creates a new UserDAO
func NewSongDAO() *SongDAO {
	return &SongDAO{}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *SongDAO) getAll() (*[]SongModel, error) {
	var songs []SongModel

	// Query Database here...
	results, err := db.GetAll("songs")

	for _, song := range results {
		var s SongModel
		bsonBytes, _ := bson.Marshal(song)

		bson.Unmarshal(bsonBytes, &s)
		songs = append(songs, s)
	}

	return &songs, err
}
