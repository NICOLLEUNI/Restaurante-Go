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

	//Canal para recibir pedidos terminados por los cocineros
	resultados := make(chan models.Pedido, len(listaPedidos))

	//Numero de cocineros disponibles
	numCocineros := 3

	//Creamos un canal individual por cocinero
	canales := make(map[int]chan models.Pedido)
	for i := 1; i <= numCocineros; i++ {
		canales[i] = make(chan models.Pedido)
		go Cocinero(i, canales[i], resultados, inv)
	}

	//Goroutine que envia cada pedido al canal de su cocinero asignado
	go func() {
		for _, pedido := range listaPedidos {
			canal, ok := canales[pedido.CocineroID]
			if !ok {
				canal = canales[1] // Si el ID no existe, va al cocinero 1 por defecto
			}
			canal <- pedido
		}
		//Cerrar todos los canales cuando no hay mas pedidos
		for _, c := range canales {
			close(c)
		}
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
