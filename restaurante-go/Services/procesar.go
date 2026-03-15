package services

import (
	"fmt"
	models "restaurante-go/Models"
	"time"
)

// Procesa la lista de pedidos utilizando goroutines para simular a varios cocineros trabajando en paralelo
func ProcesarPedidos(listaPedidos []models.Pedido, inv *InventarioService, stats *models.Estadisticas) {
	if len(listaPedidos) == 0 {
		fmt.Println("No hay pedidos para procesar")
		return
	}

	//Guarda el momento en que inicia el procesamiento
	inicio := time.Now()

	//Canal para enviar pedidos a los cocineros
	pedidos := make(chan models.Pedido)
	//Canal para recibir pedidos terminados por los cocineros
	resultados := make(chan models.Pedido)

	//Numero de cocineros que trabajaran en paralelo
	numCocineros := 3

	//Creamos goroutines para cada cocinero
	for i := 1; i <= numCocineros; i++ {
		go Cocinero(i, pedidos, resultados, inv)
	}

	//Goroutine que envia los pedidos al canal "pedidos"
	go func() {
		for _, pedido := range listaPedidos {
			pedidos <- pedido
		}
		//Cerrar el canal cuando no hay mas pedidos
		close(pedidos)
	}()

	//Recibir resultados de los cocineros
	for i := 0; i < len(listaPedidos); i++ {
		pedido := <-resultados
		fmt.Println("Pedido listo: ",
			pedido.Comida,
			"(Cocinero", pedido.CocineroID, ")")
	}

	//Calcula el tiempo total de procesamiento
	tiempoTotal := time.Since(inicio)

	//Actualiza las estadisticas del restaurante
	stats.TotalPedidos += len(listaPedidos)
	stats.TiempoTotal += int(tiempoTotal.Seconds())
}
