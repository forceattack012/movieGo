package models
package models

type Movie struct {
	ID        int64      `json:"Id"`
	Name      string   `json:"Name"`
	Actors    []string `json:"Actors"`
	Directors []string `json:"Directors"`
	IMDB      string   `json:"IMDB"`
	Youtube   string   `json:"Youtube"`
	Duration  int64    `json:"Duration"`
}

type MovieService interface {
	GetAllMovies() ([]movie, error)
	GetMovieById(id int64) (movie, error)
	CreateMovie(Movie *movie) error
	UpdateMovie(id int64, movie *movie) error
	DeleteMovie(id int64) error
}

type MovieRepository struct {
	GetAllMovies() ([]movie, error)
	GetMovieById(id int64) (movie, error)
	CreateMovie(Movie *movie) error
	UpdateMovie(id int64, movie *movie) error
	DeleteMovie(id int64 int) error
}