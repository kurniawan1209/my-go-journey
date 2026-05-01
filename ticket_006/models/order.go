package models

import "sync"

type Order struct {
	ID           string  `json:"id"`
	CustomerName string  `json:"customer_name"`
	TotalAmount  float64 `json:"total_amount"`
}

type OrderResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Order   *Order `json:"order,omitempty"`
}

type OrderDB struct {
	mu     sync.RWMutex
	orders map[string]Order
}

func (db *OrderDB) Set(id string, order Order) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.orders[id] = order
}

func (db *OrderDB) Get(id string) (Order, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	order, ok := db.orders[id]
	return order, ok
}

var GlobalOrderDB = &OrderDB{
	orders: make(map[string]Order),
}
