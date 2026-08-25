package models

import "time"

type Announcement struct {
	ID          string     `json:"id" bson:"_id"`
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	StartDate   *time.Time `json:"startDate" bson:"startDate"`
	EndDate     *time.Time `json:"endDate" bson:"endDate"`
	Color       string     `json:"color" bson:"color"`
	Active      bool       `json:"active" bson:"active"`
	CreatedAt   time.Time  `json:"createdAt" bson:"createdAt"`
	CreatedBy   string     `json:"createdBy" bson:"createdBy"`
}

func (a *Announcement) GetEffectiveEndDate() *time.Time {
	if a.EndDate != nil {
		return a.EndDate
	}
	return a.StartDate
}

func (a *Announcement) IsMultiDay() bool {
	if a.EndDate == nil || a.StartDate == nil {
		return false
	}
	return !a.EndDate.Equal(*a.StartDate)
}
