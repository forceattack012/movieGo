package routes

import (
	"backend/internal/controllers"
	"backend/internal/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewMovieRouter(e *echo.Echo, movieService domain.MovieService) {
	controllerMovie := controllers.NewMovieController(e, movieService)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	r := e.Group("api/v1/movie")
	r.GET("/list", controllerMovie.GetAllMovies)
	r.GET("/:id", controllerMovie.GetMovieById)
	r.POST("/", controllerMovie.CreateMovie)
}
