package models

import "time"

type Movie struct {
	ID        int64     `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Actors    []string  `json:"actors" validate:"required"`
	Directors []string  `json:"directors" validate:"required"`
	IMDB      string    `json:"IMDB"`
	Youtube   string    `json:"youtube"`
	Duration  int64     `json:"duration" validate:"required"`
	StartDate time.Time `json:"startDate" validate:"required"`
	Synopsis  string    `json:"synopsis"`
}
