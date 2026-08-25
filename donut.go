package models

type Donut struct {
	Product
	Frosting         bool `json:"frosting" bson:"frosting"`
	RequiredFrosting bool `json:"requiredFrosting" bson:"requiredFrosting"`
	Icing            bool `json:"icing" bson:"icing"`
	RequiredIcing    bool `json:"requiredIcing" bson:"requiredIcing"`
	Topping          bool `json:"topping" bson:"topping"`
	RequiredTopping  bool `json:"requiredTopping" bson:"requiredTopping"`
	AllowMix         bool `json:"allowMix" bson:"allowMix"`
}

func (d *Donut) IsFrosting() bool {
	return d.Frosting
}

func (d *Donut) SetFrosting(frosting bool) {
	d.Frosting = frosting
}

func (d *Donut) IsRequiredFrosting() bool {
	return d.RequiredFrosting
}

func (d *Donut) SetRequiredFrosting(requiredFrosting bool) {
	d.RequiredFrosting = requiredFrosting
}

func (d *Donut) IsIcing() bool {
	return d.Icing
}

func (d *Donut) SetIcing(icing bool) {
	d.Icing = icing
}

func (d *Donut) IsRequiredIcing() bool {
	return d.RequiredIcing
}

func (d *Donut) SetRequiredIcing(requiredIcing bool) {
	d.RequiredIcing = requiredIcing
}

func (d *Donut) IsTopping() bool {
	return d.Topping
}

func (d *Donut) SetTopping(topping bool) {
	d.Topping = topping
}

func (d *Donut) IsRequiredTopping() bool {
	return d.RequiredTopping
}

func (d *Donut) SetRequiredTopping(requiredTopping bool) {
	d.RequiredTopping = requiredTopping
}

func (d *Donut) IsAllowMix() bool {
	return d.AllowMix
}

func (d *Donut) SetAllowMix(allowMix bool) {
	d.AllowMix = allowMix
}
