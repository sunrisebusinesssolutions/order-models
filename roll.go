package models

type Roll struct {
	Product
	Frosting         bool `json:"frosting" bson:"frosting"`
	RequiredFrosting bool `json:"requiredFrosting" bson:"requiredFrosting"`
	Icing            bool `json:"icing" bson:"icing"`
	RequiredIcing    bool `json:"requiredIcing" bson:"requiredIcing"`
	Topping          bool `json:"topping" bson:"topping"`
	RequiredTopping  bool `json:"requiredTopping" bson:"requiredTopping"`
	AllowMix         bool `json:"allowMix" bson:"allowMix"`
}

func (r *Roll) IsFrosting() bool {
	return r.Frosting
}

func (r *Roll) SetFrosting(frosting bool) {
	r.Frosting = frosting
}

func (r *Roll) IsRequiredFrosting() bool {
	return r.RequiredFrosting
}

func (r *Roll) SetRequiredFrosting(requiredFrosting bool) {
	r.RequiredFrosting = requiredFrosting
}

func (r *Roll) IsIcing() bool {
	return r.Icing
}

func (r *Roll) SetIcing(icing bool) {
	r.Icing = icing
}

func (r *Roll) IsRequiredIcing() bool {
	return r.RequiredIcing
}

func (r *Roll) SetRequiredIcing(requiredIcing bool) {
	r.RequiredIcing = requiredIcing
}

func (r *Roll) IsTopping() bool {
	return r.Topping
}

func (r *Roll) SetTopping(topping bool) {
	r.Topping = topping
}

func (r *Roll) IsRequiredTopping() bool {
	return r.RequiredTopping
}

func (r *Roll) SetRequiredTopping(requiredTopping bool) {
	r.RequiredTopping = requiredTopping
}

func (r *Roll) IsAllowMix() bool {
	return r.AllowMix
}

func (r *Roll) SetAllowMix(allowMix bool) {
	r.AllowMix = allowMix
}
