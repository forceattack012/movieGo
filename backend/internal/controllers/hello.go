package controllers

import (
	"net/http"

	"backend/internal/services"

	"github.com/labstack/echo/v4"
)

func SayHello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello World")
}

func SayHelloName(ctx echo.Context) error {
	var name = ctx.Param("name")
	var hello = services.Hello(name)

	return ctx.String(http.StatusOK, hello)
}
