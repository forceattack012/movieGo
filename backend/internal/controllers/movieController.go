package controllers

import (
	"backend/internal/domain"
	"backend/internal/models"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type movieController struct {
	echo    *echo.Echo
	service domain.MovieService
}

func NewMovieController(eh *echo.Echo, s domain.MovieService) *movieController {
	return &movieController{
		echo:    eh,
		service: s,
	}
}

// GetAllMovies implements domain.MovieService
func (mc *movieController) GetAllMovies(ec echo.Context) error {
	list, err := mc.service.GetAllMovies(context.Background())

	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, list)
}

func (mc *movieController) GetMovieById(ec echo.Context) error {
	id, err := strconv.ParseInt(ec.Param("id"), 0, 64)

	if err != nil {
		return err
	}

	movie, err := mc.service.GetMovieById(context.Background(), id)

	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, movie)
}

func (mc *movieController) CreateMovie(ec echo.Context) error {

	movie := new(models.Movie)

	if err := ec.Bind(movie); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ec.Validate(movie); err != nil {
		return err
	}

	err := mc.service.CreateMovie(context.Background(), movie)

	if err != nil {
		return err
	}

	return ec.String(http.StatusCreated, "Create Successful")
}
