package handler

import (
	"../catalog"
	"github.com/gorilla/mux"
)

// TrainingsRouter for routing
func TrainingsRouter(scedule catalog.Scedule) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/trainings", MakeTrainingsHandler(scedule)).Methods("GET")
	r.HandleFunc("/trainings", MakeAddTrainingHandler(scedule)).Methods("POST")
	r.HandleFunc("/trainings/{id}", MakeGetTrainingHandler(scedule)).Methods("GET")
	r.HandleFunc("/trainings/{id}", MakeDeleteTrainingHandler(scedule)).Methods("DELETE")
	r.HandleFunc("/trainings/{id}", MakeUpdateTrainingHandler(scedule)).Methods("PUT")

	return r
}
