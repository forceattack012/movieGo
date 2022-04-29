package services

import (
	"backend/internal/domain"
	"backend/internal/dtos"
	"backend/internal/models"
	"context"
	"fmt"
	"time"
)

type BookingService struct {
	repo         domain.BookingRepository
	showTimeRepo domain.ShowTimeRepository
	seatRepo     domain.SeatRepository
	movieRepo    domain.MovieRepository
}

func NewBookingService(repo domain.BookingRepository, sr domain.ShowTimeRepository, seat domain.SeatRepository, movieRepo domain.MovieRepository) domain.BookingService {
	return &BookingService{
		repo:         repo,
		showTimeRepo: sr,
		seatRepo:     seat,
		movieRepo:    movieRepo,
	}
}

// CreateBooking implements domain.BookingService
func (bs *BookingService) CreateBooking(m *models.Booking) error {
	err := bs.repo.CreateBooking(m)

	if err != nil {
		return err
	}

	return nil
}

// GetAllSeatsByShowTimeId implements domain.BookingService
func (s *BookingService) GetAllSeatsByShowTimeId(showTime int32) (*dtos.ReponseBooking, error) {

	result, err := s.showTimeRepo.GetAllSeatByShowTimeId(showTime)
	if err != nil {
		return nil, err
	}

	seats, _ := s.seatRepo.GetAllSeatByTheaterId(int32(result.TheaterId))
	seatsResponse := make([]models.Seat, 0)
	for _, seat := range seats {
		seatsResponse = append(seatsResponse, models.Seat{
			Id:         seat.Id,
			SeatNumber: seat.SeatNumber,
			Price:      seat.Price,
			TheaterId:  seat.TheaterId,
		})
	}

	seatsBooking := make([]models.Seat, 0)
	if result.Booking != nil {
		fmt.Print(seatsBooking, result.Booking)
	}

	movie, _ := s.movieRepo.GetMovieById(context.Background(), result.MovieId)
	res := &dtos.ReponseBooking{
		MovieName:    movie.Name,
		Image:        movie.Image,
		Duration:     int(movie.Duration),
		TheaterName:  result.Theater.Name,
		ShowTimeId:   int64(result.Id),
		Time:         time.Time{},
		BookingSeats: seatsBooking,
		Seats:        seatsResponse,
	}

	return res, nil

}
