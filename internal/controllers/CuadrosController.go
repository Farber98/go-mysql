package controllers

import (
	"fmt"
	"go-mysql/internal/db"
	"go-mysql/internal/helpers"
	"go-mysql/internal/interfaces"
	"go-mysql/internal/managers"
	"go-mysql/internal/models"
	"go-mysql/internal/structs"
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

func (controller *CuadrosController) DarBaja(c echo.Context) error {

	cuadro := structs.Cuadros{}

	if err := c.Bind(&cuadro); err != nil {
		return interfaces.GenerarRespuestaError(err, http.StatusUnprocessableEntity)
	}

	cuadrosModel := models.CuadrosModel{
		Db:      controller.Db,
		Cuadros: &cuadro,
	}

	result, err := cuadrosModel.DarBaja()
	fmt.Println(result, err)

	if err != nil || result == nil {
		return interfaces.GenerarRespuestaError(err, http.StatusBadRequest)
	}

	response := interfaces.Response{
		Error: nil,
	}

	response.AddModels(result)

	return c.JSON(http.StatusOK, response)
}
