package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/monikalilhare01/survey-backend/config"
	"github.com/monikalilhare01/survey-backend/models"
)

type CreateSurveyResponse struct {
	SurveyID int    `json:"survey_id"`
	Message  string `json:"message"`
}

func CreateSurvey(w http.ResponseWriter, r *http.Request) {
	var survey models.Survey
	if err := json.NewDecoder(r.Body).Decode(&survey); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	survey.CreatedAt = time.Now()
	// Then insert the questions
	for _, q := range survey.Questions {
		q.SurveyID = survey.ID
		config.DB.Create(&q)
	}

	json.NewEncoder(w).Encode(CreateSurveyResponse{
		SurveyID: int(survey.ID),
		Message:  "Survey created successfully",
	})
}

func GetSurveyByID(w http.ResponseWriter, r *http.Request) {

}
