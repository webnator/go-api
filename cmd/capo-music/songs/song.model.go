package songs

// NewSongModel represents the song entity in the db.
func NewSongModel() SongModel {
	song := SongModel{}
	song.Category = make([]string, 0)
	song.Tags = make([]string, 0)
	return song
}

type SongModel struct {
	Title    string      `json:"title" bson:"title,omitempty" binding:"required"`
	Slug     string      `json:"slug" bson:"slug,omitempty"`
	Hidden   bool        `json:"hidden" bson:"hidden"`
	Lyrics   string      `json:"lyrics" bson:"lyrics,omitempty" binding:"required"`
	Category []string    `json:"category" bson:"category,omitempty" binding:"required,gt=0"`
	Tags     []string    `json:"tags" bson:"tags,omitempty"`
	Media    mediaModel  `json:"media" bson:"media,omitempty"`
	Composer artistModel `json:"composer" bson:"composer,omitempty"`
	Singer   artistModel `json:"singer" bson:"singer,omitempty"`
	Viewed   int         `json:"viewed" bson:"viewed"`
}

type SongUpdateModel struct {
	Title    string      `json:"title" bson:"title,omitempty"`
	Hidden   bool        `json:"hidden" bson:"hidden"`
	Lyrics   string      `json:"lyrics" bson:"lyrics,omitempty"`
	Category []string    `json:"category" bson:"category,omitempty" binding:"omitempty,gt=0"`
	Tags     []string    `json:"tags" bson:"tags,omitempty"`
	Media    mediaModel  `json:"media" bson:"media,omitempty"`
	Composer artistModel `json:"composer" bson:"composer,omitempty"`
	Singer   artistModel `json:"singer" bson:"singer,omitempty"`
	Viewed   int         `json:"viewed" bson:"viewed,omitempty"`
}

type mediaModel struct {
	AudioLink string `json:"audioLink"`
	VideoLink string `json:"videoLink"`
}

type artistModel struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
