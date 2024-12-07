package models

import "time"

type Booking struct {
	Id        int       `json:"id"`
	PetID     int       `json:"pet_id"`
	UserID    uint      `json:"userId"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
