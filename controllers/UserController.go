package controllers

import (
	"encoding/json"
	"net/http"
	"user_management_go/storage"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(storage.Users)
	if err != nil {
		http.Error(w, "Encountered an Error", http.StatusInternalServerError) // error 500
		return
	}
}
