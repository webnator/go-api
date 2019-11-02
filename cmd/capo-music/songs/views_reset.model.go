package songs

type ResetViewsModel struct {
	Slug  string `json:"slug" bson:"slug,omitempty"`
	Since string `json:"since" bson:"since,omitempty"`
	Until string `json:"until" bson:"until,omitempty"`
}
