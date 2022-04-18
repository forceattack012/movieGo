package models

type Seat struct {
	Id         int `json:"id" autoIncrement:"true"`
	SeatNumber int `json:"seatNumber" validate:"required"`
	TheaterId  int `json:"theaterId" validate:"required"`
	Theater    Theater
}
