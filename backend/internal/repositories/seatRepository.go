package repositories

import (
	"backend/internal/domain"
	"backend/internal/models"

	"gorm.io/gorm"
)

type SeatRepository struct {
	db *gorm.DB
}

func NewSeatRepository(db *gorm.DB) domain.SeatRepository {
	return &SeatRepository{
		db: db,
	}
}

// GetAllSeatByShowTimeId implements domain.ShowTimeRepository
func (s *SeatRepository) GetAllSeatByShowTimeId(showTimeId int32) ([]models.Seat, error) {
	seats := make([]models.Seat, 0)
	if err := s.db.Where(&models.ShowTime{Id: showTimeId}).Joins("Theater").Joins("Seat").Find(&seats).Error; err != nil {
		return nil, err
	}

	return seats, nil
}

// CreateSeat implements domain.SeatRepository
func (s *SeatRepository) CreateSeat(seat models.Seat) error {
	if err := s.db.Create(seat).Error; err != nil {
		return err
	}

	return nil
}

// CreateSeatList implements domain.SeatRepository
func (s *SeatRepository) CreateSeatList(seat []models.Seat) error {
	if err := s.db.Create(seat).Error; err != nil {
		return err
	}

	return nil
}

// GetAllSeat implements domain.SeatRepository
func (*SeatRepository) GetAllSeat() ([]models.Seat, error) {
	panic("unimplemented")
}
