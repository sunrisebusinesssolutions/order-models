package models

type Item struct {
	ID              string   `json:"id" bson:"_id"`
	OrderID         string   `json:"orderId" bson:"orderId"`
	ItemID          string   `json:"itemId" bson:"itemId"`
	ItemType        ItemType `json:"itemType" bson:"itemType"`
	Qty             int      `json:"qty" bson:"qty"`
	ProductName     string   `json:"productName" bson:"productName"`
	ProductImageURL string   `json:"productImageUrl" bson:"productImageUrl"`
	ProductUnit     string   `json:"productUnit" bson:"productUnit"`
	UnitPrice       *float64 `json:"unitPrice" bson:"unitPrice"`
	LineTotal       *float64 `json:"lineTotal" bson:"lineTotal"`
	FrostingID      string   `json:"frostingId" bson:"frostingId"`
	FrostingName    string   `json:"frostingName" bson:"frostingName"`
	IcingID         string   `json:"icingId" bson:"icingId"`
	IcingName       string   `json:"icingName" bson:"icingName"`
	ToppingID       string   `json:"toppingId" bson:"toppingId"`
	ToppingName     string   `json:"toppingName" bson:"toppingName"`
	Notes           string   `json:"notes" bson:"notes"`
	Mix             bool     `json:"mix" bson:"mix"`
	Brand           string   `json:"brand" bson:"brand"`
	Size            string   `json:"size" bson:"size"`
	Color           string   `json:"color" bson:"color"`
	Material        string   `json:"material" bson:"material"`
}
