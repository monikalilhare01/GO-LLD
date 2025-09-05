package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"survey-backend/config"
	"survey-backend/models"

	"gorm.io/gorm"
)

func CreateSurvey(w http.ResponseWriter, r *http.Request) {
	var survey models.Survey
	if err := json.NewDecoder(r.Body).Decode(&survey); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := config.DB.Table("surveys").Where("id = ?", survey.ID).First(&survey).Error
	if err == nil {
		http.Error(w, "Survey already exists", http.StatusOK)
		return
	} else if err.Error() != "record not found" && !errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Manually set SurveyID for each question
	for i := range survey.Questions {
		survey.Questions[i].SurveyID = survey.ID
	}

	// First create the survey
	if err := config.DB.Create(&survey).Error; err != nil {
		http.Error(w, "Failed to create survey", http.StatusInternalServerError)
		return
	}

	// Then insert the questions
	for _, q := range survey.Questions {
		q.SurveyID = survey.ID
		config.DB.Create(&q)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"survey_id": survey.ID,
		"message":   "Survey created successfully",
	})
}
