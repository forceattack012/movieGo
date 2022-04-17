package controllers

import (
	"backend/internal/domain"
	"backend/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TheaterController struct {
	TheaterService domain.TheaterService
}

func NewTheaterController(service domain.TheaterService) *TheaterController {
	return &TheaterController{
		TheaterService: service,
	}
}

func (t *TheaterController) GetAllTheater(ec echo.Context) error {

	list, err := t.TheaterService.GetAllTheater()

	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, list)
}

func (t *TheaterController) CreateTheater(ec echo.Context) error {
	theater := new(models.Theater)

	if err := ec.Bind(&theater); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ec.Validate(theater); err != nil {
		return err
	}

	err := t.TheaterService.CreateTheater(theater)

	if err != nil {
		return err
	}

	return ec.JSON(http.StatusCreated, "create theater sucessful")
}
