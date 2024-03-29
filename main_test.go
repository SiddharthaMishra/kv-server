package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a *App

// setup for the tests
func TestMain(t *testing.T) {
	a = initApp()
	go http.ListenAndServe(":8000", a.router)
}

// checks post requests storing values in the server
func TestPostRequest(t *testing.T) {

	payload := []byte(`{"Key": "TEST", "Value":"TESTING"}`)

	req, _ := http.NewRequest("POST", "/api", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["Key"] != "TEST" {
		t.Errorf("Expected response key to be 'TEST'. Got '%v'", m["Key"])
	}

	if m["Value"] != "TESTING" {
		t.Errorf("Expected response value to be 'TESTING'. Got '%v'", m["Value"])
	}

}

// check if a wrong post request is giving the right error
func TestWrongPostRequest(t *testing.T) {

	payload := []byte(`{"Key": "TEST"`)

	req, _ := http.NewRequest("POST", "/api", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	i, ok := m["error"]

	if !ok {
		t.Errorf("Key 'error' not present")
	}

	if i != "Invalid request payload" {
		t.Errorf("Unexpected error text %v", i)
	}

}

// checks if get requests work properly
func TestGetRequest(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/TEST", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["Key"] != "TEST" {
		t.Errorf("Expected response key to be 'TEST'. Got '%v'", m["Key"])
	}

	if m["Value"] != "TESTING" {
		t.Errorf("Expected response value to be 'TESTING'. Got '%v'", m["Value"])
	}

}

// check if get requests with incorrect key give the right error
func TestWrongGetRequest(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/NotThere", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]interface{}

	json.Unmarshal(response.Body.Bytes(), &m)

	i, ok := m["error"]

	if !ok {
		t.Errorf("Key 'error' not present")
	}

	if i != "Key not found" {
		t.Errorf("Unexpected error text %v", i)
	}
}

// check if the response codes are equal
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.router.ServeHTTP(rr, req)

	return rr
}
