package controllers

import (
	"go-mysql/internal/db"
	"go-mysql/internal/helpers"
	"go-mysql/internal/managers"
	"net/http"

	"github.com/labstack/echo"
)

type CuadrosController struct {
	Db *db.MySQLHandler
}

func (controller *CuadrosController) EstadoCuadros(c echo.Context) error {
	id, _ := helpers.URLParamInt(c, "id")

	manager := managers.CuadrosManager{
		Db: controller.Db,
	}

	cuadros, err := manager.EstadoCuadros(id)
	if err != nil {
		return err
	}

	if cuadros == nil {
		return echo.NewHTTPError(http.StatusNotFound, "No existe")
	}

	return c.JSON(http.StatusOK, cuadros)
}
