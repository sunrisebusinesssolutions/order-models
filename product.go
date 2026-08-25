package models

import "time"

type Product struct {
	ID               string      `json:"id" bson:"_id"`
	Order            int         `json:"order" bson:"order"`
	Description      string      `json:"description" bson:"description"`
	Price            string      `json:"price" bson:"price"`
	SpecialPrice     string      `json:"specialPrice" bson:"specialPrice"`
	SpecialPriceDate *time.Time  `json:"specialPriceDate" bson:"specialPriceDate"`
	ItemType         ItemType    `json:"itemType" bson:"itemType"`
	Unit             string      `json:"unit" bson:"unit"`
	URL              string      `json:"url" bson:"url"`
	AvailableDays    string      `json:"availableDays" bson:"availableDays"`
	ImageOriginal    string      `json:"imageOriginal" bson:"imageOriginal"`
	ImageMedium      string      `json:"imageMedium" bson:"imageMedium"`
	ImageSmall       string      `json:"imageSmall" bson:"imageSmall"`
	ImageIcon        string      `json:"imageIcon" bson:"imageIcon"`
	ImageThumbnail   string      `json:"imageThumbnail" bson:"imageThumbnail"`
}

func (p *Product) GetImageOriginal() string {
	return p.ImageOriginal
}

func (p *Product) SetImageOriginal(url string) {
	p.ImageOriginal = url
}

func (p *Product) GetImageMedium() string {
	return p.ImageMedium
}

func (p *Product) SetImageMedium(url string) {
	p.ImageMedium = url
}

func (p *Product) GetImageSmall() string {
	return p.ImageSmall
}

func (p *Product) SetImageSmall(url string) {
	p.ImageSmall = url
}

func (p *Product) GetImageIcon() string {
	return p.ImageIcon
}

func (p *Product) SetImageIcon(url string) {
	p.ImageIcon = url
}

func (p *Product) GetImageThumbnail() string {
	return p.ImageThumbnail
}

func (p *Product) SetImageThumbnail(url string) {
	p.ImageThumbnail = url
}
