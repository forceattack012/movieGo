package routes

import (
	"backend/internal/controllers"
	"backend/internal/domain"

	"github.com/labstack/echo/v4"
)

func NewShowTimeRoute(e *echo.Echo, service domain.ShowTimeService) {
	controler := controllers.NewShowTimeController(service)

	r := e.Group("api/v1/showtime")
	r.GET("/list", controler.GetShowTimes)
	r.GET("/:movieId", controler.GetShowTimesByMovieId)
	r.POST("/", controler.CreateShowTime)
	r.DELETE("/remove/:showtimeId", controler.DeleteShowTime)
}
