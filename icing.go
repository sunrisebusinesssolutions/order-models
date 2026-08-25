package models

type Icing struct {
	ID          string `json:"id" bson:"_id"`
	Description string `json:"description" bson:"description"`
	URL         string `json:"url" bson:"url"`
}
