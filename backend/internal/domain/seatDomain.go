package domain

import "backend/internal/models"

type SeatRepository interface {
	CreateSeat(seat models.Seat) error
	CreateSeatList(seat []models.Seat) error
	GetAllSeat() ([]models.Seat, error)
	GetAllSeatByShowTimeId(showTimeId int32) ([]models.Seat, error)
	GetAllSeatByTheaterId(theaterId int32) ([]models.Seat, error)
}

type SeatService interface {
	CreateSeat(seat models.Seat) error
	CreateSeatList(seat []models.Seat) error
	GetAllSeat() ([]models.Seat, error)
	GetAllSeatByShowTimeId(showTimeId int32) ([]models.Seat, error)
}
