package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	echo.GET("/", hello)

	echo.Logger.Fatal(echo.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}
