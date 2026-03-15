package models

import "sync"

type Inventario struct {
	Mu           sync.Mutex
	Ingredientes map[string]int
}
