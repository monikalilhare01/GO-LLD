package routes

import (
	"survey-backend/handlers"

	"github.com/gorilla/mux"
)

func RegisterSurveyRoutes(r *mux.Router) {
	r.HandleFunc("/surveys", handlers.CreateSurvey).Methods("POST")
	r.HandleFunc("/surveys/{survey_id}/responses", handlers.SubmitResponse).Methods("POST")
	r.HandleFunc("/responses/{response_id}", handlers.EditResponse).Methods("PUT")
	r.HandleFunc("/surveys/{survey_id}/responses", handlers.GetResponse).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
}
