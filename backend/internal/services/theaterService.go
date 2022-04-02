package services

import (
	"backend/internal/domain"
	"backend/internal/models"
	"context"
)

type TheaterService struct {
	TheaterRepo domain.TheaterRepository
}

func NewTheaterService(repo domain.TheaterRepository) domain.TheaterService {
	return &TheaterService{
		TheaterRepo: repo,
	}
}

// CreateTheater implements domain.TheaterRepository
func (*TheaterService) CreateTheater(ctx context.Context, theater *models.Theater) error {
	panic("unimplemented")
}

// DeleteTheater implements domain.TheaterRepository
func (*TheaterService) DeleteTheater(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetAllTheater implements domain.TheaterRepository
func (*TheaterService) GetAllTheater(ctx context.Context) ([]models.Theater, error) {
	panic("unimplemented")
}

// GetTheaterById implements domain.TheaterRepository
func (*TheaterService) GetTheaterById(ctx context.Context, id int) (models.Theater, error) {
	panic("unimplemented")
}

// UpdateTheater implements domain.TheaterRepository
func (*TheaterService) UpdateTheater(ctx context.Context, id int, theater *models.Theater) error {
	panic("unimplemented")
}
