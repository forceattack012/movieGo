package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	Id         int32 `json:"id" autoIncrement:"true"`
	ShowTimeId int64 `json:"showtimeId" validate:"required"`
	ShowTime   ShowTime
	Timestamp  time.Time `json:"timestamp" validate:"required"`
	Seat       []Seat    `json:"seat" gorm:"many2many:user_reversion;"`
	Total      float64   `json:"total" validate:"required"`
}
