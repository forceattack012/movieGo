package main

import (
	"backend/internal/database/mongo"
	"backend/internal/repositories"
	"backend/internal/routes"
	"backend/internal/services"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	config := mongo.BuildConfig()
	mongoDb, err := mongo.Connect(config)

	mongo.DeleteSeed(mongoDb, "Movies")
	mongo.Seed(mongoDb, "Movies")

	if err != nil {
		log.Println(err)
		return
	}

	movieRepo := repositories.NewMovieRepository(mongoDb, "Movies")
	movieService := services.NewMovieService(movieRepo)

	e := echo.New()
	e.Validator = &ValidateRequest{
		v: validator.New(),
	}

	routes.HelloRoutes(e)
	routes.NewMovieRouter(e, movieService)

	e.Logger.Fatal(e.Start(":3000"))
}

type ValidateRequest struct {
	v *validator.Validate
}

func (cv *ValidateRequest) Validate(i interface{}) error {
	if err := cv.v.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
