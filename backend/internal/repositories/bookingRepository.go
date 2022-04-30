package repositories

import (
	"backend/internal/domain"
	"backend/internal/models"

	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) domain.BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

// CreateBooking implements domain.BookingRepository
func (br *BookingRepository) CreateBooking(booking *models.Booking) error {
	err := br.db.Create(booking).Error

	if err != nil {
		return err
	}

	return nil
}

// GetAllSeatsByShowTimeId implements domain.BookingRepository
func (br *BookingRepository) GetAllSeatsByShowTimeId(showTime int64) (*models.Booking, error) {
	booking := models.Booking{}
	if err := br.db.Where(models.ShowTime{Id: int32(showTime)}).Preload("Seat").Find(&booking).Error; err != nil {
		return nil, err
	}

	return &booking, nil
}

func (br *BookingRepository) GetAllBooking() ([]models.Booking, error) {
	bookings := make([]models.Booking, 0)
	if err := br.db.Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}
