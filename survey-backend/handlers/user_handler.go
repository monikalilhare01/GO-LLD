package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/monikalilhare01/survey-backend/config"
	"github.com/monikalilhare01/survey-backend/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var existingUser models.User
	if err := config.DB.Where("user_email = ?", user.UserEmail).First(&existingUser).Error; err == nil {
		http.Error(w, "User with this email already exists", http.StatusConflict)
		return
	}
	config.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
