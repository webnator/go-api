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

func (dao *SongDAO) find(search string) (*[]SongModel, error) {
	var songs []SongModel = make([]SongModel, 0)

	query := bson.M{}
	if search != "" {
		query = bson.M{
			"$text": bson.M{
				"$search": search,
			},
		}
	}
	results, err := dao.db.Find("songs", query)

	for _, song := range results {
		var s SongModel = NewSongModel()
		bsonBytes, _ := bson.Marshal(song)

		bson.Unmarshal(bsonBytes, &s)
		songs = append(songs, s)
	}

	return &songs, err
}

func (dao *SongDAO) findByKey(query map[string]string) (*SongModel, error) {
	bsonQuery := bson.M{}
	bsonQueryBytes, _ := bson.Marshal(query)
	bson.Unmarshal(bsonQueryBytes, &bsonQuery)

	result, err := dao.db.FindOne("songs", bsonQuery)

	if result == nil {
		return nil, nil
	}
	var song SongModel
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &song)

	return &song, err
}
