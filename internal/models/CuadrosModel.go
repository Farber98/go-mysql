package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-mysql/internal/db"
	"go-mysql/internal/structs"

	"github.com/mitchellh/mapstructure"
)

type CuadrosModel struct {
	Db      *db.MySQLHandler
	Cuadros *structs.Cuadros
}

func (cm *CuadrosModel) DarBaja() (*structs.Cuadros, error) {
	objeto := struct {
		IDCuadro int `json:"IDCuadro"`
	}{IDCuadro: cm.Cuadros.IDCuadro}
	out, err := cm.Db.CallSP("sp_cuadros_baja", objeto)
	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, errors.New("ERROR_DEFAULT")
	}

	var response map[string]interface{}

	err = json.Unmarshal(*out, &response)
	if err != nil {
		return nil, err
	}
	var cuadro structs.Cuadros
	if response["Cuadro"] != nil {
		err = mapstructure.Decode(response["Cuadro"], &cuadro)
		fmt.Println(err)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &cuadro, nil
}
