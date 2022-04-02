package routes

import (
	"backend/internal/controllers"
	"backend/internal/domain"

	"github.com/labstack/echo/v4"
)

func NewMovieRouter(e *echo.Echo, movieService domain.MovieService) {
	controllerMovie := controllers.NewMovieController(e, movieService)

	r := e.Group("api/v1/movie")
	r.GET("/list", controllerMovie.GetAllMovies)
	r.GET("/:id", controllerMovie.GetMovieById)
	r.POST("/", controllerMovie.CreateMovie)
}
