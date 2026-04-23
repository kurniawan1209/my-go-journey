package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

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

var globalOrderDB = &OrderDB{
	orders: make(map[string]Order),
}

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Printf("[%s] %s - %v\n", r.Method, r.URL.Path, duration)
	})

}

func handleOrder(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(OrderResponse{
			Status:  "method_not_allowed",
			Message: "Only POST requests are allowed",
		})
		return
	}

	var body Order
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(OrderResponse{
			Status:  "bad_request",
			Message: "Invalid JSON payload",
		})
		return
	}

	globalOrderDB.Set(body.ID, body)

	response := OrderResponse{
		Status:  "success",
		Message: "Order received",
		Order:   &body,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

}

func handleGetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(OrderResponse{
			Status:  "method_not_allowed",
			Message: "Only GET requests are allowed",
		})
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(OrderResponse{
			Status:  "bad_request",
			Message: "Missing 'id' query parameter",
		})
		return
	}

	if order, exists := globalOrderDB.Get(id); exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(OrderResponse{
			Status:  "success",
			Message: "Order found",
			Order:   &order,
		})
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(OrderResponse{
			Status:  "not_found",
			Message: "Order not found",
		})
		return
	}

}

func main() {

	http.HandleFunc("/orders", LoggerMiddleware(handleOrder))
	http.HandleFunc("/order", LoggerMiddleware(handleGetOrder))

	fmt.Println("Server is starting on http://localhost:8086...")

	if err := http.ListenAndServe(":8086", nil); err != nil {
		fmt.Printf("Gagal menjalankan server: %v\n", err)
	}

}
