package dtos

import (
	"backend/internal/models"
	"time"
)

type ResponseShowTime struct {
	Id        int32   `json:"id" autoIncrement:"true"`
	Times     []Times `json:"times" validate:"required"`
	MovieId   string  `json:"movieId" validate:"required"`
	TheaterId int64   `json:"theaterId" validate:"required"`
	Theater   models.Theater
}

type Times struct {
	Id   int32     `json:"id" autoIncrement:"true"`
	Time time.Time `json:"times" validate:"required"`
}

type TimeSlice []Times

// Forward request for length

func (p TimeSlice) Len() int {

	return len(p)
}

// Define compare

func (p TimeSlice) Less(i, j int) bool {

	return p[i].Time.Before(p[j].Time)
}

// Define swap over an array

func (p TimeSlice) Swap(i, j int) {

	p[i], p[j] = p[j], p[i]
}
