package models

import "time"

type Order struct {
	ID              string      `json:"id" bson:"_id"`
	CreatedBy       string      `json:"createdBy" bson:"createdBy"`
	CustomerID      int         `json:"customerId" bson:"customerId"`
	Name            string      `json:"name" bson:"name"`
	Email           string      `json:"email" bson:"email"`
	Phone           string      `json:"phone" bson:"phone"`
	OrderDate       *time.Time  `json:"orderDate" bson:"orderDate"`
	Type            OrderType   `json:"type" bson:"type"`
	DeliverDate     *time.Time  `json:"deliverDate" bson:"deliverDate"`
	Approved        bool        `json:"approved" bson:"approved"`
	ApprovedDate    *time.Time  `json:"approvedDate" bson:"approvedDate"`
	LocationID      string      `json:"locationId" bson:"locationId"`
	Status          OrderStatus `json:"status" bson:"status"`
	CalendarEventID string      `json:"calendarEventId" bson:"calendarEventId"`
	Items           []Item      `json:"items" bson:"items"`
	Notes           string      `json:"notes" bson:"notes"`
	CreatedAt       time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt" bson:"updatedAt"`
}
