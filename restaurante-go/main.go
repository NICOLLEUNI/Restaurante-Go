package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	models "restaurante-go/Models"
	"restaurante-go/services"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	stats := models.Estadisticas{}

	for {

		fmt.Println("\n=== SISTEMA DEL RESTAURANTE ===")
		fmt.Println("1. Hacer pedido")
		fmt.Println("2. Ver estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		opcion, _ := strconv.Atoi(input)

		switch opcion {

		case 1:

			pedidos := services.IngresarPedidos()
			services.ProcesarPedidos(pedidos, &stats)

		case 2:

			services.MostrarEstadisticas(stats)

		case 3:

			fmt.Println("Saliendo del sistema...")
			return

		default:

			fmt.Println("Opción inválida")

		}

	}

}
