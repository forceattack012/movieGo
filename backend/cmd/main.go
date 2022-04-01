package main

import (
	"backend/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.HelloRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
