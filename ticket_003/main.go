package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

var orderDB = make(map[string]Order)

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

	response := OrderResponse{
		Status:  "success",
		Message: "Order received",
		Order:   &body,
	}
	w.Header().Set("Content-Type", "application//json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

}

func main() {

	http.HandleFunc("/orders", handleOrder)

	fmt.Println("Server is starting on http://localhost:8086...")

	if err := http.ListenAndServe(":8086", nil); err != nil {
		fmt.Printf("Gagal menjalankan server: %v\n", err)
	}

}
