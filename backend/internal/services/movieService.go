package services

import (
	"backend/internal/domain"
	"backend/internal/models"
	"context"
	"strings"

	"github.com/google/uuid"
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
	uuId := uuid.New()
	Movie.ID = strings.Replace(uuId.String(), "-", "", -1)

	err := m.repo.CreateMovie(ctx, Movie)

	if err != nil {
		return err
	}
	return nil
}

// DeleteMovie implements domain.MovieService
func (m *movieService) DeleteMovie(ctx context.Context, id string) (int64, error) {
	result, err := m.repo.DeleteMovie(ctx, id)

	if err != nil {
		return 0, err
	}

	return result, err
}

// GetMovieById implements domain.MovieService
func (m *movieService) GetMovieById(ctx context.Context, id string) (models.Movie, error) {
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
func (m *movieService) UpdateMovie(ctx context.Context, id string, movie *models.Movie) (int64, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := m.repo.UpdateMovie(ctx, id, movie)

	if err != nil {
		return 0, err
	}

	return result, nil
}
