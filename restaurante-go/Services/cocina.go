package services

import (
	"fmt"
	"time"

	"restaurante-go/Models"
)

func Cocinero(id int, pedidos <-chan models.Pedido, resultados chan<- models.Pedido) {

	for pedido := range pedidos {

		fmt.Println("Cocinero", id, "preparando:", pedido.Comida)

		if pedido.Especificacion != "" {
			fmt.Println("   Especificación:", pedido.Especificacion)
		}

		time.Sleep(time.Duration(pedido.Tiempo) * time.Second)

		pedido.Estado = "Listo"

		resultados <- pedido
	}

}
