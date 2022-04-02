package controllers

import (
	"backend/internal/domain"
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

	return ec.JSON(http.StatusOK, "OK Theater")
	// //list, err := t.TheaterService.GetAllTheater()()

	// if err != nil {
	// 	return err
	// }

	// ec.JSON(http.StatusOK, list)
}
