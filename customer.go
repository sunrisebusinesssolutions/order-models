package models

type Customer struct {
	ID         string `json:"id" bson:"_id"`
	FirstName  string `json:"firstName" bson:"firstName"`
	MiddleName string `json:"middleName" bson:"middleName"`
	LastName   string `json:"lastName" bson:"lastName"`
	Address1   string `json:"address1" bson:"address1"`
	Address2   string `json:"address2" bson:"address2"`
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	Zip        string `json:"zip" bson:"zip"`
	Phone      string `json:"phone" bson:"phone"`
	Email      string `json:"email" bson:"email"`
}
