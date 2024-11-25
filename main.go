package main

import (
	"fmt"
	"net/http"
	"receipt-processor-challenge/handlers"
)

func main() {
	http.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler)
	http.HandleFunc("/receipts/", handlers.GetPointsHandler)

	port := ":8181"

	fmt.Printf("Server running on port %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
