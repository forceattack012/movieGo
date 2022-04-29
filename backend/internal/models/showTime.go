package models

import (
	"time"

	"gorm.io/gorm"
)

type ShowTime struct {
	gorm.Model
	Id        int32     `json:"id" autoIncrement:"true"`
	Time      time.Time `json:"time" validate:"required"`
	MovieId   string    `json:"movieId" validate:"required"`
	TheaterId int64     `json:"theaterId" validate:"required"`
	Booking   []Booking `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL, ForeignKey:ShowTimeId;references:Id"`
	Theater   Theater
}
