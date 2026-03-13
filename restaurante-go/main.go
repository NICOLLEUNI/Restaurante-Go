package main

import (
	"restaurante-go/Services"
)

func main() {

	pedidos := services.IngresarPedidos()

	services.ProcesarPedidos(pedidos)

}
