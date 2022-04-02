package domain

import (
	"backend/internal/models"
	"context"
)

type MovieService interface {
	GetAllMovies(ctx context.Context) ([]models.Movie, error)
	GetMovieById(ctx context.Context, id int64) (models.Movie, error)
	CreateMovie(ctx context.Context, Movie *models.Movie) error
	UpdateMovie(id int64, movie *models.Movie) error
	DeleteMovie(id int64) error
}

type MovieRepository interface {
	GetAllMovies(ctx context.Context) ([]models.Movie, error)
	GetMovieById(ctx context.Context, id int64) (models.Movie, error)
	CreateMovie(ctx context.Context, Movie *models.Movie) error
	UpdateMovie(id int64, movie *models.Movie) error
	DeleteMovie(id int64) error
}
