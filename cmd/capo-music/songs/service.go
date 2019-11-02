package songs

import "github.com/stvp/slug"

type SongService struct {
	dao *SongDAO
}

func NewSongService(deps Deps) *SongService {
	dao := NewSongDAO(deps.DB)
	return &SongService{dao}
}

// Find retrieves all songs in the DB.
func (service *SongService) Find(params map[string]string) (*[]SongModel, error) {
	return service.dao.find(params)
}

// FindSongBySlug retrieves a specific song in the DB
func (service *SongService) FindSongBySlug(slug string) (*SongModel, error) {
	return service.dao.findByKey(map[string]string{"slug": slug})
}

func (service *SongService) GetSongCategories() ([]string, error) {
	return service.dao.getCategories()
}

func (service *SongService) AddSong(song SongModel) error {
	if song.Slug == "" {
		song.Slug = setSongSlug(song.Title)
	}
	return service.dao.storeSongInfo(song)
}

func (service *SongService) UpdateSong(slug string, song SongUpdateModel) error {
	return service.dao.updateSongInfo(slug, song)
}

func (service *SongService) IncreaseViewCount(slug string) error {
	return service.dao.addViewToSong(slug)
}

func (service *SongService) ResetViews(params ResetViewsModel) error {
	return service.dao.resetSongViews(params)
}

func setSongSlug(name string) string {
	slug.Replacement = '-'
	return slug.Clean(name)
}
