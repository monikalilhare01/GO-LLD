package handlers

import (
	"encoding/json"
	"net/http"
	"survey-backend/config"
	"survey-backend/models"
	"time"

	"github.com/gorilla/mux"
)

func SubmitResponse(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.SubmittedAt = time.Now()
	response.UpdatedAt = time.Now()
	config.DB.Create(&response)
	json.NewEncoder(w).Encode(response)
}

func EditResponse(w http.ResponseWriter, r *http.Request) {
	responseID := mux.Vars(r)["response_id"]
	var updatedAnswers []models.ResponseAnswer
	if err := json.NewDecoder(r.Body).Decode(&updatedAnswers); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var response models.Response
	if err := config.DB.First(&response, responseID).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	config.DB.Where("response_id = ?", responseID).Delete(&models.ResponseAnswer{})
	for _, ans := range updatedAnswers {
		ans.ResponseID = response.ID
		config.DB.Create(&ans)
	}
	response.UpdatedAt = time.Now()
	config.DB.Save(&response)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Response updated successfully"})
}

func GetResponse(w http.ResponseWriter, r *http.Request) {
	var responses []models.Response
	surveyID := mux.Vars(r)["survey_id"]
	config.DB.Preload("ResponseAnswer").Where("survey_id = ?", surveyID).Find(&responses)
	json.NewEncoder(w).Encode(responses)

}
