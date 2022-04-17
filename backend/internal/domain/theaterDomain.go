package domain

import (
	"backend/internal/models"
)

type TheaterRepository interface {
	GetAllTheater() ([]models.Theater, error)
	GetTheaterById(id int) (models.Theater, error)
	CreateTheater(theater *models.Theater) error
	UpdateTheater(id int, theater *models.Theater) error
	DeleteTheater(id int) error
}

type TheaterService interface {
	GetAllTheater() ([]models.Theater, error)
	GetTheaterById(id int) (models.Theater, error)
	CreateTheater(theater *models.Theater) error
	UpdateTheater(id int, theater *models.Theater) error
	DeleteTheater(id int) error
}
