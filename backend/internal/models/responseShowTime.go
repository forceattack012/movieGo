package models

import "time"

type ResponseShowTime struct {
	Id        int32       `json:"id" autoIncrement:"true"`
	Times     []time.Time `json:"times" validate:"required"`
	MovieId   string      `json:"movieId" validate:"required"`
	TheaterId int64       `json:"theaterId" validate:"required"`
	Theater   Theater
}
