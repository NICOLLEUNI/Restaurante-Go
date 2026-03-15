package services

import (
	"errors"
	models "restaurante-go/Models"
)

// Servicio para manejar el inventario del restaurante
type InventarioService struct {
	Inv *models.Inventario
}

// Crea una nueva instancia del servicio de inventario
func NewInventario() *InventarioService {
	return &InventarioService{
		Inv: &models.Inventario{
			Ingredientes: make(map[string]int),
		},
	}
}

// Inicializa el inventario con platos y cantidades predeterminados
func (s *InventarioService) Inicializar() {
	s.Inv.Mu.Lock()
	defer s.Inv.Mu.Unlock()

	s.Inv.Ingredientes["pollo"] = 20
	s.Inv.Ingredientes["res"] = 15
	s.Inv.Ingredientes["cerdo"] = 15
	s.Inv.Ingredientes["pasta"] = 40
	s.Inv.Ingredientes["arroz"] = 50
	s.Inv.Ingredientes["ensalada"] = 25
	s.Inv.Ingredientes["pescado"] = 10
}

// Reserva un ingrediente cuando se realiza un pedido
func (s *InventarioService) ReservarIngredientes(comida string) error {
	s.Inv.Mu.Lock()
	defer s.Inv.Mu.Unlock()

	cantidad, existe := s.Inv.Ingredientes[comida]
	if !existe || cantidad <= 0 {
		return errors.New("no hay suficiente " + comida)
	}

	s.Inv.Ingredientes[comida]--
	return nil
}

// Retorna una copia del inventario actual
func (s *InventarioService) MostrarInventario() map[string]int {
	s.Inv.Mu.Lock()
	defer s.Inv.Mu.Unlock()

	copia := make(map[string]int)
	for k, v := range s.Inv.Ingredientes {
		copia[k] = v
	}
	return copia
}

// Agrega platos al inventario
func (s *InventarioService) AgregarIngredientes(nombre string, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("cantidad invalida")
	}

	s.Inv.Mu.Lock()
	defer s.Inv.Mu.Unlock()

	s.Inv.Ingredientes[nombre] += cantidad
	return nil
}

// Elimina platos del inventario
func (s *InventarioService) EliminarIngredientes(nombre string, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("cantidad invalida")
	}

	s.Inv.Mu.Lock()
	defer s.Inv.Mu.Unlock()

	actual, existe := s.Inv.Ingredientes[nombre]
	if !existe || actual < cantidad {
		return errors.New("no hay suficiente para eliminar")
	}

	s.Inv.Ingredientes[nombre] -= cantidad
	return nil
}

// Verifica si un plato esta disponible en el inventario
func (s *InventarioService) VerificarDisponibilidad(comida string) bool {
	s.Inv.Mu.Lock()
	defer s.Inv.Mu.Unlock()

	cantidad, existe := s.Inv.Ingredientes[comida]
	return existe && cantidad > 0
}

// Muestra el menu disponible basado en los ingredientes en el inventario
func (s *InventarioService) MostrarMenu() []string {
	s.Inv.Mu.Lock()
	defer s.Inv.Mu.Unlock()

	var menu []string
	for ingrediente, cantidad := range s.Inv.Ingredientes {
		if cantidad > 0 {
			menu = append(menu, ingrediente)
		}
	}
	return menu
}
