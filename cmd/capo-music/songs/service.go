package songs

var dao *SongDAO = NewSongDAO()

type SongsService struct{}

// NewSongsService creates a new SongsService with the given user DAO.
func NewSongsService() *SongsService {
	return &SongsService{}
}

// GetAll just retrieves user using User DAO, here can be additional logic for processing data retrieved by DAOs
func GetAll() (*[]SongModel, error) {
	return dao.getAll()
}
