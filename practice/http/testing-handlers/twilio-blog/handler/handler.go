package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ProcessDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var d Data
	err := json.NewDecoder(r.Body).Decode(&d)
	fmt.Println("d -> ", d)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	if d.Name == "" || d.Email == "" {
		http.Error(w, "Missing data", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"status":  "success",
		"message": "Data processed successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
