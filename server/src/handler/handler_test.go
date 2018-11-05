package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"../catalog"
	"github.com/gorilla/mux"
)

var router *mux.Router
var scedule catalog.Scedule

func init() {

	// Setup mocks
	scedule = &catalog.DefaultScedule{}
	layout := "15:04"
	startTime, _ := time.Parse(layout, "10:30")
	t := catalog.Training{11, "Tech Inf", "Technische Informatik", catalog.Teacher{101, "Norbert Jung", 55, "jung@fbrs.de"}, startTime, 250}
	scedule.AddTraining(t)

	router = TrainingsRouter(scedule)
}

func TestGetAllTrainings(t *testing.T) {

	// When: GET trainings/trainings is called
	req, _ := http.NewRequest("GET", "/trainings", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Then: status is 200
	assert.Equal(t, http.StatusOK, rr.Code)

	// And: Body contains 1 product
	expected := `[{"Id":11,"name":"Tech Inf","description":"Technische Informatik","Lecturer":{"Id":101,"Name":"Norbert Jung","Age":55,"EMail":"jung@fbrs.de"},"Time":"0000-01-01T10:30:00Z","Price":250}]`
	result := rr.Body.String()
	result = result[0 : len(result)-1]
	assert.Equal(t, expected, result)
}

func TestGetSingleTrainings(t *testing.T) {

	// When: GET /trainings is called
	req, _ := http.NewRequest("GET", "/trainings/11", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Then: status is 200
	assert.Equal(t, http.StatusOK, rr.Code)

	// And: Body contains 1 product
	expected := `{"Id":11,"name":"Tech Inf","description":"Technische Informatik","Lecturer":{"Id":101,"Name":"Norbert Jung","Age":55,"EMail":"jung@fbrs.de"},"Time":"0000-01-01T10:30:00Z","Price":250}`
	result := rr.Body.String()
	result = result[0 : len(result)-1]
	assert.Equal(t, expected, result)
}

func TestDeleteTrainings(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/trainings/11", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	assert.Equal(t, 0, len(scedule.AllTrainings()))
}

func TestUpdateTraining(t *testing.T) {
	x, _ := scedule.TrainingByID(11)
	x.Name = "Mathematik"
	json, _ := json.Marshal(x)
	body := bytes.NewBuffer(json)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/trainings/%d", x.ID), body)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
