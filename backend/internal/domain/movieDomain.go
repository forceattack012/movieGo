package domain

import (
	"backend/internal/models"
	"context"
)

type MovieService interface {
	GetAllMovies(ctx context.Context) ([]models.Movie, error)
	GetMovieById(ctx context.Context, id string) (models.Movie, error)
	CreateMovie(ctx context.Context, Movie *models.Movie) error
	UpdateMovie(ctx context.Context, id string, movie *models.Movie) error
	DeleteMovie(ctx context.Context, id string) (int64, error)
}

type MovieRepository interface {
	GetAllMovies(ctx context.Context) ([]models.Movie, error)
	GetMovieById(ctx context.Context, id string) (models.Movie, error)
	CreateMovie(ctx context.Context, Movie *models.Movie) error
	UpdateMovie(ctx context.Context, id string, movie *models.Movie) error
	DeleteMovie(ctx context.Context, id string) (int64, error)
}
