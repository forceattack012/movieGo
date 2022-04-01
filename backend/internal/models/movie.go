package models

type Movie struct {
	ID        int      `json:"Id"`
	Name      string   `json:"Name"`
	Actors    []string `json:"Actors"`
	Directors []string `json:"Directors"`
	IMDB      string   `json:"IMDB"`
	Youtube   string   `json:"Youtube"`
	Duration  int32    `json:"Duration"`
}
