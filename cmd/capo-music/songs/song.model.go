package songs

// SongModel represents the song entity in the db.
type SongModel struct {
	Title    string      `json:"title"`
	Slug     string      `json:"slug"`
	Visible  bool        `json:"visible"`
	Lyrics   string      `json:"lyrics"`
	Category []string    `json:"category"`
	Media    mediaModel  `json:"media"`
	Composer artistModel `json:"composer"`
	Singer   artistModel `json:"singer"`
}

type mediaModel struct {
	AudioLink string `json:"audioLink"`
	VideoLink string `json:"videoLink"`
}

type artistModel struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
