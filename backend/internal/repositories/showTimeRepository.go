package repositories

import (
	"backend/internal/domain"
	"backend/internal/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ShowTimeRepository struct {
	db *gorm.DB
}

func NewShowTimeRepository(conn *gorm.DB) domain.ShowTimeRepository {
	return &ShowTimeRepository{
		db: conn,
	}
}

// GetAllSeatByShowTimeId implements domain.ShowTimeRepository
func (s *ShowTimeRepository) GetAllSeatByShowTimeId(showTimeId int32) ([]models.Seat, error) {
	seats := make([]models.Seat, 0)
	if err := s.db.Where(&models.ShowTime{Id: showTimeId}).Joins("Theater").Joins("Seat").Find(&seats).Error; err != nil {
		return nil, err
	}

	return seats, nil
}

// GetShowTimesByMovieId implements domain.ShowTimeRepository
func (repo *ShowTimeRepository) GetShowTimesByMovieId(movieId string) ([]models.ShowTime, error) {
	showTimes := make([]models.ShowTime, 0)

	if err := repo.db.Joins("Theater").Where(&models.ShowTime{MovieId: movieId}).Find(&showTimes).Error; err != nil {
		return nil, err
	}

	return showTimes, nil
}

// IsShowTimeByTheaterAndTime implements domain.ShowTimeRepository
func (repo *ShowTimeRepository) IsShowTimeByTheaterAndTime(theaterId int64, time time.Time) (bool, error) {
	showTime := models.ShowTime{}
	repo.db.Where(&models.ShowTime{TheaterId: theaterId, Time: time}).Find(&showTime)

	if showTime.MovieId != "" {
		return true, errors.New("show time is exits not create")
	}

	return false, nil
}

// CreateShowTime implements domain.ShowTimeRepository
func (repo *ShowTimeRepository) CreateShowTime(showtime *models.ShowTime) error {
	if err := repo.db.Create(showtime).Error; err != nil {
		return err
	}

	return nil
}

// GetAllShowTimeById implements domain.ShowTimeRepository
func (repo *ShowTimeRepository) GetAllShowTimeById(id int64) (*models.ShowTime, error) {
	showtime := models.ShowTime{}

	if err := repo.db.FirstOrInit(&showtime, id).Error; err != nil {
		return nil, err
	}

	return &showtime, nil
}

// DeleteShowTime implements domain.ShowTimeRepository
func (repo *ShowTimeRepository) DeleteShowTime(id int64, showtime *models.ShowTime) error {
	if err := repo.db.Delete(showtime, id).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// GetAllShowTime implements domain.ShowTimeRepository
func (repo *ShowTimeRepository) GetAllShowTime() ([]models.ShowTime, error) {
	showTimes := make([]models.ShowTime, 0)

	if err := repo.db.Joins("Theater").Find(&showTimes).Error; err != nil {
		return nil, err
	}

	return showTimes, nil
}

// GetShowTimeSlot implements domain.ShowTimeRepository
func (repo *ShowTimeRepository) GetShowTimeSlot(theaterId int64) ([]time.Time, error) {
	panic("unimplemented")
}
