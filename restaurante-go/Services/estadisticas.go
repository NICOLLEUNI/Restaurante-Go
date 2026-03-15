package services

import (
	"fmt"
	models "restaurante-go/Models"
)

func MostrarEstadisticas(stats models.Estadisticas) {

	fmt.Println("\n=== ESTADÍSTICAS DEL RESTAURANTE ===")

	if stats.TotalPedidos == 0 {
		fmt.Println("Aún no hay pedidos procesados")
		return
	}

	promedio := float64(stats.TiempoTotal) / float64(stats.TotalPedidos)

	fmt.Println("Total de pedidos:", stats.TotalPedidos)
	fmt.Println("Tiempo total de preparación:", stats.TiempoTotal, "segundos")
	fmt.Println("Tiempo promedio por pedido:", promedio, "segundos")

}
