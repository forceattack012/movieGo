package routes

import (
	"backend/internal/controllers"

	"github.com/labstack/echo/v4"
)

var HelloRoutes = func(e *echo.Echo) {
	e.GET("/hello", controllers.SayHello)
	e.GET("/hello/:name", controllers.SayHelloName)
}
