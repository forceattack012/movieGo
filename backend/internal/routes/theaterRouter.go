package routes

import (
	"backend/internal/controllers"
	"backend/internal/domain"

	"github.com/labstack/echo/v4"
)

func NewTheaterRouter(e *echo.Echo, service domain.TheaterService) {
	controller := controllers.NewTheaterController(service)

	r := e.Group("api/v1/theater")
	r.GET("/list", controller.GetAllTheater)
}
