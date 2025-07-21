package handlers

import (
	"encoding/json"
	"net/http"
	"survey-backend/config"
	"survey-backend/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
