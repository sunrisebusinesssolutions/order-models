package models

type OrderPlacementStatus string

const (
	OrderPlacementStatusOff       OrderPlacementStatus = "OFF"
	OrderPlacementStatusScheduled OrderPlacementStatus = "SCHEDULED"
	OrderPlacementStatusOpen      OrderPlacementStatus = "OPEN"
	OrderPlacementStatusClosed    OrderPlacementStatus = "CLOSED"
)
