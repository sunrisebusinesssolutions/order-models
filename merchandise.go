package models

type Merchandise struct {
	Product
	Size            string `json:"size" bson:"size"`
	Color           string `json:"color" bson:"color"`
	Material        string `json:"material" bson:"material"`
	Brand           string `json:"brand" bson:"brand"`
	HasSizeOptions  bool   `json:"hasSizeOptions" bson:"hasSizeOptions"`
	HasColorOptions bool   `json:"hasColorOptions" bson:"hasColorOptions"`
}
