package songs

import (
	"go.mongodb.org/mongo-driver/bson"
)

// SongDAO persists user data in database
type SongDAO struct {
	db DBLibrary
}

// NewSongDAO creates a new UserDAO
func NewSongDAO(dbLib DBLibrary) *SongDAO {
	db := dbLib
	return &SongDAO{db}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *SongDAO) getAll() (*[]SongModel, error) {
	var songs []SongModel = make([]SongModel, 0)

	// Query Database here...
	results, err := dao.db.GetAll("songs")

	for _, song := range results {
		var s SongModel = NewSongModel()
		bsonBytes, _ := bson.Marshal(song)

		bson.Unmarshal(bsonBytes, &s)
		songs = append(songs, s)
	}

	return &songs, err
}
