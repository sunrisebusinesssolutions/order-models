package models

type Frosting struct {
	ID           string `json:"id" bson:"_id"`
	Description  string `json:"description" bson:"description"`
	URL          string `json:"url" bson:"url"`
	ImageOriginal string `json:"imageOriginal" bson:"imageOriginal"`
	ImageMedium  string `json:"imageMedium" bson:"imageMedium"`
	ImageSmall   string `json:"imageSmall" bson:"imageSmall"`
	ImageIcon    string `json:"imageIcon" bson:"imageIcon"`
}
