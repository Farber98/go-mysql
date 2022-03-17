package interfaces

type IDBHandler interface {
	CallSP(sp string, objeto interface{}) (*[]byte, error)
}
