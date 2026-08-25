package models

type OrderType string

const (
	OrderTypePickup   OrderType = "PICKUP"
	OrderTypeDelivery OrderType = "DELIVERY"
	OrderTypeStandard OrderType = "STANDARD"
)
