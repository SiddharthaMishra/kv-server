package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a *App

func TestMain(t *testing.T) {
	a = initApp()
	go http.ListenAndServe(":8000", a.router)
}

func TestPutValueInStorage(t *testing.T) {
	a.Storage.putValue(&Pair{Key: "testKey", Value: "testValue"})

	i, ok := a.Storage.hashmap["testKey"]

	if ok == false {
		t.Errorf("Key 'testKey' does not exist in the database")
	}

	if i != "testValue" {
		t.Errorf("Expected value 'testValue' got %v", a.Storage.hashmap["testKey"])
	}
}

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

func TestWrongPostRequest(t *testing.T) {

	payload := []byte(`{"Key": "TEST"`)

	req, _ := http.NewRequest("POST", "/api", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	i, ok := m["error"]

	if ok == false {
		t.Errorf("Key error not present")
	}

	if i != "Invalid request payload" {
		t.Errorf("Unexpected error text %v", i)
	}

}

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
