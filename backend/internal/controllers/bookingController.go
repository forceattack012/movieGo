package controllers

import (
	"backend/internal/domain"
	"backend/internal/dtos"
	"backend/internal/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookingController struct {
	service domain.BookingService
}

func NewBookingController(s domain.BookingService) *BookingController {
	return &BookingController{
		service: s,
	}
}

func (bc *BookingController) GetBookingByShowTimeId(ec echo.Context) error {
	id := ec.Param("showTimeId")
	showtimeId, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	result, err := bc.service.GetAllSeatsByShowTimeId(int32(showtimeId))

	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, result)

}

func (t *BookingController) CreateBooking(ec echo.Context) error {
	createBooking := new(dtos.CreateBooking)

	if err := ec.Bind(&createBooking); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ec.Validate(createBooking); err != nil {
		return err
	}

	seats := make([]models.Seat, 0)
	book := make([]models.Booking, 0)
	for _, s := range createBooking.Seat {
		seats = append(seats, models.Seat{
			TheaterId:  s.TheaterId,
			Price:      s.Price,
			SeatNumber: s.SeatNumber,
			Id:         s.SeatNumber,
			Booking: append(book, models.Booking{
				Id:         int32(s.SeatNumber),
				ShowTimeId: createBooking.ShowTimeId,
			}),
		})
	}

	fmt.Println(seats)
	booking := &models.Booking{
		ShowTimeId: createBooking.ShowTimeId,
		Total:      createBooking.Total,
		Timestamp:  createBooking.Timestamp,
		Seat:       seats,
	}
	err := t.service.CreateBooking(booking)

	if err != nil {
		return err
	}

	return ec.JSON(http.StatusCreated, "create theater sucessful")
}

func (bc *BookingController) GetBookingAll(ec echo.Context) error {
	result, err := bc.service.GetAllBooking()

	if err != nil {
		return err
	}
	return ec.JSON(http.StatusOK, result)

}
