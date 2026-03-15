package services

import (
	"fmt"
	models "restaurante-go/Models"
	"time"
)

// Funcion que simula a un cocinero preparando pedidos
func Cocinero(id int, pedidos <-chan models.Pedido, resultados chan<- models.Pedido, inv *InventarioService) {
	//El cocinero recibe pedidos del canal "pedidos" y los procesa
	for pedido := range pedidos {
		fmt.Println("Cocinero", id, "preparando:", pedido.Comida)

		//Si el pedido tiene especificaciones, las muestra
		if pedido.Especificacion != "" {
			fmt.Println("   Especificacion:", pedido.Especificacion)
		}

		//Valida el inventario antes de preparar el pedido
		if err := inv.ReservarIngredientes(pedido.Comida); err != nil {
			//Si no hay ingredientes suficientes, marca el pedido como fallido
			fmt.Println("   Error:", err)
			pedido.Estado = "Fallido"
			resultados <- pedido
			continue
		}

		//Simula el tiempo de preparacion del pedido
		time.Sleep(time.Duration(pedido.Tiempo) * time.Second)

		//Actualiza el estado del pedido a "Listo"
		pedido.Estado = "Listo"
		//Asigna el ID del cocinero que preparo el pedido
		pedido.CocineroID = id

		//Envia el resultado al canal "resultados"
		resultados <- pedido
	}
}

// Retorna una lista de cocineros disponibles
func ObtenerCocineros() []models.Cocinero {
	//Lista estatica de cocineros con sus IDs y nombres
	return []models.Cocinero{
		{ID: 1, Nombre: "Juan"},
		{ID: 2, Nombre: "Maria"},
		{ID: 3, Nombre: "Carlos"},
	}
}
