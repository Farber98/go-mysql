package managers

import (
	"encoding/json"
	"go-mysql/internal/db"
	"go-mysql/internal/structs"
)

type CuadrosManager struct {
	Db *db.MySQLHandler
}

func (manager *CuadrosManager) EstadoCuadros(IDPintor int) (*structs.CuadrosEstado, error) {
	objeto := struct {
		IDPintor int `json:"IDPintor"`
	}{IDPintor: IDPintor}

	out, err := manager.Db.CallSP("sp_cuadros_estado", objeto)
	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, nil
	}
	cuadros := structs.CuadrosEstado{}
	if err := json.Unmarshal(*out, &cuadros); err != nil {
		return nil, err
	}

	return &cuadros, nil
}
