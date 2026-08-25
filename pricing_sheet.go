package models

type PricingSheet struct {
	ID          string      `json:"id" bson:"_id"`
	LocationID  string      `json:"locationId" bson:"locationId"`
	Order       int         `json:"order" bson:"order"`
	ProductType ProductType `json:"productType" bson:"productType"`
	ProductCode string      `json:"productCode" bson:"productCode"`
	Description string      `json:"description" bson:"description"`
	Unit        string      `json:"unit" bson:"unit"`
	Price       string      `json:"price" bson:"price"`
}
