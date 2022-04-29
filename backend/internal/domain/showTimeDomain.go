package domain

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"time"
)

type ShowTimeService interface {
	CreateShowTime(showtime *models.ShowTime) error
	GetAllShowTime() ([]models.ShowTime, error)
	DeleteShowTime(id int64) error
	GetShowTimeSlot(theaterId int64) []time.Time
	GetShowTimesByMovieId(movieId string) ([]dtos.ResponseShowTime, error)
}

type ShowTimeRepository interface {
	CreateShowTime(showtime *models.ShowTime) error
	GetAllShowTimeById(id int64) (*models.ShowTime, error)
	GetAllShowTime() ([]models.ShowTime, error)
	DeleteShowTime(id int64, showtime *models.ShowTime) error
	GetShowTimeSlot(theaterId int64) ([]time.Time, error)
	IsShowTimeByTheaterAndTime(theaterId int64, time time.Time) (bool, error)
	GetShowTimesByMovieId(movieId string) ([]models.ShowTime, error)
	GetAllSeatByShowTimeId(showTimeId int32) (*models.ShowTime, error)
}
