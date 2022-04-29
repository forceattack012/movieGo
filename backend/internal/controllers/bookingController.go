package controllers

import (
	"backend/internal/domain"
	"backend/internal/models"
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
	booking := new(models.Booking)

	if err := ec.Bind(&booking); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ec.Validate(booking); err != nil {
		return err
	}

	err := t.service.CreateBooking(booking)

	if err != nil {
		return err
	}

	return ec.JSON(http.StatusCreated, "create theater sucessful")
}
