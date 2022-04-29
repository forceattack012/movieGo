package domain

import (
	"backend/internal/dtos"
	"backend/internal/models"
)

type BookingService interface {
	GetAllSeatsByShowTimeId(showTime int32) (*dtos.ReponseBooking, error)
	CreateBooking(*models.Booking) error
}

type BookingRepository interface {
	GetAllSeatsByShowTimeId(showTime int64) (*models.Booking, error)
	CreateBooking(*models.Booking) error
}
