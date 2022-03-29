package interfaces

import (
	"go-mysql/internal/helpers"
	"go-mysql/internal/structs"

	"github.com/labstack/echo"
)

type Response struct {
	Error     *structs.Errores `json:"error"`
	Respuesta interface{}      `json:"respuesta"`
}

func GenerarRespuestaError(err error, code int) *echo.HTTPError {
	codigo, mensaje := helpers.GetError(err)
	objError := structs.Errores{
		Codigo:  codigo,
		Mensaje: mensaje,
	}
	response := Response{
		Error:     &objError,
		Respuesta: nil,
	}
	return echo.NewHTTPError(code, response)
}
