package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../catalog"
)

func MakeGetTrainingHandler(scedule catalog.Scedule) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if products, ok := scedule.TrainingByID(int32(id)); ok {
			if json, err := json.Marshal(products); err == nil {
				fmt.Fprintln(w, string(json))
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}
}

func MakeDeleteTrainingHandler(scedule catalog.Scedule) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if scedule.Delete(int32(id)) {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func MakeTrainingsHandler(scedule catalog.Scedule) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		trainings := scedule.AllTrainings()

		if json, err := json.Marshal(trainings); err == nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, string(json))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func MakeAddTrainingHandler(scedule catalog.Scedule) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var p catalog.Training
		if err = json.Unmarshal(body, &p); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id := scedule.AddTraining(p)
		w.Header().Set("Location", fmt.Sprintf("%s/%d", r.URL.String(), id))
		w.WriteHeader(http.StatusCreated)
	}
}

func MakeUpdateTrainingHandler(scedule catalog.Scedule) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var p catalog.Training
		if err = json.Unmarshal(body, &p); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if scedule.Update(int32(id), p) {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
