package models

import "time"

type ContactMessage struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Subject   string    `json:"subject" bson:"subject"`
	Message   string    `json:"message" bson:"message"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
