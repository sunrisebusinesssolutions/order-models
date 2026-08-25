package models

import "time"

type Status struct {
	ID               string               `json:"id" bson:"_id"`
	Status           OrderPlacementStatus `json:"status" bson:"status"`
	StartTime        *time.Time           `json:"startTime" bson:"startTime"`
	EndTime          *time.Time           `json:"endTime" bson:"endTime"`
	PendingMessage   string               `json:"pendingMessage" bson:"pendingMessage"`
	PendingSubject   string               `json:"pendingSubject" bson:"pendingSubject"`
	ApprovedMessage  string               `json:"approvedMessage" bson:"approvedMessage"`
	ApprovedSubject  string               `json:"approvedSubject" bson:"approvedSubject"`
	DeclinedMessage  string               `json:"declinedMessage" bson:"declinedMessage"`
	DeclinedSubject  string               `json:"declinedSubject" bson:"declinedSubject"`
	CompletedMessage string               `json:"completedMessage" bson:"completedMessage"`
	CompletedSubject string               `json:"completedSubject" bson:"completedSubject"`
	DeniedMessage    string               `json:"deniedMessage" bson:"deniedMessage"`
	DeniedSubject    string               `json:"deniedSubject" bson:"deniedSubject"`
	DeletedMessage   string               `json:"deletedMessage" bson:"deletedMessage"`
	DeletedSubject   string               `json:"deletedSubject" bson:"deletedSubject"`
	StaffMessage     string               `json:"staffMessage" bson:"staffMessage"`
}
