package models

type Pedido struct {
	ID             int
	Comida         string
	Especificacion string
	Estado         string
	Tiempo         int
	CocineroID     int
}
