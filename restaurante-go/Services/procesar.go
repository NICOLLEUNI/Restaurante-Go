package services

import (
	"fmt"

	models "restaurante-go/Models"
)

func ProcesarPedidos(listaPedidos []models.Pedido) {

	pedidos := make(chan models.Pedido)
	resultados := make(chan models.Pedido)

	numCocineros := 3

	for i := 1; i <= numCocineros; i++ {
		go Cocinero(i, pedidos, resultados)
	}

	go func() {

		for _, pedido := range listaPedidos {
			pedidos <- pedido
		}

		close(pedidos)

	}()

	for i := 0; i < len(listaPedidos); i++ {

		pedido := <-resultados

		fmt.Println("Pedido listo:", pedido.Comida)

	}

}
