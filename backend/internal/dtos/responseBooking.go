package dtos

import (
	"backend/internal/models"
	"time"
)

type ReponseBooking struct {
	MovieName    string `json:"movieName"`
	Duration     int    `json:"duration"`
	TheaterName  string `json:"theaterName"`
	ShowTimeId   int64  `json:"showTimeId"`
	Image        string `json:"image"`
	Time         time.Time
	BookingSeats []models.Seat `json:"bookingSeats"`
	Seats        []models.Seat `json:"seats"`
}
