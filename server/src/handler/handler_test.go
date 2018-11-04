package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"bitbucket.org/rwirdemann/rest-apis-go/catalog"
	"github.com/gorilla/mux"
)

var router *mux.Router
var repository catalog.Repository

func init() {

	// Setup mocks
	repository = &catalog.DefaultRepository{}
	p := catalog.Product{Id: 1, Name: "Schuhe"}
	repository.AddProduct(p)

	router = NewRouter(repository)
}

func TestGetAllProducts(t *testing.T) {

	// When: GET /catalog/products is called
	req, _ := http.NewRequest("GET", "/catalog/products", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Then: status is 200
	assert.Equal(t, http.StatusOK, rr.Code)

	// And: Body contains 1 product
	expected := `[{"Id":1,"Name":"Schuhe","Description":"","Category":"","Price":0}]`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetSingleProducts(t *testing.T) {

	// When: GET /catalog/products is called
	req, _ := http.NewRequest("GET", "/catalog/products/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Then: status is 200
	assert.Equal(t, http.StatusOK, rr.Code)

	// And: Body contains 1 product
	expected := `{"Id":1,"Name":"Schuhe","Description":"","Category":"","Price":0}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestAddProduct(t *testing.T) {
	json, _ := json.Marshal(catalog.Product{Name: "Jacke"})
	body := bytes.NewBuffer(json)
	req, _ := http.NewRequest("POST", "/catalog/products", body)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Regexp(t, "/catalog/products/[0-9]+", rr.Header()["Location"][0])
	assert.True(t, repository.Contains("Jacke"))
}

func TestDeleteProducts(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/catalog/products/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	assert.Equal(t, 0, len(repository.AllProducts()))
}

func TestUpdateProduct(t *testing.T) {
	p, _ := repository.ProductById(1)
	p.Name = "Herrenschuhe"
	json, _ := json.Marshal(p)
	body := bytes.NewBuffer(json)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/catalog/products/%d", p.Id), body)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.True(t, repository.Contains("Herrenschuhe"))
}
