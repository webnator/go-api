package songs

import (
	"github.com/webnator/capo-music-api/cmd/capo-music/services/db"
)

// SongDAO persists user data in database
type SongDAO struct{}

// NewSongDAO creates a new UserDAO
func NewSongDAO() *SongDAO {
	return &SongDAO{}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *SongDAO) Get(id uint) (*SongModel, error) {
	var songs []SongModel

	// Query Database here...
	results, err := db.GetAll("songs")

	for _, song := range results {
		append(songs, song)
	}

	return &songs, err
}
