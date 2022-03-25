package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Prueba(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Mensaje string }{"Metodo de prueba"})
}
