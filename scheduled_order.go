package models

import "time"

type ScheduledOrder struct {
	Order
	WeekDay           *int              `json:"weekDay" bson:"weekDay"`
	RecurrenceType    RecurrenceType    `json:"recurrenceType" bson:"recurrenceType"`
	MonthlyPattern    MonthlyDayPattern `json:"monthlyPattern" bson:"monthlyPattern"`
	MonthlyDayOfWeek  *time.Weekday     `json:"monthlyDayOfWeek" bson:"monthlyDayOfWeek"`
}
