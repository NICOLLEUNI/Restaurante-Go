package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"restaurante-go/Models"
)

func IngresarPedidos() []models.Pedido {

	var pedidos []models.Pedido

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Sistema de Pedidos ===")

	for {

		fmt.Print("Ingrese comida: ")
		comida, _ := reader.ReadString('\n')
		comida = strings.TrimSpace(comida)

		fmt.Print("Especificación (ej: sin cebolla) o enter si no hay: ")
		esp, _ := reader.ReadString('\n')
		esp = strings.TrimSpace(esp)

		fmt.Print("Tiempo de preparación (segundos): ")
		tiempoStr, _ := reader.ReadString('\n')
		tiempoStr = strings.TrimSpace(tiempoStr)

		tiempo, _ := strconv.Atoi(tiempoStr)

		pedido := models.Pedido{
			ID:             len(pedidos) + 1,
			Comida:         comida,
			Especificacion: esp,
			Tiempo:         tiempo,
			Estado:         "Pendiente",
		}

		pedidos = append(pedidos, pedido)

		fmt.Print("¿Desea agregar otro pedido? (s/n): ")
		resp, _ := reader.ReadString('\n')
		resp = strings.TrimSpace(resp)

		if resp != "s" {
			break
		}

	}

	return pedidos
}
