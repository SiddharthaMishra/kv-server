package main

import "testing"

var s = &Storage{hashmap: make(map[string]string)}

// checks if values are getting put in storage correctly
func TestPutValueInStorage(t *testing.T) {
	s.putValue(&Pair{Key: "testKey", Value: "testValue"})

	i, ok := s.hashmap["testKey"]

	if !ok {
		t.Errorf("Key 'testKey' does not exist in the database")
	}

	if i != "testValue" {
		t.Errorf("Expected value 'testValue' got %v", a.Storage.hashmap["testKey"])
	}
}

// checks if values are getting retrieved from sotrage properly
func TestGetValueFromStorage(t *testing.T) {
	s.getValue("testKey")

	i, ok := s.hashmap["testKey"]

	if !ok {
		t.Errorf("Key 'testKey' does not exist in the database")
	}

	if i != "testValue" {
		t.Errorf("Expected value 'testValue' got %v", i)
	}
}

// checks if values not present are rejected properly
func TestGetWrongValueFromStorage(t *testing.T) {
	_, ok := s.getValue("notpresent")
	if ok {
		t.Errorf("Key 'testKey' should not exist in the database")
	}
}
