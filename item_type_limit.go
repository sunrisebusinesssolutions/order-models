package models

type ItemTypeLimit struct {
	ID       string  `json:"id" bson:"_id"`
	ItemType ItemType `json:"itemType" bson:"itemType"`
	Capacity int     `json:"capacity" bson:"capacity"`
	SheetSize *int   `json:"sheetSize" bson:"sheetSize"`
}
