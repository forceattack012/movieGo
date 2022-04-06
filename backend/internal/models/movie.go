package models

import "time"

type Movie struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Actors    []string  `json:"actors" validate:"required"`
	Directors []string  `json:"directors" validate:"required"`
	IMDB      string    `json:"IMDB"`
	Youtube   string    `json:"youtube"`
	Duration  int64     `json:"duration" validate:"required"`
	StartDate time.Time `json:"startDate" validate:"required"`
	Synopsis  string    `json:"synopsis"`
	Image     string    `json:"image"`
}
