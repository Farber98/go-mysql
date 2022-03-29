package structs

type CuadrosEstado []struct {
	Cuadro    string  `json:"Cuadro"`
	Precio    float64 `json:"Precio"`
	Vendido   string  `json:"Vendido"`
	Ganancia  float64 `json:"Ganancia"`
	IDCuadro  int     `json:"IDCuadro"`
	Propuesta float64 `json:"Propuesta"`
}
