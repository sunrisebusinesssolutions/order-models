package models

type Location struct {
	ID              string   `json:"id" bson:"_id"`
	Name            string   `json:"name" bson:"name"`
	Address1        string   `json:"address1" bson:"address1"`
	Address2        string   `json:"address2" bson:"address2"`
	City            string   `json:"city" bson:"city"`
	State           string   `json:"state" bson:"state"`
	Zip             string   `json:"zip" bson:"zip"`
	Phone           string   `json:"phone" bson:"phone"`
	Email           string   `json:"email" bson:"email"`
	PrimaryLocation bool     `json:"primaryLocation" bson:"primaryLocation"`
	Visible         bool     `json:"visible" bson:"visible"`
	Latitude        *float64 `json:"latitude" bson:"latitude"`
	Longitude       *float64 `json:"longitude" bson:"longitude"`
}

func (l *Location) HasGeolocation() bool {
	return l.Latitude != nil && l.Longitude != nil
}
