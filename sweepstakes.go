package models

import "time"

type Sweepstakes struct {
	ID              string     `json:"id" bson:"_id"`
	Title           string     `json:"title" bson:"title"`
	Description     string     `json:"description" bson:"description"`
	BeginDate       *time.Time `json:"beginDate" bson:"beginDate"`
	EndDate         *time.Time `json:"endDate" bson:"endDate"`
	ImageURL        string     `json:"imageUrl" bson:"imageUrl"`
	WinnerEmail     string     `json:"winnerEmail" bson:"winnerEmail"`
	WinnerDate      *time.Time `json:"winnerDate" bson:"winnerDate"`
	TicketNumber    string     `json:"ticketNumber" bson:"ticketNumber"`
	EnteredUsers    []string   `json:"enteredUsers" bson:"enteredUsers"`
	CalendarEventID string     `json:"calendarEventId" bson:"calendarEventId"`
	CreatedAt       time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt" bson:"updatedAt"`
}
