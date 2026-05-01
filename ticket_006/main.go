package main

import (
	"fmt"
	"net/http"
	"order-api/handlers"
	"time"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Printf("[%s] %s - %v\n", r.Method, r.URL.Path, duration)
	})
}

func main() {

	http.HandleFunc("/order", LoggerMiddleware(handlers.HandleOrder))
	http.HandleFunc("/order/get", LoggerMiddleware(handlers.HandleGetOrder))

	fmt.Println("Server is starting on http://localhost:8086...")

	if err := http.ListenAndServe(":8086", nil); err != nil {
		fmt.Printf("Gagal menjalankan server: %v\n", err)
	}

}
