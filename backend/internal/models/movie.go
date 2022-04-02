package models

type Movie struct {
	ID        int64    `json:"Id" validate:"required"`
	Name      string   `json:"Name" validate:"required"`
	Actors    []string `json:"Actors" validate:"required"`
	Directors []string `json:"Directors" validate:"required"`
	IMDB      string   `json:"IMDB"`
	Youtube   string   `json:"Youtube"`
	Duration  int64    `json:"Duration" validate:"required"`
}
