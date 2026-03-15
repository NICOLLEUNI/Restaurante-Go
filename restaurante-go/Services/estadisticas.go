package services

import (
	"fmt"
	models "restaurante-go/Models"
)

// Muestra las estadísticas del restaurante
func MostrarEstadisticas(stats models.Estadisticas) {

	fmt.Println("\n=== ESTADISTICAS DEL RESTAURANTE ===")

	//Verifica si no hay pedidos procesados para evitar division por cero
	if stats.TotalPedidos == 0 {
		fmt.Println("Aun no hay pedidos procesados")
		return
	}

	//Calcula el tiempo promedio por pedido
	promedio := float64(stats.TiempoTotal) / float64(stats.TotalPedidos)

	fmt.Println("Total de pedidos:", stats.TotalPedidos)
	fmt.Println("Tiempo total de preparacion:", stats.TiempoTotal, "segundos")
	fmt.Println("Tiempo promedio por pedido:", promedio, "segundos")

}
