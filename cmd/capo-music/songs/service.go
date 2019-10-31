package songs

type Deps struct {
	DB DBLibrary
}

type SongService struct {
	dao *SongDAO
}

func NewSongService(deps Deps) *SongService {
	dao := NewSongDAO(deps.DB)
	return &SongService{dao}
}

// Find retrieves all songs in the DB.
func (this *SongService) Find(search string) (*[]SongModel, error) {
	return this.dao.find(search)
}

// FindSongBySlug retrieves a specific song in the DB
func (this *SongService) FindSongBySlug(slug string) (*SongModel, error) {
	return this.dao.findByKey(map[string]string{"slug": slug})
}
