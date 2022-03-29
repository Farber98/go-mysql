package controllers

import (
	"go-mysql/internal/db"
	"go-mysql/internal/helpers"
	"go-mysql/internal/interfaces"
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

	result, err := manager.EstadoCuadros(id)
	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error:     nil,
		Respuesta: result,
	}
	return c.JSON(http.StatusOK, response)
}
