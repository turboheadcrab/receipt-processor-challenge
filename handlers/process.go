package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/store"
	"receipt-processor-challenge/utils"
)

func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid receipt", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	points := utils.CalculatePoints(receipt)

	store.SaveReceipt(id, points)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
