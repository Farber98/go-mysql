package structs

type CuadrosEstado []struct {
	Cuadro    string  `json:"Cuadro"`
	Precio    float64 `json:"Precio"`
	Vendido   string  `json:"Vendido"`
	Ganancia  float64 `json:"Ganancia"`
	IDCuadro  int     `json:"IDCuadro"`
	Propuesta float64 `json:"Propuesta"`
}

type Cuadros struct {
	IDCuadro int
	IDPintor int
	IDMetodo int
	Titulo   string
	Fecha    string
	Precio   float64
	Estado   string
}
