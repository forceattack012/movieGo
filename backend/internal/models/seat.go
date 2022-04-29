package models

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	Id         int       `json:"id" autoIncrement:"true"`
	SeatNumber int       `json:"seatNumber" validate:"required"`
	TheaterId  int       `json:"theaterId" validate:"required"`
	Price      int64     `json:"price" validate:"required"`
	Booking    []Booking `gorm:"many2many:user_reversion;"`
	Theater    Theater
}
