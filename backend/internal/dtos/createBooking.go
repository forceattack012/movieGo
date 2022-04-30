package dtos

import "time"

type CreateBooking struct {
	Id         int32        `json:"id" autoIncrement:"true"`
	ShowTimeId int64        `json:"showtimeId"`
	Timestamp  time.Time    `json:"timestamp"`
	Seat       []CreateSeat `json:"seat"`
	Total      float64      `json:"total"`
}

type CreateSeat struct {
	SeatNumber int   `json:"seatNumber"`
	TheaterId  int   `json:"theaterId"`
	Price      int64 `json:"price"`
}
