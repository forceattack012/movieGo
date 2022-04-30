package routes

import (
	"backend/internal/controllers"
	"backend/internal/domain"

	"github.com/labstack/echo/v4"
)

func NewBookingRouter(e *echo.Echo, bookingService domain.BookingService) {
	controler := controllers.NewBookingController(bookingService)

	r := e.Group("api/v1/booking")
	r.GET("/list", controler.GetBookingAll)
	r.GET("/:showTimeId", controler.GetBookingByShowTimeId)
	r.POST("/", controler.CreateBooking)

}
