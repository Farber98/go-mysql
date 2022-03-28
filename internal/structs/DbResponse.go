package structs

type ResponseMySQL struct {
	Error     *string     `json:"error,omitempty"`
	Respuesta interface{} `json:"respuesta,omitempty"`
}
