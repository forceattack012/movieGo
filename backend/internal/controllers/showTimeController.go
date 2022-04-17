package controllers

import (
	"backend/internal/domain"
	"backend/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ShowTimeController struct {
	service domain.ShowTimeService
}

func NewShowTimeController(ser domain.ShowTimeService) *ShowTimeController {
	return &ShowTimeController{
		service: ser,
	}
}

func (controller *ShowTimeController) GetShowTimes(ec echo.Context) error {
	list, err := controller.service.GetAllShowTime()

	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, list)
}

func (controller *ShowTimeController) CreateShowTime(ec echo.Context) error {
	createShowTime := new(models.CreateShowTime)

	if err := ec.Bind(&createShowTime); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ec.Validate(createShowTime); err != nil {
		return err
	}

	showTime := &models.ShowTime{
		Time:      createShowTime.Time,
		MovieId:   createShowTime.MovieId,
		TheaterId: createShowTime.TheaterId,
	}
	err := controller.service.CreateShowTime(showTime)

	if err != nil {
		return err
	}

	return ec.JSON(http.StatusCreated, "create showtime sucessful")
}

func (controller *ShowTimeController) DeleteShowTime(ec echo.Context) error {
	id := ec.Param("showtimeId")
	showtimeId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return err
	}

	if err := controller.service.DeleteShowTime(showtimeId); err != nil {
		return err
	}

	return ec.JSON(http.StatusNotFound, "delete successful")
}

func (controller *ShowTimeController) GetShowTimesByMovieId(ec echo.Context) error {
	id := ec.Param("movieId")

	list, err := controller.service.GetShowTimesByMovieId(id)

	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, list)
}
