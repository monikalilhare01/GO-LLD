package main

import (
	"log"
	"net/http"

	"survey-backend/config"
	"survey-backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()
	r := mux.NewRouter()
	routes.RegisterSurveyRoutes(r)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is up and running"))
	})

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
