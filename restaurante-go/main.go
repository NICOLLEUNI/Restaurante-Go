package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	models "restaurante-go/Models"
	services "restaurante-go/Services"
)

func main() {
	//Crea servicio de inventario
	inventario := services.NewInventario()

	//Inicia con ingredientes predeterminados
	inventario.Inicializar()

	reader := bufio.NewReader(os.Stdin)

	//Guarda estadisticas del restaurante
	stats := models.Estadisticas{}

	for {
		//Menu principal del sistema
		fmt.Println("\n=== SISTEMA DEL RESTAURANTE ===")
		fmt.Println("1. Hacer pedido")
		fmt.Println("2. Ver estadisticas")
		fmt.Println("3. Ver inventario")
		fmt.Println("4. Ver menu")
		fmt.Println("5. Agregar ingredientes")
		fmt.Println("6. Eliminar ingredientes")
		fmt.Println("7. Salir")
		fmt.Print("Seleccione una opcion: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		opcion, _ := strconv.Atoi(input)

		switch opcion {
		case 1:
			pedidos, cancelados := services.IngresarPedidos(inventario)

			//Muestra pedidos cancelados si hay alguno
			if len(cancelados) > 0 {
				fmt.Println("\nPedidos cancelados:", len(cancelados))
				for _, p := range cancelados {
					fmt.Printf("  - %s (ID %d)\n", p.Comida, p.ID)
				}
			}

			//Procesa los pedidos con cocineros
			services.ProcesarPedidos(pedidos, inventario, &stats)

		case 2:
			//Muestra las estadisticas del restaurante
			services.MostrarEstadisticas(stats)

		case 3:
			//Muestra el inventario actual
			fmt.Println("\n=== INVENTARIO ACTUAL ===")
			stock := inventario.MostrarInventario()
			for ingrediente, cantidad := range stock {
				fmt.Printf("  %s: %d unidades\n", ingrediente, cantidad)
			}

		case 4:
			//Muestra el menu disponible
			fmt.Println("\n=== MENU DISPONIBLE ===")
			menu := inventario.MostrarMenu()
			for i, item := range menu {
				fmt.Printf("  %d. %s\n", i+1, item)
			}

		case 5:
			//Agregar ingredientes al inventario
			fmt.Print("Nombre del ingrediente: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			fmt.Print("Cantidad: ")
			cantidadStr, _ := reader.ReadString('\n')
			cantidadStr = strings.TrimSpace(cantidadStr)
			cantidad, _ := strconv.Atoi(cantidadStr)

			if err := inventario.AgregarIngredientes(nombre, cantidad); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Agregado correctamente")
			}

		case 6:
			//Eliminar ingredientes del inventario
			fmt.Print("Nombre del ingrediente: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			fmt.Print("Cantidad: ")
			cantidadStr, _ := reader.ReadString('\n')
			cantidadStr = strings.TrimSpace(cantidadStr)
			cantidad, _ := strconv.Atoi(cantidadStr)

			if err := inventario.EliminarIngredientes(nombre, cantidad); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Eliminado correctamente")
			}

		case 7:
			//Salir del sistema
			fmt.Println("Saliendo del sistema...")
			return

		default:
			fmt.Println("Opcion invalida")
		}
	}
}
