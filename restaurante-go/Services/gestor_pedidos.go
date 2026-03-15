package services

import (
	"bufio"
	"fmt"
	"os"
	models "restaurante-go/Models"
	"strconv"
	"strings"
)

// Permite al usuario ingresar pedidos, con opciones para agregar, cancelar y finalizar la orden
func IngresarPedidos(inv *InventarioService) ([]models.Pedido, []models.Pedido) {
	var pedidos []models.Pedido
	var pedidosCancelar []models.Pedido
	reader := bufio.NewReader(os.Stdin)

	//Muestra cocineros disponibles
	cocineros := ObtenerCocineros()
	fmt.Println("\n=== COCINEROS DISPONIBLES ===")
	for _, c := range cocineros {
		fmt.Printf("  %d. %s\n", c.ID, c.Nombre)
	}

	//Muestra el menu disponible
	menu := inv.MostrarMenu()
	fmt.Println("\n=== MENU DISPONIBLE ===")
	for i, item := range menu {
		fmt.Printf("  %d. %s\n", i+1, item)
	}

	//Bucle para gestionar pedidos
	fmt.Println("\n=== SISTEMA DE PEDIDOS DEL RESTAURANTE ===")

	for {
		fmt.Println("\nOpciones:")
		fmt.Println("  1. Agregar pedido")
		fmt.Println("  2. Cancelar pedido")
		fmt.Println("  3. Terminar y procesar")
		fmt.Print("Seleccione opcion: ")

		opcionStr, _ := reader.ReadString('\n')
		opcionStr = strings.TrimSpace(opcionStr)
		opcion, _ := strconv.Atoi(opcionStr)

		switch opcion {
		case 1:
			//Solicitar detalles del pedido
			fmt.Print("Ingrese comida: ")
			comida, _ := reader.ReadString('\n')
			comida = strings.TrimSpace(comida)

			//Verificar disponibilidad del plato
			if !inv.VerificarDisponibilidad(comida) {
				fmt.Println("Error: Comida no disponible en el menu")
				continue
			}

			//Solicitar especificaciones
			fmt.Print("Especificacion (ej: sin cebolla) o ENTER si no hay: ")
			esp, _ := reader.ReadString('\n')
			esp = strings.TrimSpace(esp)

			//Solicitar tiempo de preparacion
			fmt.Print("Tiempo de preparacion (segundos): ")
			tiempoStr, _ := reader.ReadString('\n')
			tiempoStr = strings.TrimSpace(tiempoStr)
			tiempo, _ := strconv.Atoi(tiempoStr)

			//Solicitar cocinero
			fmt.Print("Seleccione cocinero (1-3): ")
			cocineroStr, _ := reader.ReadString('\n')
			cocineroStr = strings.TrimSpace(cocineroStr)
			cocineroID, _ := strconv.Atoi(cocineroStr)
			if cocineroID < 1 || cocineroID > 3 {
				cocineroID = 1
			}

			//Crear pedido y agregar a la lista
			pedido := models.Pedido{
				ID:             len(pedidos) + len(pedidosCancelar) + 1,
				Comida:         comida,
				Especificacion: esp,
				Tiempo:         tiempo,
				Estado:         "Pendiente",
				CocineroID:     cocineroID,
			}

			pedidos = append(pedidos, pedido)
			fmt.Println("Pedido agregado")

		case 2:
			//Verificar que hay pedidos para cancelar
			if len(pedidos) == 0 {
				fmt.Println("No hay pedidos para cancelar")
				continue
			}

			//Mostrar pedidos pendientes
			fmt.Println("\nPedidos pendientes:")
			for _, p := range pedidos {
				fmt.Printf("  ID %d: %s (Cocinero %d)\n", p.ID, p.Comida, p.CocineroID)
			}

			//Solicitar ID del pedido a cancelar
			fmt.Print("Ingrese ID del pedido a cancelar: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)

			//Buscar y cancelar el pedido
			for i, p := range pedidos {
				if p.ID == id {
					pedidosCancelar = append(pedidosCancelar, p)
					pedidos = append(pedidos[:i], pedidos[i+1:]...)
					fmt.Println("Pedido cancelado")
					break
				}
			}

		case 3:
			//Retornar pedidos y cancelados para procesar
			return pedidos, pedidosCancelar

		default:
			fmt.Println("Opcion invalida")
		}
	}

	return pedidos, pedidosCancelar
}
