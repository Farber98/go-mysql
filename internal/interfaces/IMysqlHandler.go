package interfaces

// IDBHandler interface para la conexión a la base de datos
type IDBHandler interface {
	CallSP(sp string, objeto interface{}) (*[]byte, error)
}
