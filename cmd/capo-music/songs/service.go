package songs

type songsDAO interface {
	getAll() (*[]SongModel, error)
}

type Deps struct {
	DB DBLibrary
}

type SongService struct {
	dao songsDAO
}

func NewSongService(deps Deps) *SongService {
	dao := NewSongDAO(deps.DB)
	return &SongService{dao}
}

// GetAll retrieves all songs in the DB.
func (this *SongService) GetAll() (*[]SongModel, error) {
	return this.dao.getAll()
}
