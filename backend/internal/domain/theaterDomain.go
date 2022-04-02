package domain

import (
	"backend/internal/models"
	"context"
)

type TheaterRepository interface {
	GetAllTheater(ctx context.Context) ([]models.Theater, error)
	GetTheaterById(ctx context.Context, id int) (models.Theater, error)
	CreateTheater(ctx context.Context, theater *models.Theater) error
	UpdateTheater(ctx context.Context, id int, theater *models.Theater) error
	DeleteTheater(ctx context.Context, id int) error
}

type TheaterService interface {
	GetAllTheater(ctx context.Context) ([]models.Theater, error)
	GetTheaterById(ctx context.Context, id int) (models.Theater, error)
	CreateTheater(ctx context.Context, theater *models.Theater) error
	UpdateTheater(ctx context.Context, id int, theater *models.Theater) error
	DeleteTheater(ctx context.Context, id int) error
}
