package models

import "time"

type CreateShowTime struct {
	Id        int32     `json:"id" autoIncrement:"true"`
	Time      time.Time `json:"time" validate:"required"`
	MovieId   string    `json:"movieId" validate:"required"`
	TheaterId int64     `json:"theaterId" validate:"required"`
}
