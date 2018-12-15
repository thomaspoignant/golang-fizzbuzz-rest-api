package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FizzbuzzResponse struct {
	Status int      `json:"status"`
	Value  []string `json:"value"`
}

// utility func who make the request
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// test a succesful request to the API
func TestSuccess(t *testing.T) {
	router := SetupRouter()

	w := performRequest(router, "GET", "/v1/fizzbuzz?string1=Fizz&string2=Buzz&int1=3&int2=5&limit=15")
	assert.Equal(t, http.StatusOK, w.Code)

	var response FizzbuzzResponse
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	assert.Equal(t, want, response.Value)
	assert.Equal(t, 200, response.Status)
}

// test the api with error in the parameters but with the good type
func TestInvalidParameter_withGoodType(t *testing.T) {
	router := SetupRouter()

	w := performRequest(router, "GET", "/v1/fizzbuzz?string1=Fizz&string2=Buzz&int1=3&int2=5&limit=-15")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response FizzbuzzResponse
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	assert.Equal(t, 400, response.Status)
}

// test the api with a wrong type parameter
func TestInvalidParameter_withWrongType(t *testing.T) {
	router := SetupRouter()

	w := performRequest(router, "GET", "/v1/fizzbuzz?string1=Fizz&string2=Buzz&int1=3&int2=5&limit=a")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response FizzbuzzResponse
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	assert.Equal(t, 400, response.Status)
}

// test the api with a missing parameter
func TestInvalidParameter_missingParameter(t *testing.T) {
	router := SetupRouter()

	w := performRequest(router, "GET", "/v1/fizzbuzz?string1=Fizz&string2=Buzz&int1=3&int2=5")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response FizzbuzzResponse
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	assert.Equal(t, 400, response.Status)
}
