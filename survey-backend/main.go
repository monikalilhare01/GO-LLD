package main

import (
	"fmt"
	"log"
	"net/http"

	"survey-backend/config"
	"survey-backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	_, err := config.Connect()
	if err != nil {
		fmt.Println("Unable to connect to db ")
		return
	}
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is up and running"))
	})
	log.Println("Server started at :8080")
	routes.RegisterSurveyRoutes(r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
