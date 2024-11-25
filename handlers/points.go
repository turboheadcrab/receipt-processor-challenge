package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge/store"
	"strings"
)

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasSuffix(r.URL.Path, "/points") {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 {
		http.Error(w, "Invalid URL structure", http.StatusBadRequest)
		return
	}

	id := parts[2]
	points, exists := store.GetPoints(id)
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
