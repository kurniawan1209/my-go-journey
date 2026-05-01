package handlers

import (
	"encoding/json"
	"net/http"
	"order-api/models"
)

func HandleOrder(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.OrderResponse{
			Status:  "method_not_allowed",
			Message: "Only POST requests are allowed",
		})
		return
	}

	var body models.Order
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.OrderResponse{
			Status:  "bad_request",
			Message: "Invalid JSON payload",
		})
		return
	}

	models.GlobalOrderDB.Set(body.ID, body)
	response := models.OrderResponse{
		Status:  "success",
		Message: "Order received",
		Order:   &body,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

}

func HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.OrderResponse{
			Status:  "method_not_allowed",
			Message: "Only GET requests are allowed",
		})
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.OrderResponse{
			Status:  "bad_request",
			Message: "Missing 'id' query parameter",
		})
		return
	}

	if order, exists := models.GlobalOrderDB.Get(id); exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.OrderResponse{
			Status: "success",
			Order:  &order,
		})
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.OrderResponse{
			Status:  "not_found",
			Message: "Order not found",
			Order:   nil,
		})
	}

}
