package songs

import (
	"fmt"
	"time"

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

func queryModel(params map[string]string) bson.M {
	query := bson.M{}
	if params["search"] != "" {
		query = bson.M{
			"$text": bson.M{
				"$search": params["search"],
			},
		}
	}
	if params["category"] != "" {
		query["category"] = params["category"]
	}
	return query
}

func (dao *SongDAO) find(params map[string]string) (*[]SongModel, error) {
	var songs []SongModel = make([]SongModel, 0)

	query := queryModel(params)
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

func (dao *SongDAO) getCategories() ([]string, error) {
	var categories []string = make([]string, 0)
	bsonQuery := []bson.M{
		{"$unwind": "$category"},
		{"$project": bson.M{"category": 1}},
		{"$group": bson.M{"_id": "$category"}},
	}

	results, err := dao.db.Aggregate("songs", bsonQuery)

	for _, category := range results {
		categories = append(categories, category["_id"].(string))
	}

	return categories, err
}

func (dao *SongDAO) storeSongInfo(song SongModel) error {
	err := dao.db.Insert("songs", song)
	return err
}

func (dao *SongDAO) updateSongInfo(slug string, song SongUpdateModel) error {
	err := dao.db.UpdateOne("songs", bson.M{"slug": slug}, bson.M{"$set": song})
	return err
}

func (dao *SongDAO) addViewToSong(slug string) error {
	err := dao.db.UpdateOne("songs", bson.M{"slug": slug}, bson.M{"$inc": bson.M{"viewed": 1}})
	if err != nil {
		fmt.Println("Failed to register view to song")
		return err
	}
	err = dao.db.Insert("view_count", bson.M{
		"time": time.Now(),
		"song": slug,
	})
	return err
}

func resetFilterModel(slug string, params ResetViewsModel) bson.M {
	filter := bson.M{}
	var timeFilter bson.M

	if slug != "" {
		filter["song"] = slug
	}
	if params.Since != "" {
		timeFilter = timeModel(params.Since, timeFilter, "$gt")
		filter["time"] = timeFilter
	}
	if params.Until != "" {
		timeFilter = timeModel(params.Until, timeFilter, "$lt")
		filter["time"] = timeFilter
	}
	return filter
}

func timeModel(date string, filter bson.M, key string) bson.M {
	layout := "2006-01-02 15:04:05.000Z"
	modelTime, _ := time.Parse(layout, date)
	if filter == nil {
		filter = bson.M{}
	}
	filter[key] = modelTime
	return filter
}

func (dao *SongDAO) getSlugsToReset(params ResetViewsModel) []string {
	var slugs []string
	if params.Slug == "" {

		bsonQuery := []bson.M{
			{"$match": resetFilterModel("", params)},
			{"$group": bson.M{"_id": "$song"}},
		}
		results, err := dao.db.Aggregate("view_count", bsonQuery)

		if err != nil {
			fmt.Println("Failed to fetch slugs in views")
			return nil
		}

		for _, songView := range results {
			slugs = append(slugs, songView["_id"].(string))
		}

	} else {
		slugs = append(slugs, params.Slug)
	}
	return slugs
}

func (dao *SongDAO) resetSongViews(params ResetViewsModel) error {
	slugs := dao.getSlugsToReset(params)
	for _, slug := range slugs {
		deletedRows, err := dao.db.DeleteMany("view_count", resetFilterModel(slug, params))
		if err != nil {
			fmt.Println("Failed to delete song counts", slug)
			return err
		}
		err = dao.db.UpdateOne("songs", bson.M{"slug": slug}, bson.M{"$inc": bson.M{"viewed": deletedRows.DeletedCount * -1}})
		if err != nil {
			fmt.Println("Failed to update song", slug)
			return err
		}
	}
	return nil
}
