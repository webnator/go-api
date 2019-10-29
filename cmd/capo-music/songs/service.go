package songs

var dao *SongDAO = NewSongDAO()

// GetAll retrieves all songs in the DB.
func GetAll() (*[]SongModel, error) {
	return dao.getAll()
}
