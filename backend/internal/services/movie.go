package services

import (
	"backend/internal/domain"
	"backend/internal/models"
	"context"
)

type movieService struct {
	repo domain.MovieRepository
}

func NewMovieService(repo domain.MovieRepository) domain.MovieService {
	return &movieService{
		repo: repo,
	}
}

// GetAllMovies
func (m *movieService) GetAllMovies(ctx context.Context) ([]models.Movie, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	list, err := m.repo.GetAllMovies(ctx)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// CreateMovie implements domain.MovieService
func (m *movieService) CreateMovie(ctx context.Context, Movie *models.Movie) error {
	err := m.repo.CreateMovie(ctx, Movie)

	if err != nil {
		return err
	}
	return nil
}

// DeleteMovie implements domain.MovieService
func (*movieService) DeleteMovie(id int64) error {
	panic("unimplemented")
}

// GetMovieById implements domain.MovieService
func (m *movieService) GetMovieById(ctx context.Context, id int64) (models.Movie, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := m.repo.GetMovieById(ctx, id)

	if err != nil {
		return models.Movie{}, err
	}

	return result, nil

}

// UpdateMovie implements domain.MovieService
func (*movieService) UpdateMovie(id int64, movie *models.Movie) error {
	panic("unimplemented")
}
